package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/go-graphql-boilerplate/config"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/handler"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func NewServer() *fiber.App {
	app := fiber.New()
	conf := config.NewAppConfig()

	gh := handler.NewHandler(app, conf)

	app.All("/graphql", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gh.ServeHTTP(w, r)
		})(c.Context())
		return nil
	})

	return app
}
