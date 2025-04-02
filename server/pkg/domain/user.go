package domain

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
}

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Email     string
}

type UserService interface {
	CreateUser(createUserParams CreateUserParams) (user User, err error)
	GetUserIdByEmail(email string) (userId uuid.UUID, err error)
	SetContext(ctx context.Context)
}

type UserRepository interface {
	CreateUser(createUserParams CreateUserParams) (user User, err error)
	GetUserIdByEmail(email string) (userId uuid.UUID, err error)
	SetContext(ctx context.Context)
}

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
