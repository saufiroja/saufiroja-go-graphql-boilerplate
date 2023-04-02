package resolver

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/saufiroja/go-graphql-boilerplate/models/dto"
	"github.com/saufiroja/go-graphql-boilerplate/utils"
)

func (r *Resolver) FindAllPost(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("email").(utils.Token)
	fmt.Println(token)
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

func (r *Resolver) FindPostById(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)

	return r.PostService.FindPostById(id)
}

func (r *Resolver) DeletePostById(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)

	err := r.PostService.DeletePostById(id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (r *Resolver) UpdatePostById(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	input, _ := params.Args["input"].(map[string]interface{})
	title := input["title"].(string)
	content := input["content"].(string)

	data := &dto.UpdatePost{
		Title:   title,
		Content: content,
	}

	err := r.PostService.UpdatePostById(id, data)
	if err != nil {
		return nil, err
	}

	return input, nil
}
