package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/peske/lsp/protocol"
	"github.com/peske/x-tools-internal/jsonrpc2"
)

func (s *Server) initialize(_ context.Context, params *protocol.ParamInitialize) (*protocol.InitializeResult, error) {
	s.mu.Lock()
	if s.state >= serverInitializing {
		defer s.mu.Unlock()
		return nil, fmt.Errorf("%w: initialize called while server in %v state", jsonrpc2.ErrInvalidRequest, s.state)
	}
	s.state = serverInitializing
	s.mu.Unlock()

	log.Println("initialize:")
	if d, err := json.MarshalIndent(&params, "", "  "); err == nil {
		log.Println(string(d))
	} else {
		log.Println(err)
	}

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

func (s *Server) initialized(ctx context.Context, params *protocol.InitializedParams) error {
	s.mu.Lock()
	if s.state >= serverInitialized {
		defer s.mu.Unlock()
		return fmt.Errorf("%w: initialized called while server in %v state", jsonrpc2.ErrInvalidRequest, s.state)
	}
	s.state = serverInitialized
	s.mu.Unlock()

	log.Println("initialized:")
	if d, err := json.MarshalIndent(&params, "", "  "); err == nil {
		log.Println(string(d))
	} else {
		log.Println(err)
	}

	//TODO(peske): Ensure that this implementation is equivalent to the original.
	var rs []protocol.Registration
	if s.clientCapabilities.Workspace.Configuration {
		rs = append(rs, protocol.Registration{
			ID:     "workspace/didChangeConfiguration",
			Method: "workspace/didChangeConfiguration",
		})
	}
	if s.clientCapabilities.Workspace.WorkspaceFolders {
		rs = append(rs, protocol.Registration{
			ID:     "workspace/didChangeWorkspaceFolders",
			Method: "workspace/didChangeWorkspaceFolders",
		})
	}

	if len(rs) > 0 {
		if err := s.client.RegisterCapability(ctx, &protocol.RegistrationParams{
			Registrations: rs,
		}); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (s *Server) shutdown(_ context.Context) error {
	log.Println("shutdown")
	return nil
}
