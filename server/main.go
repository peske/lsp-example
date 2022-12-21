package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/peske/lsp-example/server/lsp"
	"github.com/peske/lsp/cmd"
)

func main() {
	// Configure logging:
	f, err := os.OpenFile("/tmp/lsp-example.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func() {
		log.Println("exiting")
		_ = f.Close()
	}()
	log.SetOutput(f)
	log.Println("starting...")

	// Listen for SIGINT, just to check if/when the server process gets killed:
	c1 := make(chan os.Signal, 1)
	c2 := make(chan bool, 1)
	signal.Notify(c1, os.Interrupt, os.Kill)
	go func() {
		select {
		case sig := <-c1:
			log.Printf("os.Signal received: %v", sig)
		case <-c2:
			break
		}
	}()

	// Launching the server:
	serve := &cmd.Serve{Server: &lsp.Server{}}
	if err = serve.Run(context.Background()); err != nil {
		log.Println(err)
	} else {
		log.Println("stopped")
	}

	// Exit SIGINT listener:
	c2 <- true
}
