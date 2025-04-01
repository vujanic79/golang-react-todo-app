package sql

import (
	"context"
	"github.com/vujanic79/golang-react-todo-app/internal/database"
	"github.com/vujanic79/golang-react-todo-app/todo"
)

var _ todo.TaskStatusService = (*TaskStatusService)(nil)

type TaskStatusService struct {
}

func NewTaskStatusService() *TaskStatusService {
	return &TaskStatusService{}
}

func (taskStatusService *TaskStatusService) CreateTaskStatus(
	ctx context.Context,
	taskStatus string,
	DB *database.Queries) (dbTaskStatus database.TaskStatus, err error) {
	taskStatus, err = DB.CreateTaskStatus(ctx, taskStatus)
	return database.TaskStatus{Status: taskStatus}, err
}

func (taskStatusService *TaskStatusService) GetTaskStatuses(
	ctx context.Context,
	DB *database.Queries) (dbTaskStatuses []database.TaskStatus, err error) {
	taskStatuses, err := DB.GetTaskStatuses(ctx)
	dbTaskStatuses = make([]database.TaskStatus, len(taskStatuses))
	for index, taskStatus := range taskStatuses {
		dbTaskStatuses[index] = database.TaskStatus{Status: taskStatus}
	}
	return dbTaskStatuses, err
}

func (taskStatusService *TaskStatusService) GetTaskStatusByStatus(
	ctx context.Context,
	taskStatus string,
	DB *database.Queries) (dbTaskStatus database.TaskStatus, err error) {
	taskStatus, err = DB.GetTaskStatusByStatus(ctx, taskStatus)
	return database.TaskStatus{Status: taskStatus}, err
}
