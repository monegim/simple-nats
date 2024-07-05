package main

import (
	"log"
	"os"
	"simple-nats/server"
	"strconv"
)

func main() {
	var port int
	var err error
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port_env := os.Getenv("PORT")
	if port_env == "" {
		port = 4222
	} else {
		port, err = strconv.Atoi(port_env)
		if err != nil {
			log.Fatal(err)
		}
	}

	opts := &server.Options{
		Host: host,
		Port: port,
	}
	// if port == "" {
	// 	port = "4222"
	// }

	s, err := server.NewServer(opts)
	if err != nil {
		server.PrintAndDie(err.Error())
	}

	// TODO: ConfigureLogger
	// s.ConfigureLogger()
	if err := server.Run(s); err != nil {

	}
}
