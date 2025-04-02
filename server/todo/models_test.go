package todo

import (
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"slices"
	"testing"
)

func TestMapDbUserToUser(t *testing.T) {
	dbUser := generateDbUser("John", "Doe", "john.doe@gmail.com")
	want := generateUser(dbUser)

	user := domain.MapDbUserToUser(dbUser)
	areEqual := checkUserEquality(want, user)

	if !areEqual {
		t.Errorf("MapDbUserToUser(dbUser) = %v, want %v", user, want)
	}
}

func TestMapDbTaskToTask(t *testing.T) {
	dbTask := generateDbTask("Task 1", "Task 1 description", "ACTIVE")
	want := generateTask(dbTask)

	task := domain.MapDbTaskToTask(dbTask)
	areEqual := checkTaskEquality(want, task)

	if !areEqual {
		t.Errorf("MapDbTaskToTask(dbTask) = %v, want %v", task, want)
	}
}

func TestMapDbTasksToTasks(t *testing.T) {
	dbTasks := generateDbTasks()
	want := generateTasks(dbTasks)

	tasks := domain.MapDbTasksToTasks(dbTasks)
	areEqual := slices.Equal(want, tasks)

	if !areEqual {
		t.Errorf("MapDbTasksToTasks(dbTasks) = %v, want %v", tasks, want)
	}
}

func TestMapDbTaskStatusToTaskStatus(t *testing.T) {
	status := "PENDING"
	want := domain.TaskStatus{Status: status}

	taskStatus := domain.MapDbTaskStatusToTaskStatus(status)
	if taskStatus.Status != want.Status {
		t.Errorf("MapDbTaskStatusToTaskStatus(dbTaskStatus) = %v, want %v", taskStatus.Status, want.Status)
	}
}

func TestMapDbTaskStatusesToTaskStatuses(t *testing.T) {
	dbStatuses := []string{"PENDING", "ACTIVE", "COMPLETED"}
	want := generateTaskStatuses(dbStatuses)

	taskStatuses := domain.MapDbTaskStatusesToTaskStatuses(dbStatuses)
	areEqual := slices.Equal(want, taskStatuses)

	if !areEqual {
		t.Errorf("MapDbTaskStatusesToTaskStatuses(dbStatuses) = %v, want %v", taskStatuses, want)
	}
}
