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

func (pDb *TaskRepository) CreateTask(
	ctx context.Context,
	userId uuid.UUID,
	createTaskParams domain.CreateTaskParams) (task domain.Task, err error) {

	l := logger.FromContext(ctx)
	layout := "2006-01-02T15:04:05.999999Z"
	date, err := time.Parse(layout, createTaskParams.CompleteDeadline)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.CreateTask_params", zerolog.Dict().
				Str("completeDeadline", createTaskParams.CompleteDeadline)).
			Msg("Parsing completeDeadline error")
		return task, err
	}

	dbTask, err := pDb.Db.CreateTask(ctx, database.CreateTaskParams{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		Title:            createTaskParams.Title,
		Description:      createTaskParams.Description,
		Status:           strings.ToUpper(createTaskParams.Status),
		CompleteDeadline: date,
		UserID:           userId,
	})

	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.CreateTask_params", zerolog.Dict().
				Object("createTaskParams", createTaskParams)).
			Msg("Creating task error")
	}
	// [*] END
	return mapDbTaskToTask(dbTask), err
}

func (pDb *TaskRepository) DeleteTask(ctx context.Context, taskId uuid.UUID) (err error) {
	l := logger.FromContext(ctx)
	err = pDb.Db.DeleteTask(ctx, taskId)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.DeleteTask_params", zerolog.Dict().
				Interface("taskId", taskId)).
			Msg("Deleting task error")
	}
	return err
}

func (pDb *TaskRepository) UpdateTask(ctx context.Context, updateTaskParams domain.UpdateTaskParams) (task domain.Task, err error) {
	l := logger.FromContext(ctx)

	layout := "2006-01-02T15:04:05.999999Z"
	date, err := time.Parse(layout, updateTaskParams.CompleteDeadline)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.UpdateTask_params", zerolog.Dict().
				Str("completeDeadline", updateTaskParams.CompleteDeadline)).
			Msg("Parsing completeDeadline error")
		return task, err
	}

	dbTask, err := pDb.Db.UpdateTask(ctx, database.UpdateTaskParams{
		ID:               updateTaskParams.ID,
		Title:            updateTaskParams.Title,
		Description:      updateTaskParams.Description,
		CompleteDeadline: date,
		Status:           strings.ToUpper(updateTaskParams.Status),
	})
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.UpdateTask_params", zerolog.Dict().
				Object("updateTaskParams", updateTaskParams)).
			Msg("Updating task error")
	}
	// [*] END
	return mapDbTaskToTask(dbTask), err
}
func (pDb *TaskRepository) GetTasksByUserId(ctx context.Context, userID uuid.UUID) (tasks []domain.Task, err error) {
	l := logger.FromContext(ctx)
	dbTasks, err := pDb.Db.GetTasksByUserId(ctx, userID)
	// [*] START - Log repository data with context
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("db.GetTaskByUserId_params", zerolog.Dict().
				Interface("userId", userID)).
			Msg("Updating task error")
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
