package app

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
)

type TaskStatusService struct {
	TaskStatusRepository domain.TaskStatusRepository
}

var _ domain.TaskStatusService = (*TaskStatusService)(nil)

func NewTaskStatusService(taskStatusRepository domain.TaskStatusRepository) *TaskStatusService {
	return &TaskStatusService{TaskStatusRepository: taskStatusRepository}
}

func (tss *TaskStatusService) CreateTaskStatus(ctx context.Context, taskStatusStr string) (taskStatus domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.CreateTaskStatus_params", zerolog.Dict().
			Str("taskStatusStr", taskStatusStr)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	taskStatus, err = tss.TaskStatusRepository.CreateTaskStatus(ctx, taskStatusStr)
	return taskStatus, err
}

func (tss *TaskStatusService) GetTaskStatuses(ctx context.Context) (taskStatuses []domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.GetTaskStatusByStatus_params", zerolog.Dict()).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	taskStatuses, err = tss.TaskStatusRepository.GetTaskStatuses(ctx)
	return taskStatuses, err
}

func (tss *TaskStatusService) GetTaskStatusByStatus(ctx context.Context, taskStatusStr string) (taskStatus domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.GetTaskStatusByStatus_params", zerolog.Dict().
			Str("taskStatusStr", taskStatusStr)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	taskStatus, err = tss.TaskStatusRepository.GetTaskStatusByStatus(ctx, taskStatusStr)
	return taskStatus, err
}
