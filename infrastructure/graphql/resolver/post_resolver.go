package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/saufiroja/go-graphql-boilerplate/models/dto"
)

func (r *Resolver) FindAllPost(params graphql.ResolveParams) (interface{}, error) {
	return r.PostService.FindAllPost()
}

func (r *Resolver) CreatePost(params graphql.ResolveParams) (interface{}, error) {
	input := &dto.CreatePost{
		Title:   params.Args["title"].(string),
		Content: params.Args["content"].(string),
	}

	err := r.PostService.CreatePost(input)
	if err != nil {
		return nil, err
	}

	return input, nil
}
