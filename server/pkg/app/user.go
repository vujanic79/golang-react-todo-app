package app

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
)

type UserService struct {
	UserRepository domain.UserRepository
}

var _ domain.UserService = (*UserService)(nil)

func NewUserService(db domain.UserRepository) *UserService {
	return &UserService{UserRepository: db}
}

func (us *UserService) CreateUser(ctx context.Context, createUserParams domain.CreateUserParams) (user domain.User, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.CreateUser_params", zerolog.Dict().
			Object("create_user_params", createUserParams)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	user, err = us.UserRepository.CreateUser(ctx, createUserParams)
	return user, err
}

func (us *UserService) GetUserIdByEmail(ctx context.Context, email string) (userId uuid.UUID, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.GetUserIdByEmail_params", zerolog.Dict().
			Str("email", email)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	userId, err = us.UserRepository.GetUserIdByEmail(ctx, email)
	return userId, err
}
