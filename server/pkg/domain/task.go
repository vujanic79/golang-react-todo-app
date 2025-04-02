package domain

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Task struct {
	ID               uuid.UUID `json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	CompleteDeadline time.Time `json:"completeDeadline"`
	UserID           uuid.UUID `json:"userId"`
}

type CreateTaskParams struct {
	ID               uuid.UUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Title            string
	Description      string
	Status           string
	CompleteDeadline string
	UserEmail        string
}

type UpdateTaskParams struct {
	ID               uuid.UUID
	Title            string
	Description      string
	Status           string
	CompleteDeadline string
}

type GetTasksByUserIdParams struct {
	UserID uuid.UUID
}

type TaskService interface {
	CreateTask(userId uuid.UUID, createTaskParams CreateTaskParams) (dbTask Task, err error)
	DeleteTask(id uuid.UUID) (err error)
	UpdateTask(updateTaskParams UpdateTaskParams) (dbTask Task, err error)
	GetTasksByUserId(userID uuid.UUID) (dbTasks []Task, err error)
	SetContext(ctx context.Context)
}

type TaskRepository interface {
	CreateTask(userId uuid.UUID, createTaskParams CreateTaskParams) (dbTask Task, err error)
	DeleteTask(id uuid.UUID) (err error)
	UpdateTask(updateTaskParams UpdateTaskParams) (dbTask Task, err error)
	GetTasksByUserId(userID uuid.UUID) (dbTasks []Task, err error)
	SetContext(ctx context.Context)
}

type TaskController interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	GetTasksByUserId(w http.ResponseWriter, r *http.Request)
}
