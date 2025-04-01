package sql

import (
	"context"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/internal/database"
	"github.com/vujanic79/golang-react-todo-app/todo"
	"time"
)

var _ todo.UserService = (*UserService)(nil)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (usrService *UserService) CreateUser(ctx context.Context,
	params todo.ReqParamsForCreateUser,
	DB *database.Queries) (dbUser database.User, err error) {
	dbUser, err = DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
	})
	return dbUser, err
}

func (usrService *UserService) GetUserIdByEmail(
	ctx context.Context,
	email string,
	DB *database.Queries) (userId uuid.UUID, err error) {
	dbUserId, err := DB.GetUserIdByEmail(ctx, email)
	return dbUserId, err
}
