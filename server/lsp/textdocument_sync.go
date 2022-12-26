package lsp

import (
	"context"

	"github.com/peske/lsp-srv/lsp/protocol"
)

func (s *Server) didChange(_ context.Context, params *protocol.DidChangeTextDocumentParams) error {
	s.checkDiagnosticsThread(params.TextDocument.URI)
	return nil
}

func (s *Server) didOpen(_ context.Context, params *protocol.DidOpenTextDocumentParams) error {
	s.checkDiagnosticsThread(params.TextDocument.URI)
	return nil
}
