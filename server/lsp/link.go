package lsp

import (
	"context"
	"encoding/json"
	"log"

	"github.com/peske/lsp/protocol"
)

func (s *Server) documentLink(_ context.Context, params *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	log.Println("documentLink:")
	if d, err := json.MarshalIndent(&params, "", "  "); err == nil {
		log.Println(string(d))
	} else {
		log.Println(err)
	}
	return nil, nil
}
