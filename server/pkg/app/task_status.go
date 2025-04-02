package app

import (
	"context"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
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

func (tss *TaskStatusService) CreateTaskStatus(taskStatusStr string) (taskStatus domain.TaskStatus, err error) {
	tss.TaskStatusRepository.SetContext(tss.Ctx)
	taskStatus, err = tss.TaskStatusRepository.CreateTaskStatus(taskStatusStr)
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
