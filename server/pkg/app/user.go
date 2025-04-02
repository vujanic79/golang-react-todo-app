package app

import (
	"context"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
)

type UserService struct {
	UserRepository domain.UserRepository
	Ctx            context.Context
}

var _ domain.UserService = (*UserService)(nil)

func NewUserService(db domain.UserRepository) *UserService {
	return &UserService{UserRepository: db}
}

func (us *UserService) SetContext(ctx context.Context) {
	us.Ctx = ctx
}

func (us *UserService) CreateUser(createUserParams domain.CreateUserParams) (user domain.User, err error) {
	us.UserRepository.SetContext(us.Ctx)
	user, err = us.UserRepository.CreateUser(createUserParams)
	return user, err
}

func (us *UserService) GetUserIdByEmail(email string) (userId uuid.UUID, err error) {
	us.UserRepository.SetContext(us.Ctx)
	userId, err = us.UserRepository.GetUserIdByEmail(email)
	return userId, err
}
