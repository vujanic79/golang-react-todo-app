package todo

import (
	"context"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/internal/database"
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
	Title            string `json:"title"`
	Description      string `json:"description"`
	Status           string `json:"status"`
	CompleteDeadline string `json:"completeDeadline"`
	UserEmail        string `json:"userEmail"`
}

type UpdateTaskParams struct {
	ID               uuid.UUID `json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	CompleteDeadline time.Time `json:"completeDeadline"`
	UserID           uuid.UUID `json:"userId"`
}

type GetTasksByUserIdParams struct {
	UserID uuid.UUID `json:"user_id"`
}

type TaskService interface {
	CreateTask(ctx context.Context, createTaskParams CreateTaskParams, DB *database.Queries, userId uuid.UUID) (dbTask database.Task, err error)
	DeleteTask(ctx context.Context, id uuid.UUID, DB *database.Queries) (err error)
	UpdateTask(ctx context.Context, updateTaskParams UpdateTaskParams, DB *database.Queries) (dbTask database.Task, err error)
	GetTasksByUserId(ctx context.Context, userID uuid.UUID, DB *database.Queries) (dbTasks []database.Task, err error)
}

type TaskController interface {
	CreateTaskHandler(w http.ResponseWriter, r *http.Request)
	DeleteTaskHandler(w http.ResponseWriter, r *http.Request)
	UpdateTaskHandler(w http.ResponseWriter, r *http.Request)
	GetTasksByUserIdHandler(w http.ResponseWriter, r *http.Request)
}

func MapDbTaskToTask(dbTask database.Task) Task {
	return Task{
		ID:               dbTask.ID,
		CreatedAt:        dbTask.CreatedAt,
		UpdatedAt:        dbTask.UpdatedAt,
		Title:            dbTask.Title,
		Description:      dbTask.Description,
		Status:           dbTask.Status,
		CompleteDeadline: dbTask.CompleteDeadline,
		UserID:           dbTask.UserID,
	}
}

func MapDbTasksToTasks(dbTasks []database.Task) []Task {
	tasks := make([]Task, len(dbTasks))
	for index, dbTask := range dbTasks {
		tasks[index] = MapDbTaskToTask(dbTask)
	}
	return tasks
}
