package auth

import (
	"time"

	"github.com/google/uuid"
	"github.com/saufiroja/go-graphql-boilerplate/interfaces"
	"github.com/saufiroja/go-graphql-boilerplate/models/dto"
	"github.com/saufiroja/go-graphql-boilerplate/utils"
)

type authService struct {
	repo interfaces.AuthRepository
}

func NewAuthService(repo interfaces.AuthRepository) interfaces.AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(input *dto.Register) error {
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return err
	}

	data := &dto.Register{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  hash,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	return s.repo.Register(data)
}
