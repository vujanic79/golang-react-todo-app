package app

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
)

type TaskStatusService struct {
	Tsr domain.TaskStatusRepository
}

var _ domain.TaskStatusService = (*TaskStatusService)(nil)

func NewTaskStatusService(tsr domain.TaskStatusRepository) (tss *TaskStatusService) {
	return &TaskStatusService{Tsr: tsr}
}

func (tss *TaskStatusService) CreateTaskStatus(ctx context.Context, status string) (ts domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.CreateTaskStatus_params", zerolog.Dict().
			Str("status", status)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	ts, err = tss.Tsr.CreateTaskStatus(ctx, status)
	return ts, err
}

func (tss *TaskStatusService) GetTaskStatuses(ctx context.Context) (taskStatuses []domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.GetTaskStatusByStatus_params", zerolog.Dict()).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	taskStatuses, err = tss.Tsr.GetTaskStatuses(ctx)
	return taskStatuses, err
}

func (tss *TaskStatusService) GetTaskStatusByStatus(ctx context.Context, status string) (ts domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	// [*] START - Add service data to context
	l = l.With().
		Dict("app.GetTaskStatusByStatus_params", zerolog.Dict().
			Str("status", status)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	ts, err = tss.Tsr.GetTaskStatusByStatus(ctx, status)
	return ts, err
}
