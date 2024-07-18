package server

import (
	"bufio"
	"net"
	"sync"
)

const (
	_CRLF_ = "\r\n"
	CLIENT = iota
	ROUTER
	pongProto = "PONG" + _CRLF_
)

type client struct {
	reader *bufio.Reader
	conn   *net.Conn
	mu     sync.Mutex
	parseState
}

func (c *client) sendErr(err string) {

}

func (c *client) processConnect(arg []byte) error {
	return nil
}

func (c *client) processPing() {
	c.sendPong()
}

func (c *client) sendPong() {

}
