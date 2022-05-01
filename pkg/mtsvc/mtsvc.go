package mtsvc

import (
	"context"
	"net"
	"strings"
)

type Core struct {
	NodeName string
	Peers    []string
}

type Peer struct {
	NodeName  string
	addresses []string
	Peers     map[string]Peer // peers of this peer
	Hop       string          // next hop to reach this peer
}

func NewCore() Core {
	// Roll a node name
	var nodeName string
	for {
		nodeName = randomString(4)
		if strings.Compare(nodeName, "0000") != 0 {
			break
		}
	}
	return Core{
		NodeName: nodeName,
	}
}

func (c *Core) NewLocalListener(port int, handler func()) {
	go c.startListener(port)
}

func (c *Core) startListener(ctx context.Context, port int) error {
	listener, err := net.Listen("tcp", ":"+string(port))
	if err != nil {
		return err
	}
	err = c.dispatch(listener, handler)
	if err != nil {
		return err
	}
	return nil
}

func (c *Core) dispatch(listener net.Listener, handler func()) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handler(conn)
	}
}
