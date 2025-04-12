package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"strings"
	"time"
)

type TaskRepository struct {
	Db *database.Queries
}

var _ domain.TaskRepository = (*TaskRepository)(nil)

func NewTaskRepository(db *database.Queries) *TaskRepository {
	return &TaskRepository{Db: db}
}

func (tr *TaskRepository) CreateTask(
	ctx context.Context,
	userId uuid.UUID,
	params domain.CreateTaskParams) (task domain.Task, err error) {
	l := logger.FromContext(ctx)
	layout := "2006-01-02T15:04:05.999999Z"
	date, err := time.Parse(layout, params.CompleteDeadline)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.CreateTask_params", zerolog.Dict().
				Str("completeDeadline", params.CompleteDeadline)).
			Msg("Parsing completeDeadline error")
		return task, err
	}

	dbTask, err := tr.Db.CreateTask(ctx, database.CreateTaskParams{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		Title:            params.Title,
		Description:      params.Description,
		Status:           strings.ToUpper(params.Status),
		CompleteDeadline: date,
		UserID:           userId,
	})

	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.CreateTask_params", zerolog.Dict().
				Object("params", params)).
			Msg("Creating task error")
	}
	// [*] END
	return mapDbTaskToTask(dbTask), err
}

func (tr *TaskRepository) DeleteTask(ctx context.Context, taskId uuid.UUID) (err error) {
	l := logger.FromContext(ctx)
	err = tr.Db.DeleteTask(ctx, taskId)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.DeleteTask_params", zerolog.Dict().
				Interface("taskId", taskId)).
			Msg("Deleting task error")
	}
	return err
}

func (tr *TaskRepository) UpdateTask(ctx context.Context, params domain.UpdateTaskParams) (task domain.Task, err error) {
	l := logger.FromContext(ctx)
	layout := "2006-01-02T15:04:05.999999Z"
	date, err := time.Parse(layout, params.CompleteDeadline)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.UpdateTask_params", zerolog.Dict().
				Str("completeDeadline", params.CompleteDeadline)).
			Msg("Parsing completeDeadline error")
		return task, err
	}

	dbTask, err := tr.Db.UpdateTask(ctx, database.UpdateTaskParams{
		ID:               params.ID,
		Title:            params.Title,
		Description:      params.Description,
		CompleteDeadline: date,
		Status:           strings.ToUpper(params.Status),
	})
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.UpdateTask_params", zerolog.Dict().
				Object("params", params)).
			Msg("Updating task error")
	}
	// [*] END
	return mapDbTaskToTask(dbTask), err
}
func (tr *TaskRepository) GetTasksByUserId(ctx context.Context, userID uuid.UUID) (tasks []domain.Task, err error) {
	l := logger.FromContext(ctx)
	dbTasks, err := tr.Db.GetTasksByUserId(ctx, userID)
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.GetTaskByUserId_params", zerolog.Dict().
				Interface("userId", userID)).
			Msg("Getting tasks by userId error")
	}
	// [*] END
	return mapDbTasksToTasks(dbTasks), err
}

type GetTasksByUserIdParams struct {
	UserID uuid.UUID `json:"user_id"`
}

func mapDbTaskToTask(dbTask database.Task) domain.Task {
	return domain.Task{
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

func mapDbTasksToTasks(dbTasks []database.Task) (tasks []domain.Task) {
	tasks = make([]domain.Task, len(dbTasks))
	for index, dbTask := range dbTasks {
		tasks[index] = mapDbTaskToTask(dbTask)
	}
	return tasks
}
