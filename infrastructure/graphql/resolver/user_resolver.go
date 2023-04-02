package resolver

import (
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
