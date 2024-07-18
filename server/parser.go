package server

import "fmt"

type parserState int
type parseState struct {
	state   parserState
	op      byte
	as      int
	drop    int
	argBuff []byte
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
	CONNECT_ARG
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
		case OP_C:
			switch b {
			case 'o', 'O':
				c.state = OP_CO
			default:
				goto parseErr
			}
		case OP_CO:
			switch b {
			case 'N', 'n':
				c.state = OP_CON
			default:
				goto parseErr
			}
		case OP_CON:
			switch b {
			case 'N', 'n':
				c.state = OP_CONN
			default:
				goto parseErr
			}
		case OP_CONN:
			switch b {
			case 'E', 'e':
				c.state = OP_CONNE
			default:
				goto parseErr
			}
		case OP_CONNE:
			switch b {
			case 'C', 'c':
				c.state = OP_CONNEC
			default:
				goto parseErr
			}
		case OP_CONNEC:
			switch b {
			case 'T', 't':
				c.state = OP_CONNECT
			default:
				goto parseErr
			}
		case OP_CONNECT:
			switch b {
			case ' ', '\t':
				continue
			default:
				c.state = CONNECT_ARG
				c.as = i
			}
		case CONNECT_ARG:
			switch b {
			case '\r':
				c.drop = 1
			case '\n':
				var arg []byte
				if c.argBuff != nil {
					arg = c.argBuff
					c.argBuff = nil
				} else {
					arg = buf[c.as : i-c.drop]
				}
				if err := c.processConnect(arg); err != nil {
					return err
				}
				c.drop, c.state = 0, OP_START
			default:
				if c.argBuff != nil {
					c.argBuff = append(c.argBuff, b)
				}
			}
		case OP_PI:
			switch b {
			case 'N', 'n':
				c.state = OP_PIN
			default:
				goto parseErr
			}
		case OP_PIN:
			switch b {
			case 'G', 'g':
				c.state = OP_PING
			default:
				goto parseErr
			}
		case OP_PING:
			switch b {
			case '\n':
				c.processPing()
				c.drop, c.state = 0, OP_START
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
