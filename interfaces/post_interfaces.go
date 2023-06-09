package interfaces

import "github.com/saufiroja/go-graphql-boilerplate/models/dto"

type PostRepository interface {
	FindAllPost() ([]dto.FindAllPost, error)
	FindPostById(id string) (dto.FindPostById, error)
	CreatePost(input *dto.CreatePost) error
	DeletePostById(id string) error
	UpdatePostById(id string, input *dto.UpdatePost) error
}

type PostService interface {
	FindAllPost() ([]dto.FindAllPost, error)
	CreatePost(input *dto.CreatePost) error
	FindPostById(id string) (dto.FindPostById, error)
	DeletePostById(id string) error
	UpdatePostById(id string, input *dto.UpdatePost) error
}
