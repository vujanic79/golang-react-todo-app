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
	TaskService domain.TaskService
	UserService domain.UserService
}

var _ domain.TaskController = (*TaskController)(nil)

func NewTaskController(taskService domain.TaskService, userService domain.UserService) TaskController {
	return TaskController{TaskService: taskService, UserService: userService}
}

func (tc *TaskController) CreateTask(
	w http.ResponseWriter,
	r *http.Request) {
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
	var createTaskParams domain.CreateTaskParams
	err = decoder.Decode(&createTaskParams)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("requestBody", string(rBodyBytes)). // Raw string
			Msg("Creating task error")
		RespondWithError(w, http.StatusBadRequest, "Parsing task data from the body error")
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.CreateTask_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			RawJSON("requestBody", rBodyBytes)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END
	userId, err := tc.UserService.GetUserIdByEmail(ctx, createTaskParams.UserEmail)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	task, err := tc.TaskService.CreateTask(ctx, userId, createTaskParams)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error creating task")
		return
	}

	RespondWithJson(w, http.StatusCreated, task)
}

func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()
	taskIdStr := chi.URLParam(r, "taskId")
	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("taskIdStr", taskIdStr).
			Msg("Parsing taskIdStr error")
		RespondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.DeleteTask_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("urlParam", taskIdStr)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END
	err = tc.TaskService.DeleteTask(ctx, taskId)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error deleting task")
		return
	}
	RespondWithJson(w, http.StatusOK, fmt.Sprintf("Task with id %s successfully deleted", taskId))
}

func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()
	taskIdStr := chi.URLParam(r, "taskId")
	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("taskIdStr", taskIdStr).
			Msg("Parsing taskIdStr error")
		RespondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

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
	var updateTaskParams domain.UpdateTaskParams
	err = decoder.Decode(&updateTaskParams)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("requestBody", string(rBodyBytes)). // Raw string
			Msg("Updating task error")
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.UpdateTask_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			RawJSON("requestBody", rBodyBytes)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END
	updateTaskParams.ID = taskId
	task, err := tc.TaskService.UpdateTask(ctx, updateTaskParams)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error updating task")
		return
	}

	RespondWithJson(w, http.StatusOK, task)
}

func (tc *TaskController) GetTasksByUserId(w http.ResponseWriter, r *http.Request) {
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
	var getTasksByUserIdParams domain.GetTasksByUserIdParams
	err = decoder.Decode(&getTasksByUserIdParams)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("requestBody", string(rBodyBytes)). // Raw string
			Msg("Getting task by userId error")
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.GetTasksByUserId_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			RawJSON("requestBody", rBodyBytes)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END
	tasks, err := tc.TaskService.GetTasksByUserId(ctx, getTasksByUserIdParams.UserID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error getting tasks")
		return
	}

	RespondWithJson(w, http.StatusOK, tasks)
}
