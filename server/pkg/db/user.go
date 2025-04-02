package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"time"
)

type UserRepository struct {
	Db  *database.Queries
	Ctx context.Context
}

var _ domain.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *database.Queries) *UserRepository {
	return &UserRepository{Db: db, Ctx: context.Background()}
}

func (ur *UserRepository) SetContext(ctx context.Context) {
	ur.Ctx = ctx
}

func (ur *UserRepository) CreateUser(createUserParams domain.CreateUserParams) (user domain.User, err error) {
	dbUser, err := ur.Db.CreateUser(ur.Ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: createUserParams.FirstName,
		LastName:  createUserParams.LastName,
		Email:     createUserParams.Email,
	})

	return mapDbUserToUser(dbUser), err
}

func (ur *UserRepository) GetUserIdByEmail(email string) (userId uuid.UUID, err error) {
	userId, err = ur.Db.GetUserIdByEmail(ur.Ctx, email)
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
