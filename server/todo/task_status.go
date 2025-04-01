package todo

import (
	"context"
	"github.com/vujanic79/golang-react-todo-app/internal/database"
	"net/http"
)

type TaskStatus struct {
	Status string `json:"status"`
}

type TaskStatusService interface {
	CreateTaskStatus(ctx context.Context, taskStatus string, DB *database.Queries) (dbTaskStatus database.TaskStatus, err error)
	GetTaskStatuses(ctx context.Context, DB *database.Queries) (dbTaskStatus []database.TaskStatus, err error)
	GetTaskStatusByStatus(ctx context.Context, taskStatus string, DB *database.Queries) (dbTaskStatus database.TaskStatus, err error)
}

type TaskStatusController interface {
	CreateTaskStatusHandler(w http.ResponseWriter, r *http.Request)
	GetTaskStatusesHandler(w http.ResponseWriter, r *http.Request)
	GetTaskStatusByStatusHandler(w http.ResponseWriter, r *http.Request)
}

type ReqParamsForCreateTaskStatus struct {
	Status string `json:"status"`
}

func MapDbTaskStatusToTaskStatus(dbTaskStatus database.TaskStatus) TaskStatus {
	return TaskStatus{
		Status: dbTaskStatus.Status,
	}
}

func MapDbTaskStatusesToTaskStatuses(dbTaskStatuses []database.TaskStatus) []TaskStatus {
	taskStatuses := make([]TaskStatus, len(dbTaskStatuses))
	for index, dbTaskStatus := range dbTaskStatuses {
		taskStatuses[index] = MapDbTaskStatusToTaskStatus(dbTaskStatus)
	}
	return taskStatuses
}
