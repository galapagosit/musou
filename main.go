package main

import (
	"flag"
	"fmt"
	"github.com/galapagosit/musou/server"
	"github.com/galapagosit/musou/client"
)

func main() {
	mode := flag.String("mode", "client", "mode")
	host := flag.String("host", "localhost", "host")
	port := flag.String("port", "8888", "port")
	flag.Parse()

	if *mode == "server" {
		server.StartServer(*port)
	} else {
		client.StartClient(*host, *port)
		fmt.Println("I'm client connect to " + *host)
	}
}
