package http_rest

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"io"
	"net/http"
)

type TaskStatusController struct {
	TaskStatusService domain.TaskStatusService
}

var _ domain.TaskStatusController = (*TaskStatusController)(nil)

func NewTaskStatusController(taskStatusService domain.TaskStatusService) TaskStatusController {
	return TaskStatusController{TaskStatusService: taskStatusService}
}

func (tsc *TaskStatusController) CreateTaskStatus(w http.ResponseWriter, r *http.Request) {
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
	var decodedParams domain.CreateTaskStatusParams
	err = decoder.Decode(&decodedParams)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("request_body", string(rBodyBytes)). // Raw string
			Msg("Creating task status error")
		RespondWithError(w, http.StatusBadRequest, "Error parsing task status data from the body")
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.CreateTaskStatus_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			RawJSON("request_body", rBodyBytes)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END
	taskStatus, err := tsc.TaskStatusService.CreateTaskStatus(ctx, decodedParams.Status)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error creating task status")
		return
	}

	RespondWithJson(w, http.StatusCreated, taskStatus)
}

func (tsc *TaskStatusController) GetTaskStatuses(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()
	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.GetTaskStatuses_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END

	taskStatuses, err := tsc.TaskStatusService.GetTaskStatuses(ctx)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error getting task statuses")
		return
	}

	RespondWithJson(w, http.StatusOK, taskStatuses)
}

func (tsc *TaskStatusController) GetTaskStatusByStatus(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()
	taskStatusParam := chi.URLParam(r, "taskStatus")

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.CreateTaskStatusByStatus_params", zerolog.Dict().
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("urlParam", taskStatusParam)).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END

	taskStatus, err := tsc.TaskStatusService.GetTaskStatusByStatus(ctx, taskStatusParam)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error getting task status")
		return
	}

	RespondWithJson(w, http.StatusOK, taskStatus)
}
