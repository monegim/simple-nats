package server

import (
	"bufio"
	"net"
	"sync"
)

const (
	CLIENT = iota
	ROUTER
)

type Client struct {
	reader *bufio.Reader
	conn   *net.Conn
	mu     sync.Mutex
}
