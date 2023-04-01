package post

import (
	"context"
	"database/sql"
	"time"

	"github.com/saufiroja/go-graphql-boilerplate/interfaces"
	"github.com/saufiroja/go-graphql-boilerplate/models/dto"
	"github.com/saufiroja/go-graphql-boilerplate/utils"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) interfaces.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) FindAllPost() ([]dto.FindAllPost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var posts []dto.FindAllPost

	rows, err := r.db.QueryContext(ctx, "SELECT id, title, content, created_at, updated_at FROM posts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post dto.FindAllPost

		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *postRepository) CreatePost(input *dto.CreatePost) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	defer utils.Transaction(r.db)

	_, err := r.db.ExecContext(ctx, "INSERT INTO posts (id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", input.ID, input.Title, input.Content, input.CreatedAt, input.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
