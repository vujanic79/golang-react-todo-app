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

func NewTaskStatusRepository(db *database.Queries) (tsr *TaskStatusRepository) {
	return &TaskStatusRepository{Db: db}
}

func (tsr *TaskStatusRepository) CreateTaskStatus(ctx context.Context, status string) (ts domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	dbStatus, err := tsr.Db.CreateTaskStatus(ctx, status)
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.CreateTaskStatus_params", zerolog.Dict().
				Str("status", status)).
			Msg("Creating task status error")
	}
	// [*] END
	return domain.TaskStatus{Status: dbStatus}, err
}

func (tsr *TaskStatusRepository) GetTaskStatuses(ctx context.Context) (tss []domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	dbTss, err := tsr.Db.GetTaskStatuses(ctx)
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.GetTaskStatuses_params", zerolog.Dict()).
			Msg("Getting task statuses from database error")
	}
	// [*] END
	return mapDbTaskStatusesToTaskStatuses(dbTss), err
}

func (tsr *TaskStatusRepository) GetTaskStatusByStatus(ctx context.Context, status string) (ts domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	dbTs, err := tsr.Db.GetTaskStatusByStatus(ctx, status)
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.GetTaskStatusByStatus_params", zerolog.Dict().
				Str("status", status)).
			Msg("Getting task status from database error")
	}
	// [*] END
	return mapDbTaskStatusToTaskStatus(dbTs), err
}

func mapDbTaskStatusToTaskStatus(dbTs string) (ts domain.TaskStatus) {
	return domain.TaskStatus{Status: dbTs}
}

func mapDbTaskStatusesToTaskStatuses(dbTss []string) (tss []domain.TaskStatus) {
	tss = make([]domain.TaskStatus, len(dbTss))
	for index, taskStatus := range dbTss {
		tss[index] = mapDbTaskStatusToTaskStatus(taskStatus)
	}
	return tss
}
