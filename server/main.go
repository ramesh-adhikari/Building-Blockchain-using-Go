package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port number for Server")
	flag.Parse()
	// server := new(Server)
	app := NewServer(uint16(*port))
	app.Run()
}
