package http_rest

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"log"
	"log/slog"
	"net/http"
)

type TaskStatusController struct {
	TaskStatusService domain.TaskStatusService
}

var _ domain.TaskStatusController = (*TaskStatusController)(nil)

func NewTaskStatusController(taskStatusService domain.TaskStatusService) *TaskStatusController {
	return &TaskStatusController{TaskStatusService: taskStatusService}
}

func (tsc *TaskStatusController) CreateTaskStatus(w http.ResponseWriter, r *http.Request) {
	tsc.TaskStatusService.SetContext(r.Context())
	decoder := json.NewDecoder(r.Body)
	var decodedParams domain.CreateTaskStatusParams
	err := decoder.Decode(&decodedParams)
	if err != nil {
		log.Printf("Error parsing task status data from the body: %s", err.Error())
		RespondWithError(w, http.StatusBadRequest, "Error parsing task status data from the body")
		return
	}

	taskStatus, err := tsc.TaskStatusService.CreateTaskStatus(decodedParams.Status)
	if err != nil {
		log.Printf("Error creating task status: %s", err.Error())
		RespondWithError(w, http.StatusInternalServerError, "Error creating task status")
		return
	}

	RespondWithJson(w, http.StatusCreated, taskStatus)
}

func (tsc *TaskStatusController) GetTaskStatuses(w http.ResponseWriter, r *http.Request) {
	tsc.TaskStatusService.SetContext(r.Context())
	taskStatuses, err := tsc.TaskStatusService.GetTaskStatuses()
	if err != nil {
		log.Printf("Error getting task statuses: %s", err.Error())
		RespondWithError(w, http.StatusInternalServerError, "Error getting task statuses")
		return
	}

	RespondWithJson(w, http.StatusOK, taskStatuses)
}

func (tsc *TaskStatusController) GetTaskStatusByStatus(w http.ResponseWriter, r *http.Request) {
	tsc.TaskStatusService.SetContext(r.Context())
	taskStatusParam := chi.URLParam(r, "taskStatus")

	taskStatus, err := tsc.TaskStatusService.GetTaskStatusByStatus(taskStatusParam)
	if err != nil {
		slog.LogAttrs(r.Context(), slog.LevelError, err.Error(),
			slog.Group("requestData", slog.String("url", r.URL.String()),
				slog.String("method", r.Method), slog.String("url_param", taskStatusParam)))
		RespondWithError(w, http.StatusInternalServerError, "Error getting task status")

		slog.GroupValue()
		return
	}

	RespondWithJson(w, http.StatusOK, taskStatus)
}
