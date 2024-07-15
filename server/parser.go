package server

import "fmt"

type parserState int
type parseState struct {
	state parserState
	op    byte
}

const (
	OP_START parserState = iota
	OP_C
	OP_CO
	OP_CON
	OP_CONN
	OP_CONNE
	OP_CONNEC
	OP_CONNECT
	OP_P
	OP_PU
	OP_PUB
	OP_PUB_SPC
	PUB_ARG
	OP_PI
	OP_PIN
	OP_PING
	OP_PO
	OP_PON
	OP_PONG
	OP_S
	OP_SU
	OP_SUB
)

func (c *client) parse(buf []byte) error {
	var b byte
	var i int
	for i = 0; i < len(buf); i++ {
		b = buf[i]
		switch c.state {
		case OP_START:
			c.op = b
			switch b {
			case 'P', 'p':
				c.state = OP_P
			case 'C', 'c':
				c.state = OP_C
			case 'S':
				c.state = OP_S
			default:
				goto parseErr
			}
		default:
			goto parseErr

		}
	}
	return nil
parseErr:
	c.sendErr("")
	return fmt.Errorf("")
}
