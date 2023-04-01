package post

import (
	"time"

	"github.com/google/uuid"
	"github.com/saufiroja/go-graphql-boilerplate/interfaces"
	"github.com/saufiroja/go-graphql-boilerplate/models/dto"
)

type PostService struct {
	repo interfaces.PostRepository
}

func NewPostService(repo interfaces.PostRepository) interfaces.PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) FindAllPost() ([]dto.FindAllPost, error) {
	return s.repo.FindAllPost()
}

func (s *PostService) CreatePost(input *dto.CreatePost) error {
	data := &dto.CreatePost{
		ID:        uuid.New().String(),
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	return s.repo.CreatePost(data)
}