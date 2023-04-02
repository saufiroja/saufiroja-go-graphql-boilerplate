package resolver

import "github.com/saufiroja/go-graphql-boilerplate/interfaces"

type Resolver struct {
	PostService interfaces.PostService
	AuthService interfaces.AuthService
}

func NewResolver(
	postService interfaces.PostService,
	AuthService interfaces.AuthService,
) *Resolver {
	return &Resolver{
		PostService: postService,
		AuthService: AuthService,
	}
}
