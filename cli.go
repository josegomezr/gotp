package main
import (
    "fmt"
    "os"
	"flag"
)

// ConstSecretPrefix a comment
var ConstSecretPrefix = ""

// ConstIssuer a comment
var ConstIssuer = "Diggi.io"


func generateCommand(generateCmd *flag.FlagSet, generateSecret *string) {
	generateCmd.Parse(os.Args[2:])
	
	secret := ConstSecretPrefix + *generateSecret

	if !validBase32(secret) {
		fmt.Println("Invalid Secret Format")
		os.Exit(2)
	}

	query := GenerateCodeQuery{
		Secret: secret,
	}
	if ! query.Validate() {
		fmt.Println("Invalid Secret Format")
		os.Exit(3)
	}
	
	code, err := currentCode(query)
	
	if err != nil {
		fmt.Println("Invalid Secret Checksum")
		os.Exit(4)
	}
	
	fmt.Println("Secret Key  :", *generateSecret)
	fmt.Println("Current Code:", code)
	os.Exit(0)
}


func checkCommand(checkCmd *flag.FlagSet, checkSecret *string, checkCode *string) {
	checkCmd.Parse(os.Args[2:])
	secret := ConstSecretPrefix + *checkSecret

	if !validBase32(secret) {
		fmt.Println("Result: Invalid")
		fmt.Println("Detail: Invalid Secret Format")
		os.Exit(2)
	}

	query := ValidateQuery{
		Secret: secret,
		Code: *checkCode,
	}

	if ! query.Validate() {
		fmt.Println("Result: Invalid")
		fmt.Println("Detail: Invalid Input")
		os.Exit(3)
	}

	result, err := verify(query)

	if err != nil {
		fmt.Println("Result: Invalid")
		fmt.Println("Detail: Validation error.", err)
		os.Exit(4)
	}
	
	if ! result {
		fmt.Println("Result: Invalid")
		fmt.Println("Detail: Invalid Code")
		os.Exit(5)	
	}
	
	stringResponse := "Valid"
	if ! result {
		stringResponse = "Invalid"
	}
	fmt.Println("Result: ", stringResponse)
	fmt.Println("Code: ", query.Code)
}

func serverCommand(serverCmd *flag.FlagSet, serverHost *string, serverPort *int){
	serverCmd.Parse(os.Args[2:])
	fmt.Printf("Starting server at port: %s:%d", *serverHost, *serverPort)
	serve(*serverHost, *serverPort)
	return
}

func main() {
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
    serverPort :=  serverCmd.Int("port", 2444, "Server Port")
	serverHost :=  serverCmd.String("host", "localhost", "Server Host")
	
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	generateSecret := generateCmd.String("secret", "", "Secret Key")
	
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	checkSecret := checkCmd.String("secret", "", "Secret Key")
	checkCode := checkCmd.String("code", "", "Code to check")
	
	if len(os.Args) < 2 {
		fmt.Println("Missing subcommand")
		fmt.Printf("%s { server | generate | check } [options]\n", os.Args[0])
		serverCmd.Usage()
		generateCmd.Usage()
		checkCmd.Usage()
		os.Exit(1)
	}
	
	switch os.Args[1] {
	case "generate":
		generateCommand(generateCmd, generateSecret)
	case "check":
		checkCommand(checkCmd, checkSecret, checkCode)
	case "server":
		serverCommand(serverCmd, serverHost, serverPort)
		os.Exit(0)
	default:
		fmt.Println("Invalid subcommand. Expecting 'generate', 'check' or 'server'.")
		os.Exit(2)
	}
}
