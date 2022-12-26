package lsp

import (
	"context"

	"github.com/peske/lsp-srv/lsp/protocol"
	"go.uber.org/zap"
)

func (s *Server) initialize(_ context.Context, params *protocol.ParamInitialize) (*protocol.InitializeResult, error) {
	s.clientCapabilities = params.Capabilities

	return &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			CompletionProvider: protocol.CompletionOptions{
				ResolveProvider: false,
			},
			Workspace: protocol.Workspace6Gn{
				WorkspaceFolders: protocol.WorkspaceFolders5Gn{
					Supported: s.clientCapabilities.Workspace.WorkspaceFolders,
				},
			},
			HoverProvider: false,
		},
		ServerInfo: protocol.PServerInfoMsg_initialize{
			Name:    "lsp-sample-server",
			Version: "1.0.0",
		},
	}, nil
}

func (s *Server) initialized(ctx context.Context, params *protocol.InitializedParams) error {
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
			s.logger.Error("initialized", zap.Error(err))
			return err
		}
	}

	return nil
}
