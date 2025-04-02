package db

import (
	"testing"
)

func TestMapDbTaskToTask(t *testing.T) {
	dbTask := generateDbTask("Task 1", "Task 1 description", "ACTIVE")
	want := generateTask(dbTask)

	task := mapDbTaskToTask(dbTask)
	areEqual := checkTaskEquality(want, task)

	if !areEqual {
		t.Errorf("MapDbTaskToTask(dbTask) = %v, want %v", task, want)
	}
}

func TestMapDbTasksToTasks(t *testing.T) {
	dbTasks := generateDbTasks()
	want := generateTasks(dbTasks)

	tasks := mapDbTasksToTasks(dbTasks)
	areEqual := checkTasksEquality(want, tasks)

	if !areEqual {
		t.Errorf("MapDbTasksToTasks(dbTasks) = %v, want %v", tasks, want)
	}
}
