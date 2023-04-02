package user

import (
	"time"

	"github.com/saufiroja/go-graphql-boilerplate/interfaces"
	"github.com/saufiroja/go-graphql-boilerplate/models/dto"
)

type userService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) FindUserById(id string) (dto.FindUserById, error) {
	return s.repo.FindUserById(id)
}

func (s *userService) UpdateUserById(id string, data *dto.UpdateUserById) error {
	input := &dto.UpdateUserById{
		Name:      data.Name,
		Email:     data.Email,
		UpdatedAt: time.Now().Unix(),
	}

	return s.repo.UpdateUserById(id, input)
}

func (s *userService) DeleteUserById(id string) error {
	return s.repo.DeleteUserById(id)
}
