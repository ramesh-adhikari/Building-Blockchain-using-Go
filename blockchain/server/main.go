package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5001, "TCP Port number for Server")
	flag.Parse()
	// server := new(Server)
	// fmt.Println(*port)
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
