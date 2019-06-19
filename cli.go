package main
import (
    "fmt"
    "os"
	"flag"
)

var SECRET_PREFIX = ""

func main() {
	var startServer = false

	flag.BoolVar(&startServer, "server", false, "start server")
	portPtr := flag.Int("port", 2444, "server port")
	codePtr := flag.String("code", "", "Code to check")
	secretPtr := flag.String("secret", "", "Secret key")
	
	flag.Parse()

	if startServer {
		fmt.Println("Starting server at port: ", *portPtr)
		serve(*portPtr)
		return
	}

	query := ValidateQuery{}
	query.Code = *codePtr
	query.Secret = *secretPtr
	
	inputValid := query.validate()

	if !inputValid {
    	fmt.Println("Invalid Secret or Code")
		os.Exit(128)
		return
	}
	query.Secret = SECRET_PREFIX + query.Secret

	result, err := verify(query)
	if err != nil || !result {
		os.Exit(1)
		return
	}
    os.Exit(0)
    return
}
