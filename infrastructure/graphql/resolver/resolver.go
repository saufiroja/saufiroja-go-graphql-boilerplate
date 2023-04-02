package resolver

import "github.com/saufiroja/go-graphql-boilerplate/interfaces"

type Resolver struct {
	PostService interfaces.PostService
	AuthService interfaces.AuthService
	UserService interfaces.UserService
}

func NewResolver(
	postService interfaces.PostService,
	AuthService interfaces.AuthService,
	UserService interfaces.UserService,
) *Resolver {
	return &Resolver{
		PostService: postService,
		AuthService: AuthService,
		UserService: UserService,
	}
}
