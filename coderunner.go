package main

import (
	"coderunner/server"
	"flag"
	"fmt"
)

var portNumber = flag.String("port", "8082", "Port number for server to listen on")

func init() {
	flag.StringVar(portNumber, "p", "8082", "Port number for server to listen on")
}
func main() {
	flag.Parse()
	fmt.Println(*portNumber)
	server.StartServer(*portNumber)
}
