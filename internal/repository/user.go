package repository

import (
	"chat/internal/model"
	"context"

	"github.com/jackc/pgx/v5"
)

func IsDuplicateUser(ctx context.Context, user model.User) (bool, error) {
	query := "SELECT COUNT(1) FROM users WHERE (email = $1 OR username = $2);"
	var exists int
	err := DB.QueryRow(ctx, query, user.Email, user.Username).Scan(&exists)

	if exists == 0 {
		return false, err
	} else {
		return true, err
	}
}

func CreateUser(ctx context.Context, user model.User) error {
	query := "INSERT INTO users (email, username, password, created_at, updated_at)\nVALUES ($1, $2, $3, $4, $5);"
	_, err := DB.Exec(ctx, query, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)

	return err
}

func GetUser(ctx context.Context, lr model.LoginRequest) (model.User, error) {
	query := "SELECT * FROM users WHERE (email = $1 OR username = $2);"
	row, err := DB.Query(ctx, query, lr.Email, lr.Username)
	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[model.User])
	return user, err
}
