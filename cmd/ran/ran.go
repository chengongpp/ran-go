package main

import (
	"flag"
	"fmt"
	"os"
	"ran/pkg/ran"
)

const (
	ClientMode = iota
	ServerMode
)

var verbose bool

func main() {
	var cPlaneUrl string
	var remoteUrl string
	var command string
	var interactiveMode bool
	flag.StringVar(&cPlaneUrl, "l", "", "CPlane listener URL setup")
	flag.StringVar(&remoteUrl, "u", "", "Remote endpoint URL setup")
	flag.StringVar(&command, "c", "", "Command to execute")
	flag.BoolVar(&interactiveMode, "i", false, "Interactive mode")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.Parse()
	runMode := ClientMode
	if command == "" && interactiveMode {
		runMode = ServerMode
	}
	switch runMode {
	case ClientMode:
		runClient(remoteUrl, command, interactiveMode)
	default:
		runServer(remoteUrl, cPlaneUrl)
	}
}

func runClient(remoteUrl string, command string, interactiveMode bool) {
	if remoteUrl == "" {
		_, _ = fmt.Fprintln(os.Stderr, "Remote URL is not set")
		os.Exit(1)
	}
	if command != "" && interactiveMode {
		_, _ = fmt.Fprintln(os.Stderr, "Will not run in interactive mode with command")
	}
	client := ran.NewClient(remoteUrl)
	if command != "" {
		client.ExecuteMode(command)
	} else {
		client.InteractiveMode()
	}
}

func runServer(remoteUrl string, cPlaneUrl string) {

}
