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

var _ todo.TaskStatusController = (*TaskStatusController)(nil)

type TaskStatusController struct {
	DB *database.Queries
}

func NewTaskStatusController(DB *database.Queries) *TaskStatusController {
	return &TaskStatusController{}
}

func (taskStatusController *TaskStatusController) CreateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var decodedParams todo.ReqParamsForCreateTaskStatus
	err := decoder.Decode(&decodedParams)
	if err != nil {
		log.Printf("Error parsing task status data from the body: %s", err.Error())
		http_response.RespondWithError(w, http.StatusBadRequest, "Error parsing task status data from the body")
		return
	}

	taskStatusService := sql.NewTaskStatusService()

	dbTaskStatus, err := taskStatusService.CreateTaskStatus(r.Context(), decodedParams.Status, taskStatusController.DB)
	if err != nil {
		log.Printf("Error creating task status: %s", err.Error())
		http_response.RespondWithError(w, http.StatusInternalServerError, "Error creating task status")
		return
	}

	http_response.RespondWithJson(w, http.StatusCreated, todo.MapDbTaskStatusToTaskStatus(dbTaskStatus))
}

func (taskStatusController *TaskStatusController) GetTaskStatusesHandler(w http.ResponseWriter, r *http.Request) {
	taskStatusService := sql.NewTaskStatusService()

	dbTaskStatuses, err := taskStatusService.GetTaskStatuses(r.Context(), taskStatusController.DB)
	if err != nil {
		log.Printf("Error getting task statuses: %s", err.Error())
		http_response.RespondWithError(w, http.StatusInternalServerError, "Error getting task statuses")
		return
	}

	http_response.RespondWithJson(w, http.StatusOK, todo.MapDbTaskStatusesToTaskStatuses(dbTaskStatuses))
}

func (taskStatusController *TaskStatusController) GetTaskStatusByStatusHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var decodedParams todo.ReqParamsForCreateTaskStatus
	err := decoder.Decode(&decodedParams)
	if err != nil {
		log.Printf("Error parsing task status data from the body: %s", err.Error())
		http_response.RespondWithError(w, http.StatusBadRequest, "Error parsing task status data from the body")
		return
	}

	taskStatusService := sql.NewTaskStatusService()
	dbTaskStatus, err := taskStatusService.GetTaskStatusByStatus(r.Context(), decodedParams.Status, taskStatusController.DB)
	if err != nil {
		log.Printf("Error getting task status: %s", err.Error())
		http_response.RespondWithError(w, http.StatusInternalServerError, "Error getting task status")
		return
	}

	http_response.RespondWithJson(w, http.StatusOK, todo.MapDbTaskStatusToTaskStatus(dbTaskStatus))
}
