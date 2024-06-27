package main

import (
	"log"
	"net"
	"os"
)

func main() {
	address := os.Getenv("address")
	if address == "" {
		address = "localhost:4222"
	}
	listener, err := net.Listen("tcp",address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
	}
	}