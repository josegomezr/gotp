package main
import (
	"fmt"
	"os"
    "errors"
    "log"
    "github.com/josegomezr/gotp/api"
    "github.com/urfave/cli"
)

func buildCli() *cli.App {
    app := cli.NewApp()
    app.Name = "gotp"
    app.Version = "0.1.0"
    app.Usage = "A TOTP Server & Utility"
    
    cli.VersionFlag = cli.BoolFlag{
        Name: "version, v",
        Usage: "print only the version",
    }

    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "port, p",
            Value: "2444",
            Usage: "GOTP HTTP Server Port",
            EnvVar: "PORT,GOTP_PORT",
        },
        cli.StringFlag{
            Name: "host",
            Value: "localhost",
            Usage: "GOTP HTTP Server Host",
            EnvVar: "HOST,GOTP_HOST",
        },
        cli.StringFlag{
            Name: "secret",
            Usage: "Secret key to check/generate the code against",
        },
        cli.StringFlag{
            Name: "code",
            Usage: "Code to be checked",
        },
    }
  
    app.Commands = []cli.Command{
      {
        Name:    "server",
        Aliases: []string{"s"},
        Usage:   "Runs GOTP HTTP server",
        Action:  func(c *cli.Context) error {
            port := c.Int("port")
            if port < 0 {
                return errors.New("Cannot start server: Invalid Port")
            }
            serve(c.String("host"), port)
            return nil
        },
        Flags: []cli.Flag {
            cli.StringFlag{
                Name: "port, p",
                Value: "2444",
                Usage: "GOTP HTTP Server Port",
                EnvVar: "PORT,GOTP_PORT",
            },
            cli.StringFlag{
                Name: "host, H",
                Value: "localhost",
                Usage: "GOTP HTTP Server Host",
                EnvVar: "HOST,GOTP_HOST",
            },
        },
      },
      {
        Name:    "check",
        Aliases: []string{"c"},
        Usage:   "Checks a TOTP code",
        Flags: []cli.Flag{
            cli.StringFlag{
                Name: "secret,s",
                Usage: "Secret key to check/generate the code against",
            },
            cli.StringFlag{
                Name: "data, d",
                Usage: "Data payload to generate a secret key",
            },
            cli.StringFlag{
                Name: "code, c",
                Usage: "Code to be checked",
            },
        },
        Action:  func(c *cli.Context) error {
            query := api.RequestValidateCode{
                Secret: c.String("secret"),
                Payload: c.String("data"),
                Code: c.String("code"),
            }
        
            if err := query.Validate(); err != nil {
                return err
            }
        
            result, err := api.Verify(query)
        
            if err != nil {
                return err
            }
            
            if ! result {
                fmt.Println("Result: Invalid")
                fmt.Println("Detail: Invalid Code")
                return err
            }
            
            stringResponse := "Valid"
            
            if ! result {
                stringResponse = "Invalid"
            }
        
            fmt.Println("Result: ", stringResponse)
            fmt.Println("Code: ", query.Code)

            return nil
        },
      },
      {
        Name:    "gencode",
        Aliases: []string{"g"},
        Usage:   "Generates a TOTP code",
        Flags: []cli.Flag{
            cli.StringFlag{
                Name: "secret,s",
                Usage: "Secret key to check/generate the code against",
            },
            cli.StringFlag{
                Name: "data, d",
                Usage: "Data payload to generate a secret key",
            },
        },
        Action:  func(c *cli.Context) error {
            query := api.RequestGenerateCode{
                Secret: c.String("secret"),
                Payload: c.String("data"),
            }
            
            err := query.Validate()

            if err != nil {
                fmt.Printf("Error in input data\n-> %s\n", err)
                return err
            }

            code, err := api.CurrentCode(query)
            
            if err != nil {
                fmt.Printf("Error in input data\n-> %s\n", err)
                return err
            }
    
            if query.Payload != "" {
                fmt.Println("Data Payload:", query.Secret)
            }

            fmt.Println("Secret Key  :", query.Secret)
            fmt.Println("Current Code:", code)
            return nil
        },
      },
    } 
    return app
}


func main() {
    app := buildCli()
    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}

