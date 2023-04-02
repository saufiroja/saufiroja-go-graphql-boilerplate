package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/go-graphql-boilerplate/config"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/handler"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/middlewares"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func NewServer() *fiber.App {
	app := fiber.New()
	conf := config.NewAppConfig()

	// middlewares
	app.Use(middlewares.Logger(app))
	app.Use(middlewares.Limiters(app))
	// // auth middleware
	// app.Use(middlewares.AuthMiddleware)

	gh := handler.NewHandler(app, conf)

	// graphql endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Welcome to GraphQL",
		})
	})

	app.All("/graphql", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gh.ServeHTTP(w, r)
		})(c.Context())
		return nil
	})

	return app
}
