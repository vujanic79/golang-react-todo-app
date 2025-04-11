package app

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
)

type TaskService struct {
	TaskRepository domain.TaskRepository
}

var _ domain.TaskService = (*TaskService)(nil)

func NewTaskService(taskRepository domain.TaskRepository) TaskService {
	return TaskService{TaskRepository: taskRepository}
}

func (ts TaskService) CreateTask(ctx context.Context, userId uuid.UUID, createTaskParams domain.CreateTaskParams) (task domain.Task, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.CreateTask_params", zerolog.Dict().
			Interface("userId", userId).
			Object("createTaskParams", createTaskParams)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	return ts.TaskRepository.CreateTask(ctx, userId, createTaskParams)
}

func (ts TaskService) DeleteTask(ctx context.Context, id uuid.UUID) (err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.DeleteTask_params", zerolog.Dict().
			Interface("id", id)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	return ts.TaskRepository.DeleteTask(ctx, id)
}

func (ts TaskService) UpdateTask(ctx context.Context, updateTaskParams domain.UpdateTaskParams) (dbTask domain.Task, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.UpdateTask_params", zerolog.Dict().
			Object("updateTaskParams", updateTaskParams)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	return ts.TaskRepository.UpdateTask(ctx, updateTaskParams)
}

func (ts TaskService) GetTasksByUserId(ctx context.Context, userID uuid.UUID) (dbTasks []domain.Task, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.GetTasksByUserId_params", zerolog.Dict().
			Interface("userId", userID)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	return ts.TaskRepository.GetTasksByUserId(ctx, userID)
}
