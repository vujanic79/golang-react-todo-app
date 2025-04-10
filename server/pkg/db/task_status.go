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
	Db  *database.Queries
	Ctx context.Context
}

var _ domain.TaskStatusRepository = (*TaskStatusRepository)(nil)

func NewTaskStatusRepository(db *database.Queries) *TaskStatusRepository {
	return &TaskStatusRepository{Db: db, Ctx: context.Background()}
}

func (tsr *TaskStatusRepository) SetContext(ctx context.Context) {
	tsr.Ctx = ctx
}

func (tsr *TaskStatusRepository) CreateTaskStatus(ctx context.Context, taskStatusStr string) (taskStatus domain.TaskStatus, err error) {
	l := logger.FromContext(ctx)
	dbTaskStatus, err := tsr.Db.CreateTaskStatus(tsr.Ctx, taskStatusStr)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.CreateTaskStatus_params", zerolog.Dict().
				Str("taskStatusStr", taskStatusStr)).
			Msg("Getting task statuses from database")
	}
	return domain.TaskStatus{Status: dbTaskStatus}, err
}

func (tsr *TaskStatusRepository) GetTaskStatuses() (taskStatuses []domain.TaskStatus, err error) {
	dbTaskStatuses, err := tsr.Db.GetTaskStatuses(tsr.Ctx)
	return mapDbTaskStatusesToTaskStatuses(dbTaskStatuses), err
}

func (tsr *TaskStatusRepository) GetTaskStatusByStatus(taskStatusStr string) (taskStatus domain.TaskStatus, err error) {
	dbTaskStatus, err := tsr.Db.GetTaskStatusByStatus(tsr.Ctx, taskStatusStr)
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
