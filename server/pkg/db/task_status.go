package db

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
)

type TaskStatusRepository struct {
	Db *database.Queries
}

var _ domain.TaskStatusRepository = (*TaskStatusRepository)(nil)

func NewTaskStatusRepository(db *database.Queries) *TaskStatusRepository {
	return &TaskStatusRepository{Db: db}
}

func (tsr *TaskStatusRepository) CreateTaskStatus(ctx context.Context, taskStatusStr string) (taskStatus domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	dbTaskStatus, err := tsr.Db.CreateTaskStatus(ctx, taskStatusStr)
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.CreateTaskStatus_params", zerolog.Dict().
				Str("taskStatusStr", taskStatusStr)).
			Msg("Creating task status error")
	}
	// [*] END
	return domain.TaskStatus{Status: dbTaskStatus}, err
}

func (tsr *TaskStatusRepository) GetTaskStatuses(ctx context.Context) (taskStatuses []domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	dbTaskStatuses, err := tsr.Db.GetTaskStatuses(ctx)
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.GetTaskStatuses_params", zerolog.Dict()).
			Msg("Getting task statuses from database error")
	}
	// [*] END
	return mapDbTaskStatusesToTaskStatuses(dbTaskStatuses), err
}

func (tsr *TaskStatusRepository) GetTaskStatusByStatus(ctx context.Context, taskStatusStr string) (taskStatus domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	dbTaskStatus, err := tsr.Db.GetTaskStatusByStatus(ctx, taskStatusStr)
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.GetTaskStatusByStatus_params", zerolog.Dict().
				Str("taskStatusStr", taskStatusStr)).
			Msg("Getting task status from database error")
	}
	// [*] END
	return mapDbTaskStatusToTaskStatus(dbTaskStatus), err
}

func mapDbTaskStatusToTaskStatus(dbTask string) domain.TaskStatus {
	return domain.TaskStatus{Status: dbTask}
}

func mapDbTaskStatusesToTaskStatuses(dbTaskStatuses []string) (taskStatuses []domain.TaskStatus) {
	taskStatuses = make([]domain.TaskStatus, len(dbTaskStatuses))
	for index, taskStatus := range dbTaskStatuses {
		taskStatuses[index] = mapDbTaskStatusToTaskStatus(taskStatus)
	}
	return taskStatuses
}
