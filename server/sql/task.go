package sql

import (
	"context"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/internal/database"
	"github.com/vujanic79/golang-react-todo-app/todo"
	"strings"
	"time"
)

type TaskService struct{}

var _ todo.TaskService = (*TaskService)(nil)

func NewTaskService() *TaskService { return &TaskService{} }

func (taskService *TaskService) CreateTask(
	ctx context.Context,
	createTaskParams todo.CreateTaskParams,
	DB *database.Queries,
	userId uuid.UUID) (dbTask database.Task, err error) {
	layout := "2006-01-02 15:04:05"
	date, err := time.Parse(layout, createTaskParams.CompleteDeadline)
	if err != nil {
		return dbTask, err
	}

	dbTask, err = DB.CreateTask(ctx, database.CreateTaskParams{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		Title:            createTaskParams.Title,
		Description:      createTaskParams.Description,
		Status:           strings.ToUpper(createTaskParams.Status),
		CompleteDeadline: date,
		UserID:           userId,
	})
	return dbTask, err
}

func (taskService *TaskService) DeleteTask(
	ctx context.Context,
	taskId uuid.UUID,
	DB *database.Queries) (err error) {
	err = DB.DeleteTask(ctx, taskId)
	return err
}

func (taskService *TaskService) UpdateTask(
	ctx context.Context,
	updateTaskParams todo.UpdateTaskParams,
	DB *database.Queries) (dbTask database.Task, err error) {
	dbTask, err = DB.UpdateTask(ctx, database.UpdateTaskParams{
		ID:               updateTaskParams.ID,
		Title:            updateTaskParams.Title,
		Description:      updateTaskParams.Description,
		CompleteDeadline: updateTaskParams.CompleteDeadline,
		Status:           strings.ToUpper(updateTaskParams.Status),
	})
	return dbTask, err
}

func (taskService *TaskService) GetTasksByUserId(
	ctx context.Context,
	userID uuid.UUID,
	DB *database.Queries) (dbTasks []database.Task, err error) {
	dbTasks, err = DB.GetTasksByUserId(ctx, userID)
	return dbTasks, err
}
