package main

import (
	"context"
	"log"
	"os"

	"github.com/peske/lsp-example/server/lsp"
	"github.com/peske/lsp/cmd"
)

func main() {
	// Configure logging
	f, err := os.OpenFile("/tmp/mend-ls.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func() {
		log.Println("exiting")
		_ = f.Close()
	}()
	log.SetOutput(f)
	log.Println("mend-ls start")

	serve := &cmd.Serve{Server: &lsp.Server{}}
	if err = serve.Run(context.Background()); err != nil {
		log.Println(err)
	}
}
