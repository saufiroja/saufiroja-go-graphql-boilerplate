package resolver

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/graphql-go/graphql"
	"github.com/saufiroja/go-graphql-boilerplate/models/dto"
)

func (r *Resolver) Register(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	user := &dto.Register{
		Name:     input["name"].(string),
		Email:    input["email"].(string),
		Password: input["password"].(string),
	}

	err := r.AuthService.Register(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Resolver) Login(params graphql.ResolveParams) (interface{}, error) {
	input := params.Args["input"].(map[string]interface{})
	user := &dto.Login{
		Email:    input["email"].(string),
		Password: input["password"].(string),
	}

	token, err := r.AuthService.Login(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (r *Resolver) FindUserById(params graphql.ResolveParams) (interface{}, error) {
	_ = params.Context.Value("email").(*jwt.Token)

	id := params.Args["id"].(string)

	user, err := r.UserService.FindUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Resolver) UpdateUserById(params graphql.ResolveParams) (interface{}, error) {
	_ = params.Context.Value("email").(*jwt.Token)

	id := params.Args["id"].(string)
	input := params.Args["input"].(map[string]interface{})
	user := &dto.UpdateUserById{
		Name:  input["name"].(string),
		Email: input["email"].(string),
	}

	err := r.UserService.UpdateUserById(id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Resolver) DeleteUserById(params graphql.ResolveParams) (interface{}, error) {
	_ = params.Context.Value("email").(*jwt.Token)

	id := params.Args["id"].(string)

	err := r.UserService.DeleteUserById(id)
	if err != nil {
		return nil, err
	}

	return id, nil
}
