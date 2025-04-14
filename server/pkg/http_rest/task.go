package http_rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"io"
	"net/http"
)

type TaskController struct {
	Ts domain.TaskService
	Us domain.UserService
}

var _ domain.TaskController = (*TaskController)(nil)

func NewTaskController(ts domain.TaskService, us domain.UserService) (tc TaskController) {
	return TaskController{Ts: ts, Us: us}
}

func (tc *TaskController) CreateTask(
	w http.ResponseWriter,
	r *http.Request) {
	l := logger.Get()

	// [*] START: Reading r.Body data, and restoring it for further usage
	b, err := io.ReadAll(r.Body)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).Msg("Reading request body error")
		http.Error(w, "Could not read user input", http.StatusInternalServerError)
		return
	}

	reader := io.NopCloser(bytes.NewBuffer(b))
	r.Body = reader
	// [*] END

	decoder := json.NewDecoder(r.Body)
	var params domain.CreateTaskParams
	err = decoder.Decode(&params)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("body", string(b)). // Raw string
			Msg("Creating t error")
		RespondWithError(w, http.StatusBadRequest, "Parsing task data from the body error")
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.CreateTask_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			RawJSON("body", b)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END
	userId, err := tc.Us.GetUserIdByEmail(ctx, params.UserEmail)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	t, err := tc.Ts.CreateTask(ctx, userId, params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error creating task")
		return
	}

	RespondWithJson(w, http.StatusCreated, t)
}

func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("id", idStr).
			Msg("Parsing id error")
		RespondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.DeleteTask_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("urlParam", idStr)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END
	err = tc.Ts.DeleteTask(ctx, id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error deleting task")
		return
	}
	RespondWithJson(w, http.StatusOK, fmt.Sprintf("Task with id %s successfully deleted", id))
}

func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("id", idStr).
			Msg("Parsing id error")
		RespondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	// [*] START: Reading r.Body data, and restoring it for further usage
	b, err := io.ReadAll(r.Body)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).Msg("Reading request body error")
		http.Error(w, "Could not read user input", http.StatusInternalServerError)
		return
	}

	reader := io.NopCloser(bytes.NewBuffer(b))
	r.Body = reader
	// [*] END

	decoder := json.NewDecoder(r.Body)
	var params domain.UpdateTaskParams
	err = decoder.Decode(&params)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("body", string(b)). // Raw string
			Msg("Updating t error")
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.UpdateTask_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			RawJSON("body", b)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END
	params.ID = id
	t, err := tc.Ts.UpdateTask(ctx, params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error updating task")
		return
	}

	RespondWithJson(w, http.StatusOK, t)
}

func (tc *TaskController) GetTasksByUserId(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()

	// [*] START: Reading r.Body data, and restoring it for further usage
	b, err := io.ReadAll(r.Body)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).Msg("Reading request body error")
		http.Error(w, "Could not read user input", http.StatusInternalServerError)
		return
	}

	reader := io.NopCloser(bytes.NewBuffer(b))
	r.Body = reader
	// [*] END
	decoder := json.NewDecoder(r.Body)
	var params domain.GetTasksByUserIdParams
	err = decoder.Decode(&params)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("body", string(b)). // Raw string
			Msg("Getting task by userId error")
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.GetTasksByUserId_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			RawJSON("body", b)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END
	ts, err := tc.Ts.GetTasksByUserId(ctx, params.UserID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error getting tasks")
		return
	}

	RespondWithJson(w, http.StatusOK, ts)
}
