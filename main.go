package main

import (
	"flag"
	"sync"

	"github.com/nut-game/nano/client"
	"github.com/nut-game/nano/session"
)

var (
	pClient        client.NanoClient
	disconnectedCh chan bool
	docsString     string
	fileName       string
	pushInfo       map[string]string
	wait           sync.WaitGroup
	prettyJSON     bool
	handshake      *session.HandshakeData
)

func main() {
	flag.StringVar(&docsString, "docs", "", "documentation route")
	flag.StringVar(&fileName, "filename", "", "file with commands")
	flag.BoolVar(&prettyJSON, "pretty", false, "print pretty jsons")
	flag.Parse()
	handshake = &session.HandshakeData{
		Sys: session.HandshakeClientData{
			Platform:    "mac",
			LibVersion:  "0.3.5-release",
			BuildNumber: "20",
			Version:     "1.0.0",
		},
		User: map[string]interface{}{
			"age": 30,
		},
	}

	switch {
	case fileName != "":
		executeFromFile(fileName)
	default:
		repl()
	}
}
