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

func (r *postRepository) FindPostById(id string) (dto.FindPostById, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var post dto.FindPostById

	rows := r.db.QueryRowContext(ctx, "SELECT id, title, content, created_at, updated_at FROM posts WHERE id = $1", id)

	err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *postRepository) DeletePostById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	defer utils.Transaction(r.db)

	_, err := r.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *postRepository) UpdatePostById(id string, input *dto.UpdatePost) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	defer utils.Transaction(r.db)

	_, err := r.db.ExecContext(
		ctx,
		"UPDATE posts SET title = $1, content = $2, updated_at = $3 WHERE id = $4",
		input.Title, input.Content, input.UpdatedAt, id,
	)
	if err != nil {
		return err
	}

	return nil
}
