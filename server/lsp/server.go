package lsp

import (
	"context"
	"fmt"
	lsp_srv_ex "github.com/peske/lsp-srv-ex"
	"github.com/peske/lsp-srv/lsp/protocol"
	"go.uber.org/zap"
)

// Server is the type mentioned in the instructions above.
type Server struct {
	client protocol.ClientCloser
	helper *lsp_srv_ex.Helper
	ctx    context.Context
	cancel func()

	clientCapabilities protocol.ClientCapabilities

	logger *zap.Logger
}

// NewServer is the factory function mentioned in the instructions above.
func NewServer(client protocol.ClientCloser, ctx context.Context, cancel func(), helper *lsp_srv_ex.Helper) protocol.Server {
	return &Server{
		client: client,
		helper: helper,
		ctx:    ctx,
		cancel: cancel,
	}
}

func notImplemented(method string) error {
	return fmt.Errorf("'%s' not implemented", method)
}

//go:generate ../lspgen -o server_gen.go -t Server -u .
