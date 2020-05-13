package main

import (
	"coderunner/server"
	"flag"
)

var portNumber = flag.String("port", "52453", "Port number for server to listen on")

func init() {
	flag.StringVar(portNumber, "p", "52453", "Port number for server to listen on")
}
func main() {
	flag.Parse()
	server.StartServer(*portNumber)
}
