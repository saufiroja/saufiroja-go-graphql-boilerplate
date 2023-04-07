package handler

import (
	"github.com/graphql-go/handler"
	"github.com/saufiroja/go-graphql-boilerplate/config"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/database/postgres"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/resolver"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/schema"
	repoAuth "github.com/saufiroja/go-graphql-boilerplate/repository/auth"
	repoPost "github.com/saufiroja/go-graphql-boilerplate/repository/post"
	repoUser "github.com/saufiroja/go-graphql-boilerplate/repository/user"
	svcAuth "github.com/saufiroja/go-graphql-boilerplate/service/auth"
	svcPost "github.com/saufiroja/go-graphql-boilerplate/service/post"
	svcUser "github.com/saufiroja/go-graphql-boilerplate/service/user"
)

func NewHandler(conf *config.AppConfig) *handler.Handler {
	db := postgres.NewPostgres(conf)

	// repository
	postRepo := repoPost.NewPostRepository(db)
	authRepo := repoAuth.NewAuthRepository(db)
	userRepo := repoUser.NewUserRepository(db)

	// service
	postService := svcPost.NewPostService(postRepo)
	authService := svcAuth.NewAuthService(authRepo)
	userService := svcUser.NewUserService(userRepo)

	// resolver
	resolver := resolver.NewResolver(
		postService,
		authService,
		userService,
	)

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
