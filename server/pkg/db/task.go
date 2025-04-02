package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"strings"
	"time"
)

type TaskRepository struct {
	Db  *database.Queries
	Ctx context.Context
}

var _ domain.TaskRepository = (*TaskRepository)(nil)

func NewTaskRepository(db *database.Queries) *TaskRepository {
	return &TaskRepository{Db: db, Ctx: context.Background()}
}

func (pDb *TaskRepository) SetContext(ctx context.Context) {
	pDb.Ctx = ctx
}

func (pDb *TaskRepository) CreateTask(
	userId uuid.UUID,
	createTaskParams domain.CreateTaskParams) (task domain.Task, err error) {

	layout := "2006-01-02T15:04:05.999999Z"
	date, err := time.Parse(layout, createTaskParams.CompleteDeadline)
	if err != nil {
		return task, err
	}

	dbTask, err := pDb.Db.CreateTask(pDb.Ctx, database.CreateTaskParams{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		Title:            createTaskParams.Title,
		Description:      createTaskParams.Description,
		Status:           strings.ToUpper(createTaskParams.Status),
		CompleteDeadline: date,
		UserID:           userId,
	})
	return mapDbTaskToTask(dbTask), err
}

func (pDb *TaskRepository) DeleteTask(taskId uuid.UUID) (err error) {
	err = pDb.Db.DeleteTask(pDb.Ctx, taskId)
	return err
}

func (pDb *TaskRepository) UpdateTask(updateTaskParams domain.UpdateTaskParams) (task domain.Task, err error) {

	layout := "2006-01-02T15:04:05.999999Z"
	date, err := time.Parse(layout, updateTaskParams.CompleteDeadline)
	if err != nil {
		return task, err
	}
	dbTask, err := pDb.Db.UpdateTask(pDb.Ctx, database.UpdateTaskParams{
		ID:               updateTaskParams.ID,
		Title:            updateTaskParams.Title,
		Description:      updateTaskParams.Description,
		CompleteDeadline: date,
		Status:           strings.ToUpper(updateTaskParams.Status),
	})
	return mapDbTaskToTask(dbTask), err
}
func (pDb *TaskRepository) GetTasksByUserId(userID uuid.UUID) (tasks []domain.Task, err error) {
	dbTasks, err := pDb.Db.GetTasksByUserId(pDb.Ctx, userID)
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
