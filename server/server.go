package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Options struct {
	Host string
	Port int
}

type Server struct {
}

func (server *Server) Start() {

}

func NewServer(opts *Options) (*Server, error) {
	return &Server{}, nil
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
func PrintAndDie(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func Run(s *Server) error {
	s.Start()
	return nil
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
