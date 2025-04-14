package http_rest

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/http_rest/util"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"net/http"
)

var _ domain.UserController = (*UserController)(nil)

type UserController struct {
	Us domain.UserService
}

func NewUserController(us domain.UserService) (uc UserController) {
	return UserController{Us: us}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()

	b, err := util.ReadBody(r)
	if err != nil {
		http.Error(w, "Could not read user input", http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params domain.CreateUserParams
	err = decoder.Decode(&params)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("body", string(b)). // Raw string
			Msg("Creating user error")
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.params", zerolog.Dict().
			Str("func", "CreateUser").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method).
				RawJSON("body", b))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END

	u, err := uc.Us.CreateUser(ctx, params)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJson(w, http.StatusCreated, u)
}
