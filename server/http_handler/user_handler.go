package http_handler

import (
	"encoding/json"
	"github.com/vujanic79/golang-react-todo-app/http_response"
	"github.com/vujanic79/golang-react-todo-app/internal/database"
	"github.com/vujanic79/golang-react-todo-app/sql"
	"github.com/vujanic79/golang-react-todo-app/todo"
	"log"
	"net/http"
)

var _ todo.UserController = (*UserController)(nil)

type UserController struct {
	DB *database.Queries
}

func NewUserController(DB *database.Queries) *UserController { return &UserController{} }

func (usrController *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var decodedParams todo.ReqParamsForCreateUser
	err := decoder.Decode(&decodedParams)
	if err != nil {
		log.Printf("Error parsing user data from the body: %s", err.Error())
		http_response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	userService := sql.NewUserService()
	dbUser, err := userService.CreateUser(r.Context(), decodedParams, usrController.DB)
	if err != nil {
		log.Printf("Error creating user: %s", err.Error())
		http_response.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http_response.RespondWithJson(w, http.StatusCreated, todo.MapDbUserToUser(dbUser))
}
