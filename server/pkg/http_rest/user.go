package http_rest

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"io"
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
	l := logger.Get()

	// [*] START: Reading r.Body data, and restoring it for further usage
	rBodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).Msg("Reading request body error")
		http.Error(w, "Could not read user input", http.StatusInternalServerError)
		return
	}

	reader := io.NopCloser(bytes.NewBuffer(rBodyBytes))
	r.Body = reader
	// [*] END

	decoder := json.NewDecoder(r.Body)
	var createUserParams domain.CreateUserParams
	err = decoder.Decode(&createUserParams)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("request_body", string(rBodyBytes)). // Raw string
			Msg("Creating user error")
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.CreateUser_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			RawJSON("request_body", rBodyBytes)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END

	user, err := uc.UserService.CreateUser(ctx, createUserParams)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusCreated, user)
}
