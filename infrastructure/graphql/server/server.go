package server

import (
	"net/http"

	"github.com/saufiroja/go-graphql-boilerplate/config"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/handler"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/middlewares"
)

func NewServer() *http.Server {
	app := http.NewServeMux()
	conf := config.NewAppConfig()

	gh := handler.NewHandler(conf)

	app.Handle("/auth", gh)
	app.Handle("/graphql", middlewares.AuthMiddleware(gh))

	server := &http.Server{
		Addr:    ":3000",
		Handler: app,
	}

	return server
}
