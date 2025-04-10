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
	GetTaskStatuses(ctx context.Context) (taskStatuses []TaskStatus, err error)
	GetTaskStatusByStatus(ctx context.Context, taskStatusStr string) (taskStatus TaskStatus, err error)
}

type TaskStatusRepository interface {
	CreateTaskStatus(ctx context.Context, taskStatusStr string) (taskStatus TaskStatus, err error)
	GetTaskStatuses(ctx context.Context) (taskStatuses []TaskStatus, err error)
	GetTaskStatusByStatus(ctx context.Context, taskStatusStr string) (taskStatus TaskStatus, err error)
}

type TaskStatusController interface {
	CreateTaskStatus(w http.ResponseWriter, r *http.Request)
	GetTaskStatuses(w http.ResponseWriter, r *http.Request)
	GetTaskStatusByStatus(w http.ResponseWriter, r *http.Request)
}
