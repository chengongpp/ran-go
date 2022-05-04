package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"ran/pkg/ran"
	"strconv"
	"strings"
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
	if len(os.Args) <= 1 {
		flag.Usage()
		os.Exit(1)
	}
	runMode := ClientMode
	if command == "" && !interactiveMode {
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
	url0, err := parseUrl(cPlaneUrl)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to parse cPlane url \"%s\": %v", cPlaneUrl, err)
		os.Exit(1)
	}
	switch strings.ToLower(url0.Scheme) {
	case "rantp", "":
		url0.Scheme = "rantp"
		svr, err := ran.NewServer(url0)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to create server: %v\n", err)
			os.Exit(1)
		}
		stmtCh, err := svr.Run()
		if remoteUrl != "" {

		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to run server: %v\n", err)
			os.Exit(1)
		}
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Unsupported protocol \"%s\"\n", url0.Scheme)
		os.Exit(1)
	}
}

// parseUrl is a wrapper of url.Parse, mainly to support input like ":8080"
func parseUrl(inputUrl string) (*url.URL, error) {
	var err error
	var url0 = &url.URL{}
	if !strings.Contains(inputUrl, "//") {
		url0.Scheme = "rantp"
		schm := strings.Split(inputUrl, ":")
		fmt.Println(schm)
		fmt.Println(len(schm))
		switch len(schm) {
		case 1:
			port, err := strconv.Atoi(schm[0])
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Failed to parse URL port from %s\n", inputUrl)
				return nil, err
			}
			if port < 1 || port > 65535 {
				_, _ = fmt.Fprintf(os.Stderr, "Invalid URL port %d\n", port)
				return nil, errors.New("invalid URL port")
			}
			url0.Host = "0.0.0.0:" + schm[0]
		case 2:
			if schm[0] == "" {
				url0.Host = "0.0.0.0:" + schm[1]
			}
		default:
			_, _ = fmt.Fprintf(os.Stderr, "Invalid URL %s\n", inputUrl)
			return nil, errors.New("invalid URL")
		}
	} else {
		url0, err = url.Parse(inputUrl)
	}
	return url0, err
}
