package db

import (
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"testing"
)

func TestMapDbTaskStatusToTaskStatus(t *testing.T) {
	status := "PENDING"
	want := domain.TaskStatus{Status: status}

	taskStatus := mapDbTaskStatusToTaskStatus(status)
	areEqual := checkTaskStatusEquality(want, taskStatus)
	
	if !areEqual {
		t.Errorf("MapDbTaskStatusToTaskStatus(dbTaskStatus) = %v, want %v", taskStatus.Status, want.Status)
	}
}

func TestMapDbTaskStatusesToTaskStatuses(t *testing.T) {
	dbStatuses := []string{"PENDING", "ACTIVE", "COMPLETED"}
	want := generateTaskStatuses(dbStatuses)

	taskStatuses := mapDbTaskStatusesToTaskStatuses(dbStatuses)
	areEqual := checkTaskStatusesEquality(want, taskStatuses)

	if !areEqual {
		t.Errorf("MapDbTaskStatusesToTaskStatuses(dbStatuses) = %v, want %v", taskStatuses, want)
	}
}
