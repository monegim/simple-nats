package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Options struct {
	Host string 
	Port int
}

func NewServer(opts *Options) *Options {
	return opts
}
func (s *Options) Run() {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	listener, err := net.Listen("tcp", address)
	defer listener.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("starting server on address: %s\n", address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		client := Client{
			conn: &conn,
		}
		go client.server()
	}
}

func (c *Client) server() {
	c.reader = bufio.NewReader(*c.conn)
	for {

		msg, err := c.reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		log.Println(msg)
	}
}
