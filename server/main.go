package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/peske/lsp-example/server/lsp"
	lsp_srv_ex "github.com/peske/lsp-srv-ex"
	"go.uber.org/zap"

	. "github.com/peske/lsp-example/server/logging"
)

func main() {
	defer func() {
		_ = Logger.Sync()
	}()
	Logger.Info("Starting...")

	// Listen for SIGINT, just to check if/when the server process gets killed:
	c1 := make(chan os.Signal, 1)
	c2 := make(chan bool, 1)
	signal.Notify(c1, os.Interrupt, os.Kill)
	go func() {
		select {
		case sig := <-c1:
			Logger.Warn(fmt.Sprintf("os.Signal received: %v", sig))
		case <-c2:
			break
		}
	}()

	// Launching the server:
	cfg := &lsp_srv_ex.Config{Caching: true}
	if err := lsp_srv_ex.Run(lsp.NewServer, cfg, Logger); err != nil {
		Logger.Error("lsp_srv_ex.Run", zap.Error(err))
	} else {
		Logger.Info("Finished")
	}

	// Exit SIGINT listener:
	c2 <- true
}
