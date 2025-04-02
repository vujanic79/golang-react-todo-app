package todo

import (
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"time"
)

func generateDbUser(firstName string, lastName string, email string) database.User {
	return database.User{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

func generateUser(dbUser database.User) domain.User {
	return domain.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
	}
}

func checkUserEquality(want domain.User, user domain.User) bool {
	return want.ID == user.ID &&
		want.CreatedAt == user.CreatedAt &&
		want.UpdatedAt == user.UpdatedAt &&
		want.FirstName == user.FirstName &&
		want.LastName == user.LastName &&
		want.Email == user.Email
}

func generateDbTask(title string, description string, status string) database.Task {
	return database.Task{
		ID:               uuid.New(),
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		Title:            title,
		Description:      description,
		Status:           status,
		CompleteDeadline: time.Now().UTC().Add(1 * time.Hour),
		UserID:           uuid.New(),
	}
}

func generateTask(dbTask database.Task) domain.Task {
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

func generateDbTasks() []database.Task {
	return []database.Task{
		generateDbTask("Task1", "Description1", "ACTIVE"),
		generateDbTask("Task2", "Description2", "PENDING"),
	}
}

func generateTasks(dbTasks []database.Task) []domain.Task {
	tasks := make([]domain.Task, len(dbTasks))
	for index, dbTask := range dbTasks {
		tasks[index] = generateTask(dbTask)
	}
	return tasks
}

func checkTaskEquality(want domain.Task, task domain.Task) bool {
	return want.ID == want.ID &&
		want.CreatedAt == task.CreatedAt &&
		want.UpdatedAt == task.UpdatedAt &&
		want.Title == task.Title &&
		want.Description == task.Description &&
		want.Status == task.Status &&
		want.CompleteDeadline == task.CompleteDeadline &&
		want.UserID == want.UserID
}

func generateTaskStatuses(dbStatuses []string) []domain.TaskStatus {
	taskStatuses := make([]domain.TaskStatus, len(dbStatuses))
	for index, dbStatus := range dbStatuses {
		taskStatuses[index] = domain.TaskStatus{Status: dbStatus}
	}
	return taskStatuses
}
