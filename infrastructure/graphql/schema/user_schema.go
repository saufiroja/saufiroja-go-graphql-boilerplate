package schema

import "github.com/graphql-go/graphql"

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
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

var LoginType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Login",
	Fields: graphql.Fields{
		"access_token": &graphql.Field{
			Type: graphql.String,
		},
		"refresh_token": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UpdateUserInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateUserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var UserTypeInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var LoginTypeInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "LoginInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

func (s *Schema) Register() *graphql.Field {
	field := &graphql.Field{
		Type:        UserType,
		Description: "Register user",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(UserTypeInput),
			},
		},
		Resolve: s.resolver.Register,
	}

	return field
}

func (s *Schema) Login() *graphql.Field {
	field := &graphql.Field{
		Type:        LoginType,
		Description: "Login user",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(LoginTypeInput),
			},
		},
		Resolve: s.resolver.Login,
	}

	return field
}

func (s *Schema) FindUserById() *graphql.Field {
	field := &graphql.Field{
		Type:        UserType,
		Description: "Find user by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: s.resolver.FindUserById,
	}

	return field
}

func (s *Schema) UpdateUserById() *graphql.Field {
	field := &graphql.Field{
		Type:        UserType,
		Description: "Update user by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(UpdateUserInput),
			},
		},
		Resolve: s.resolver.UpdateUserById,
	}

	return field
}

func (s *Schema) DeleteUserById() *graphql.Field {
	field := &graphql.Field{
		Type:        UserType,
		Description: "Delete user by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: s.resolver.DeleteUserById,
	}

	return field
}
