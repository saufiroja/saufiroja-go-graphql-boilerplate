package interfaces

import "github.com/saufiroja/go-graphql-boilerplate/models/dto"

type UserRepository interface {
	FindUserById(id string) (dto.FindUserById, error)
	UpdateUserById(id string, data *dto.UpdateUserById) error
	DeleteUserById(id string) error
}

type UserService interface {
	FindUserById(id string) (dto.FindUserById, error)
	UpdateUserById(id string, data *dto.UpdateUserById) error
	DeleteUserById(id string) error
}
