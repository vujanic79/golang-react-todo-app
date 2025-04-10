package domain

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
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
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
}

func (cup CreateUserParams) MarshalZerologObject(event *zerolog.Event) {
	event.
		Str("first_name", cup.FirstName).
		Str("last_name", cup.LastName).
		Str("email", cup.Email)
}

type UserService interface {
	CreateUser(ctx context.Context, createUserParams CreateUserParams) (user User, err error)
	GetUserIdByEmail(ctx context.Context, email string) (userId uuid.UUID, err error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, createUserParams CreateUserParams) (user User, err error)
	GetUserIdByEmail(ctx context.Context, email string) (userId uuid.UUID, err error)
}

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
