package interfaces

import "github.com/saufiroja/go-graphql-boilerplate/models/dto"

type AuthRepository interface {
	Register(input *dto.Register) error
	Login(input *dto.Login) (dto.Login, error)
}

type AuthService interface {
	Register(input *dto.Register) error
	Login(input *dto.Login) (*dto.Token, error)
}
