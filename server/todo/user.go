package todo

import (
	"context"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/internal/database"
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

type ReqParamsForCreateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UserService interface {
	CreateUser(ctx context.Context, params ReqParamsForCreateUser, DB *database.Queries) (dbUser database.User, err error)
	GetUserIdByEmail(ctx context.Context, email string, DB *database.Queries) (dbUserId uuid.UUID, err error)
}

type UserController interface {
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
}

func MapDbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
	}
}
