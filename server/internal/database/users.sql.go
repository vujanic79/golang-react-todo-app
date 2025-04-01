// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, first_name, last_name, email)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, first_name, last_name, email
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Email     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FirstName,
		arg.LastName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.Email,
	)
	return i, err
}

const getUserIdByEmail = `-- name: GetUserIdByEmail :one
SELECT u.id FROM users u
WHERE u.email = $1
`

func (q *Queries) GetUserIdByEmail(ctx context.Context, email string) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getUserIdByEmail, email)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
