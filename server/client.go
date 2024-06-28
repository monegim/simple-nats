package server

import (
	"bufio"
	"net"
)

const (
	CLIENT = iota
	ROUTER
)

type Client struct {
	reader *bufio.Reader
	conn   *net.Conn
}
