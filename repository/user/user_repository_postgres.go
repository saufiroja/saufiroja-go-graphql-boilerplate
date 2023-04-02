package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/saufiroja/go-graphql-boilerplate/interfaces"
	"github.com/saufiroja/go-graphql-boilerplate/models/dto"
	"github.com/saufiroja/go-graphql-boilerplate/utils"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindUserById(id string) (dto.FindUserById, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user dto.FindUserById
	rows := r.db.QueryRowContext(
		ctx,
		"SELECT id, name, email, created_at FROM users WHERE id = $1",
		id,
	)

	err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateUserById(id string, data *dto.UpdateUserById) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer utils.Transaction(r.db)

	_, err := r.db.ExecContext(
		ctx,
		"UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4",
		data.Name, data.Email, data.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUserById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer utils.Transaction(r.db)

	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
