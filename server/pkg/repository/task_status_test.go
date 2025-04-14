package db

import (
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"testing"
)

func TestMapDbTaskStatusToTaskStatus(t *testing.T) {
	status := "PENDING"
	want := domain.TaskStatus{Status: status}

	ts := mapDbTaskStatusToTaskStatus(status)
	areEqual := checkTaskStatusEquality(want, ts)

	if !areEqual {
		t.Errorf("MapDbTaskStatusToTaskStatus(status) = %v, want %v", ts.Status, want.Status)
	}
}

func TestMapDbTaskStatusesToTaskStatuses(t *testing.T) {
	statuses := []string{"PENDING", "ACTIVE", "COMPLETED"}
	want := generateTaskStatuses(statuses)

	tss := mapDbTaskStatusesToTaskStatuses(statuses)
	areEqual := checkTaskStatusesEquality(want, tss)

	if !areEqual {
		t.Errorf("MapDbTaskStatusesToTaskStatuses(statuses) = %v, want %v", tss, want)
	}
}
