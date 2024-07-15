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
	opts *Options
}

func NewServer(opts *Options) (*Server, error) {
	return &Server{opts: opts}, nil
}
func (s *Server) Start() {
	address := fmt.Sprintf("%s:%d", s.opts.Host, s.opts.Port)
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
		client := client{
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

func (c *client) server() {
	c.reader = bufio.NewReader(*c.conn)
	for {

		msg, err := c.reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		log.Println(msg)
	}
}
