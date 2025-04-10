package domain

import (
	"context"
	"net/http"
)

type TaskStatus struct {
	Status string `json:"status"`
}

type CreateTaskStatusParams struct {
	Status string `json:"status"`
}

type TaskStatusService interface {
	CreateTaskStatus(ctx context.Context, taskStatusStr string) (taskStatus TaskStatus, err error)
	GetTaskStatuses() (taskStatuses []TaskStatus, err error)
	GetTaskStatusByStatus(taskStatusStr string) (taskStatus TaskStatus, err error)
	SetContext(ctx context.Context)
}

type TaskStatusRepository interface {
	CreateTaskStatus(ctx context.Context, taskStatusStr string) (taskStatus TaskStatus, err error)
	GetTaskStatuses() (taskStatuses []TaskStatus, err error)
	GetTaskStatusByStatus(taskStatusStr string) (taskStatus TaskStatus, err error)
	SetContext(ctx context.Context)
}

type TaskStatusController interface {
	CreateTaskStatus(w http.ResponseWriter, r *http.Request)
	GetTaskStatuses(w http.ResponseWriter, r *http.Request)
	GetTaskStatusByStatus(w http.ResponseWriter, r *http.Request)
}
