package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/server"
)

func main() {
	app := server.NewServer()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := app.Listen(":3000")
		if err != nil {
			panic(err)
		}
	}()
	<-done
	log.Println("Shutting down server...")

	err := app.Shutdown()
	if err != nil {
		panic(err)
	}

	log.Println("Server gracefully stopped")
}
