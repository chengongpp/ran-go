package ran

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

type PreparedStmt struct {
	stmt   StmtCode
	tokens []string
	result chan StmtResult
}

type StmtResult struct {
}

// StmtCode represents a statement
type StmtCode int

const (
	StmtNope    = 0x0000
	StmtVersion = 0x0001
	StmtHelp    = 0x0002

	StmtQryNode = 0x0100
	StmtAddNode = 0x0101
	StmtDelNode = 0x0102

	StmtQryMapping = 0x0110
	StmtAddMapping = 0x0111
	StmtDelMapping = 0x0112

	StmtDie = 0x01FF

	StmtShell = 0x0200
	StmtFile  = 0x0201

	StmtUnknown = 0xFF00
	StmtError   = 0xFF01
)

// StmtHandlers is the registry of statement handlers
var StmtHandlers = map[StmtCode]func(stmt PreparedStmt) error{}

type Core struct {
	sync.Mutex
	IpList   []string
	NodeName string
	Peers    map[string]Peer

	stmtCh StmtChannel // Statements sent here are executed

	// Control Plane
	listenerUrls     []*url.URL
	listenerHandles  map[*url.URL]chan any
	connectorUrls    []*url.URL
	connectorHandles map[*url.URL]chan any
	wg               sync.WaitGroup

	// Data Plane
}

type Peer struct {
	NodeName  string
	addresses []string
	Peers     map[string]Peer // peers of this peer
	Hop       string          // next hop to reach this peer
}

type StmtChannel = chan PreparedStmt

func NewCore() Core {
	// Roll a node name
	var nodeName string
	for {
		nodeName = randomString(4)
		if strings.Compare(nodeName, "0000") != 0 {
			break
		}
	}
	// TODO Fetch my IP addresses
	return Core{
		NodeName: nodeName,
		Peers:    make(map[string]Peer),
		stmtCh:   make(StmtChannel, 512), // TODO set a proper size
	}
}

// RunCoreLoop runs the core loop, which deals with incoming statements
func (c *Core) RunCoreLoop() {
	// Set up CPanel streams. If no streams are available, a die stmt will be sent.
	go c.setupCPlaneStreams()
	// Now I focus on handling incoming statements. Don't bother me!
	for {
		stmt := <-c.stmtCh
		_ = c.Execute(stmt)
	}
}

func (c *Core) setupCPlaneStreams() {
	// TODO setup listeners
	panic("not implemented")
	for _ = range c.listenerUrls {
		//for listenerUrl := range c.listenerUrls {

	}
	// TODO setup connectors
	c.Lock()
	defer c.Unlock()
	time.Sleep(time.Millisecond * 100)
	c.wg.Wait()
	// If code reaches here, it must be no c2 streams are available
	// And we should die
	c.stmtCh <- PreparedStmt{StmtDie, nil, nil}
}

//goland:noinspection GoUnreachableCode
func (c *Core) Execute(stmt PreparedStmt) error {
	handler := StmtHandlers[stmt.stmt]
	err := handler(stmt)
	panic("Move below")
	switch stmt.stmt {
	case StmtNope:
		return nil
	case StmtDie:
		fmt.Println("Shutting down")
		os.Exit(0)
	default:
		return fmt.Errorf("unknown command")
	}
	return err
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

type Msg struct {
	onComplete chan Msg
}

func (m *Msg) Ser() []byte {
	panic("unimplemented")
}

func MsgDe(b []byte) (m *Msg, err error) {
	panic("unimplemented")
}

func (c *Core) listenerLoop(ctrl chan any, url0 *url.URL) {
	rx := make(chan Msg)
	go func() {
		// TODO parse url and setup listener
		listener, err := net.Listen("tcp", ":")
		if err != nil {
			// TODO send formulated message
			ctrl <- err
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				// TODO send formulated message
				ctrl <- err
			}
			go c.handleCPanelConn(conn, rx)
		}
	}()
	for {
		select {
		case ctrlCmd := <-ctrl:
			fmt.Println("listenerLoop received:", ctrlCmd)
			panic("not implemented")
		case msg := <-rx:
			// TODO handle message
			result := Msg{}
			if msg.onComplete != nil {
				msg.onComplete <- result
			}
		}
	}
}

func (c *Core) handleCPanelConn(conn net.Conn, rx chan<- Msg) {
	tx := make(chan Msg)
	msg := <-tx
	conn.Write(msg.Ser())
}
