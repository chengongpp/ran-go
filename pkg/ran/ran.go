package ran

import (
	"context"
	"net"
	"strconv"
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

func (c *Core) NewLocalListener(port int, handler func(conn net.Conn)) {
	go func() {
		err := c.startListener(context.Background(), port)
		if err != nil {

		}
	}()
}

func (c *Core) startListener(ctx context.Context, port int) error {
	panic("not implemented")
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	err = c.dispatch(listener, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Core) dispatch(listener net.Listener, handler func(conn net.Conn)) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handler(conn)
	}
}
