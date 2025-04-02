package http_rest

import (
	"encoding/json"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"log"
	"net/http"
)

var _ domain.UserController = (*UserController)(nil)

type UserController struct {
	UserService domain.UserService
}

func NewUserController(userService domain.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	uc.UserService.SetContext(r.Context())
	decoder := json.NewDecoder(r.Body)
	var createUserParams domain.CreateUserParams
	err := decoder.Decode(&createUserParams)
	if err != nil {
		log.Printf("Error parsing user data from the body: %s", err.Error())
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := uc.UserService.CreateUser(createUserParams)
	if err != nil {
		log.Printf("Error creating user: %s", err.Error())
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusCreated, user)
}
