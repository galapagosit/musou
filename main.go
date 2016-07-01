package main

import (
	"flag"
	"fmt"
	"github.com/galapagosit/musou/server"
)

func main() {
	mode := flag.String("mode", "client", "mode")
	host := flag.String("host", "example.com", "host")
	port := flag.String("port", "8888", "port")
	flag.Parse()

	if *mode == "server" {
		server.StartServer(*port)
	} else {
		fmt.Println("I'm client connect to " + *host)
	}
}
