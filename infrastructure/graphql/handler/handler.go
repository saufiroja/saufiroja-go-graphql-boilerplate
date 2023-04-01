package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/handler"
	"github.com/saufiroja/go-graphql-boilerplate/config"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/database/postgres"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/resolver"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/schema"
	repoPost "github.com/saufiroja/go-graphql-boilerplate/repository/post"
	svcPost "github.com/saufiroja/go-graphql-boilerplate/service/post"
)

func NewHandler(app *fiber.App, conf *config.AppConfig) *handler.Handler {
	db := postgres.NewPostgres(conf)

	// repository
	postRepo := repoPost.NewPostRepository(db)

	// service
	postService := svcPost.NewPostService(postRepo)

	// resolver
	resolver := resolver.NewResolver(postService)

	// schema
	schema := schema.NewSchema(resolver)

	// graphql handler
	gh := handler.New(&handler.Config{
		Schema:   schema.RootSchema(),
		Pretty:   true,
		GraphiQL: true,
	})

	return gh
}
