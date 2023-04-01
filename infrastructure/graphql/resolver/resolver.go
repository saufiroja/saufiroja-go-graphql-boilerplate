package resolver

import "github.com/saufiroja/go-graphql-boilerplate/interfaces"

type Resolver struct {
	PostService interfaces.PostService
}

func NewResolver(postService interfaces.PostService) *Resolver {
	return &Resolver{
		PostService: postService,
	}
}
