package interfaces

import "github.com/saufiroja/go-graphql-boilerplate/models/dto"

type AuthRepository interface {
	Register(input *dto.Register) error
}

type AuthService interface {
	Register(input *dto.Register) error
}
