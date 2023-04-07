package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/server"
)

func main() {
	app := server.NewServer()

	go func() {
		err := app.ListenAndServe()
		if err != nil {
			log.Fatalf("Server error: %v", err)
		}

		log.Println("Stopped serving new connections")
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := app.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	log.Println("Server gracefully stopped")
}
