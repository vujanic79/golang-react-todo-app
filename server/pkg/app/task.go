package app

import (
	"context"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"time"
)

type TaskService struct {
	TaskRepository domain.TaskRepository
	Ctx            context.Context
}

var _ domain.TaskService = (*TaskService)(nil)

func NewTaskService(taskRepository domain.TaskRepository) TaskService {
	return TaskService{TaskRepository: taskRepository, Ctx: context.Background()}
}

func (ts TaskService) SetContext(ctx context.Context) {
	ts.Ctx = ctx
}

func (ts TaskService) CreateTask(userId uuid.UUID, createTaskParams domain.CreateTaskParams) (task domain.Task, err error) {
	ts.TaskRepository.SetContext(ts.Ctx)
	return ts.TaskRepository.CreateTask(userId, createTaskParams)
}

func (ts TaskService) DeleteTask(id uuid.UUID) (err error) {
	ts.TaskRepository.SetContext(ts.Ctx)
	return ts.TaskRepository.DeleteTask(id)
}

func (ts TaskService) UpdateTask(updateTaskParams domain.UpdateTaskParams) (dbTask domain.Task, err error) {
	ts.TaskRepository.SetContext(ts.Ctx)
	return ts.TaskRepository.UpdateTask(updateTaskParams)
}

func (ts TaskService) GetTasksByUserId(userID uuid.UUID) (dbTasks []domain.Task, err error) {
	ts.TaskRepository.SetContext(ts.Ctx)
	return ts.TaskRepository.GetTasksByUserId(userID)
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
