package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"time"
)

type UserRepository struct {
	Db *database.Queries
}

var _ domain.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *database.Queries) *UserRepository {
	return &UserRepository{Db: db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, createUserParams domain.CreateUserParams) (user domain.User, err error) {
	l := logger.FromContext(ctx)
	dbUser, err := ur.Db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: createUserParams.FirstName,
		LastName:  createUserParams.LastName,
		Email:     createUserParams.Email,
	})
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.CreateUser_params", zerolog.Dict().
				Object("createUserParams", createUserParams)).
			Msg("Creating user error")
	}
	// [*] END
	return mapDbUserToUser(dbUser), err
}

func (ur *UserRepository) GetUserIdByEmail(ctx context.Context, email string) (userId uuid.UUID, err error) {
	l := logger.FromContext(ctx)
	userId, err = ur.Db.GetUserIdByEmail(ctx, email)
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.GetUserIdByEmail_params", zerolog.Dict().
				Str("email", email)).
			Msg("Getting user id by email error")
	}
	// [*] END
	return userId, err
}

func mapDbUserToUser(dbUser database.AppUser) domain.User {
	return domain.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
	}
}
