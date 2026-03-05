package repository

import (
	"context"
	"time"
	"todo_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(pool *pgxpool.Pool, user *models.Users) (*models.Users, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
		INSERT INTO users (email,password)
		VALUES($1,$2)
		RETURNING id,email,created_at,updated_at
	`

	err := pool.QueryRow(ctx, query, user.Email, user.Password).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(pool *pgxpool.Pool, email string) (*models.Users, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
		SELECT id,email,password,created_at,updated_at
		FROM users
		WHERE email = $1
	`

	var users models.Users

	err := pool.QueryRow(ctx, query, email).Scan(
		&users.ID,
		&users.Email,
		&users.Password,
		&users.CreatedAt,
		&users.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func GetUserByID(pool *pgxpool.Pool, id string) (*models.Users, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
		SELECT id,email,password,created_at,updated_at
		FROM users
		WHERE id = $1
	`

	var user models.Users

	err := pool.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
