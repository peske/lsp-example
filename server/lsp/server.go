package lsp

import (
	"fmt"
	"github.com/peske/lsp/protocol"
	"log"
	"sync"
)

type serverState int

const (
	serverCreated      = serverState(iota)
	serverInitializing // set once the server has received "initialize" request
	serverInitialized  // set once the server has received "initialized" request
	serverShutDown
)

type Server struct {
	clientCapabilities protocol.ClientCapabilities

	mu    sync.Mutex
	state serverState
}

func notImplemented(method string) error {
	log.Printf("'%s' method call", method)
	return fmt.Errorf("'%s' not implemented", method)
}

//go:generate ../lspgen -o server_gen.go -u .
