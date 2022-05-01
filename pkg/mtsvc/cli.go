package mtsvc

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
)

type Client struct {
	Core
	remoteUrl      string
	localAddresses []net.Addr
}

type PreparedStmt struct {
	stmt   int
	tokens []string
}

const (
	StmtNope      = 0x0000
	StmtUnknown   = 0x0001
	StmtListNodes = 0x0100
)

func NewClient(remoteUrl string) *Client {
	addrs, err := probeLocalAddresses()
	cli := &Client{
		Core:           NewCore(),
		remoteUrl:      remoteUrl,
		localAddresses: addrs,
	}
	err = cli.Connect()
	if err != nil {
		panic(err)
	}
	return cli
}

func (c *Client) Connect() error {
	url1, err := url.Parse(c.remoteUrl)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to parse remote url:", c.remoteUrl)
		return err
	}
	switch strings.ToLower(url1.Scheme) {

	case "http":
		panic("not implemented")
	default:
		// Use MTSTP
		panic("not implemented")
	}
}

func (c *Client) ExecuteMode(command string) {
	stmt, err := c.parseCommand(command)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Execute(stmt)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (c *Client) InteractiveMode() {
	for {
		cmd, err := c.readline()
		if err != nil {
			fmt.Println(err)
			continue
		}
		stmt, err := c.parseCommand(cmd)
		if err != nil {
			continue
		}
		err = c.Execute(stmt)
		fmt.Printf("\n")
	}
}

func (c *Client) readline() (string, error) {
	fmt.Print(">>>")
	return "", nil
}

func (c *Client) parseCommand(command string) (PreparedStmt, error) {
	command = strings.TrimSpace(command)
	quoted := false
	commands := strings.FieldsFunc(command, func(r rune) bool {
		if r == '"' {
			quoted = !quoted
		}
		return r == ' ' && !quoted
	})
	if len(commands) == 0 {
		return PreparedStmt{stmt: StmtNope}, nil
	}
	var stmt int
	var errStmt error
	switch strings.ToLower(commands[0]) {
	case "listnodes":
		stmt = StmtListNodes
	case "":
	default:
		stmt = StmtUnknown
	}
	return PreparedStmt{stmt: stmt}, errStmt
}

func (c *Client) Execute(stmt PreparedStmt) error {
	switch stmt.stmt {
	case StmtNope:
		return nil
	default:
		return fmt.Errorf("unknown statement")
	}
}
