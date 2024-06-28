package main

import (
	"os"
	"simple-nats/server"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "4222"
	}
	server := server.NewServer(host, port)
	server.Run()
}
