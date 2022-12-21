package lsp

import (
	"context"
	"fmt"
	"log"

	"github.com/peske/lsp/protocol"
	"github.com/peske/x-tools-internal/jsonrpc2"
)

func (s *Server) initialize(ctx context.Context, params *protocol.ParamInitialize) (*protocol.InitializeResult, error) {
	s.mu.Lock()
	if s.state >= serverInitializing {
		defer s.mu.Unlock()
		return nil, fmt.Errorf("%w: initialize called while server in %v state", jsonrpc2.ErrInvalidRequest, s.state)
	}
	s.state = serverInitializing
	s.mu.Unlock()

	log.Printf("initialize; Client capabilities:\n%v", params.Capabilities)

	s.clientCapabilities = params.Capabilities

	return &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: &protocol.TextDocumentSyncOptions{
				Change: protocol.Incremental,
			},
			CompletionProvider: protocol.CompletionOptions{
				ResolveProvider: true,
			},
			Workspace: protocol.Workspace6Gn{
				WorkspaceFolders: protocol.WorkspaceFolders5Gn{
					Supported: s.clientCapabilities.Workspace.WorkspaceFolders,
				},
			},
		},
		ServerInfo: protocol.PServerInfoMsg_initialize{
			Name:    "lsp-sample-server",
			Version: "1.0.0",
		},
	}, nil
}
