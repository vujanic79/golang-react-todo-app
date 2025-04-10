package app

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
)

type TaskStatusService struct {
	TaskStatusRepository domain.TaskStatusRepository
	Ctx                  context.Context
}

var _ domain.TaskStatusService = (*TaskStatusService)(nil)

func NewTaskStatusService(taskStatusRepository domain.TaskStatusRepository) *TaskStatusService {
	return &TaskStatusService{TaskStatusRepository: taskStatusRepository, Ctx: context.Background()}
}

func (tss *TaskStatusService) SetContext(ctx context.Context) {
	tss.Ctx = ctx
}

func (tss *TaskStatusService) CreateTaskStatus(ctx context.Context, taskStatusStr string) (taskStatus domain.TaskStatus, err error) {
	tss.TaskStatusRepository.SetContext(tss.Ctx)
	// [*] START - Add service data to context
	l := logger.FromContext(ctx)
	l = l.With().
		Dict("app.TaskStatusService_params", zerolog.Dict().
			Str("taskStatusStr", taskStatusStr)).
		Logger()
	ctx = logger.WithLogger(ctx, l)
	// [*] END
	taskStatus, err = tss.TaskStatusRepository.CreateTaskStatus(ctx, taskStatusStr)
	return taskStatus, err
}

func (tss *TaskStatusService) GetTaskStatuses() (taskStatuses []domain.TaskStatus, err error) {
	tss.TaskStatusRepository.SetContext(tss.Ctx)
	taskStatuses, err = tss.TaskStatusRepository.GetTaskStatuses()
	return taskStatuses, err
}

func (tss *TaskStatusService) GetTaskStatusByStatus(taskStatusStr string) (taskStatus domain.TaskStatus, err error) {
	tss.TaskStatusRepository.SetContext(tss.Ctx)
	taskStatus, err = tss.TaskStatusRepository.GetTaskStatusByStatus(taskStatusStr)
	return taskStatus, err
}
