package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/saufiroja/go-graphql-boilerplate/models/dto"
)

func (r *Resolver) FindAllPost(params graphql.ResolveParams) (interface{}, error) {
	return r.PostService.FindAllPost()
}

func (r *Resolver) CreatePost(params graphql.ResolveParams) (interface{}, error) {
	input, _ := params.Args["input"].(map[string]interface{})
	title := input["title"].(string)
	content := input["content"].(string)

	data := &dto.CreatePost{
		Title:   title,
		Content: content,
	}

	err := r.PostService.CreatePost(data)
	if err != nil {
		return nil, err
	}

	return input, nil
}
