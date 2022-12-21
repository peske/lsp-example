package lsp

import (
	"fmt"
	"log"
)

type Server struct{}

func notImplemented(method string) error {
	log.Printf("'%s' method call")
	return fmt.Errorf("'%s' not implemented", method)
}

//go:generate ../lspgen -o server_gen.go -u .
