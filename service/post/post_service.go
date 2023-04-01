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

func (s *PostService) FindPostById(id string) (dto.FindPostById, error) {
	return s.repo.FindPostById(id)
}

func (s *PostService) DeletePostById(id string) error {
	return s.repo.DeletePostById(id)
}

func (s *PostService) UpdatePostById(id string, input *dto.UpdatePost) error {
	data := &dto.UpdatePost{
		Title:     input.Title,
		Content:   input.Content,
		UpdatedAt: time.Now().Unix(),
	}

	return s.repo.UpdatePostById(id, data)
}
