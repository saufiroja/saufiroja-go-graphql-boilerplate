package auth

import (
	"context"
	"database/sql"
	"time"

	"github.com/saufiroja/go-graphql-boilerplate/interfaces"
	"github.com/saufiroja/go-graphql-boilerplate/models/dto"
	"github.com/saufiroja/go-graphql-boilerplate/utils"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) interfaces.AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(input *dto.Register) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer utils.Transaction(r.db)

	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1 ,$2 , $3 , $4 , $5, $6)",
		input.ID, input.Name, input.Email, input.Password, input.CreatedAt, input.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
