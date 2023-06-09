package schema

import "github.com/graphql-go/graphql"

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.Int,
		},
		"updated_at": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var PostInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PostInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"content": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

func (s *Schema) FindAllPost() *graphql.Field {
	field := &graphql.Field{
		Type:        graphql.NewList(PostType),
		Description: "Find all post",
		Resolve:     s.resolver.FindAllPost,
	}

	return field
}

func (s *Schema) CreatePost() *graphql.Field {
	field := &graphql.Field{
		Type:        PostType,
		Description: "Create post",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(PostInputType),
			},
		},
		Resolve: s.resolver.CreatePost,
	}

	return field
}

func (s *Schema) FindPostById() *graphql.Field {
	field := &graphql.Field{
		Type:        PostType,
		Description: "Find post by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: s.resolver.FindPostById,
	}

	return field
}

func (s *Schema) DeletePostById() *graphql.Field {
	field := &graphql.Field{
		Type:        PostType,
		Description: "Delete post by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: s.resolver.DeletePostById,
	}

	return field
}

func (s *Schema) UpdatePostById() *graphql.Field {
	field := &graphql.Field{
		Type:        PostType,
		Description: "Update post by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(PostInputType),
			},
		},
		Resolve: s.resolver.UpdatePostById,
	}

	return field
}
