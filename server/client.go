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

type client struct {
	reader *bufio.Reader
	conn   *net.Conn
	mu     sync.Mutex
	parseState
}

func (c *client) sendErr(err string) {

}
