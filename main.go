package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Client struct {
	reader *bufio.Reader
	conn   *net.Conn
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

type Server struct {
	host, port string
}

func NewServer(host, port string) *Server {
	return &Server{
		host: host,
		port: port,
	}
}
func (s *Server) Run() {
	address := fmt.Sprintf("%s:%s", s.host, s.port)
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
func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "4222"
	}
	server := NewServer(host, port)
	server.Run()
}
