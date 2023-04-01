package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/saufiroja/go-graphql-boilerplate/infrastructure/graphql/resolver"
)

type Schema struct {
	resolver *resolver.Resolver
}

func NewSchema(resolver *resolver.Resolver) *Schema {
	return &Schema{
		resolver: resolver,
	}
}

func (s *Schema) Query() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"findAllPost":  s.FindAllPost(),
			"findPostById": s.FindPostById(),
		},
	})
}

func (s *Schema) Mutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createPost": s.CreatePost(),
			"deletePost": s.DeletePostById(),
		},
	})
}

func (s *Schema) RootSchema() *graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    s.Query(),
		Mutation: s.Mutation(),
	})

	if err != nil {
		panic(err)
	}

	return &schema
}
