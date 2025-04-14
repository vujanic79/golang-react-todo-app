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
	Tss domain.TaskStatusService
}

var _ domain.TaskStatusController = (*TaskStatusController)(nil)

func NewTaskStatusController(tss domain.TaskStatusService) (tsc TaskStatusController) {
	return TaskStatusController{Tss: tss}
}

func (tsc *TaskStatusController) CreateTaskStatus(w http.ResponseWriter, r *http.Request) {
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
	var params domain.CreateTaskStatusParams
	err = decoder.Decode(&params)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("body", string(b)). // Raw string
			Msg("Creating task status error")
		RespondWithError(w, http.StatusBadRequest, "Error parsing task status data from the body")
		return
	}

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.params", zerolog.Dict().
			Str("func", "CreateTaskStatus").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method).
				RawJSON("body", b))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END
	ts, err := tsc.Tss.CreateTaskStatus(ctx, params.Status)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error creating task status")
		return
	}

	RespondWithJson(w, http.StatusCreated, ts)
}

func (tsc *TaskStatusController) GetTaskStatuses(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()
	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.params", zerolog.Dict().
			Str("func", "GetTaskStatuses").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END

	tss, err := tsc.Tss.GetTaskStatuses(ctx)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error getting task statuses")
		return
	}

	RespondWithJson(w, http.StatusOK, tss)
}

func (tsc *TaskStatusController) GetTaskStatusByStatus(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()
	status := chi.URLParam(r, "taskStatus")

	// [*] START - Add http request data to context
	l = l.With().
		Dict("http_rest.params", zerolog.Dict().
			Str("func", "GetTaskStatusByStatus").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method).
				Str("urlParam", status))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)
	// [*] END

	ts, err := tsc.Tss.GetTaskStatusByStatus(ctx, status)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error getting task status")
		return
	}

	RespondWithJson(w, http.StatusOK, ts)
}
