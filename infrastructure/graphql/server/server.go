package server

import (
	"net/http"

	"github.com/saufiroja/go-graphql-boilerplate/config"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/handler"
)

func NewServer() *http.Server {
	app := http.NewServeMux()
	conf := config.NewAppConfig()

	// // auth middleware
	// app.Use(middlewares.AuthMiddleware)

	gh := handler.NewHandler(conf)

	app.Handle("/graphql", gh)

	server := &http.Server{
		Addr:    ":3000",
		Handler: app,
	}

	return server
}
