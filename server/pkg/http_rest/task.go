package http_rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"log"
	"net/http"
)

type TaskController struct {
	TaskService domain.TaskService
	UserService domain.UserService
}

var _ domain.TaskController = (*TaskController)(nil)

func NewTaskController(taskService domain.TaskService, userService domain.UserService) TaskController {
	return TaskController{TaskService: taskService, UserService: userService}
}

func (tc *TaskController) CreateTask(
	w http.ResponseWriter,
	r *http.Request) {
	tc.TaskService.SetContext(r.Context())
	tc.UserService.SetContext(r.Context())
	decoder := json.NewDecoder(r.Body)
	var createTaskParams domain.CreateTaskParams
	err := decoder.Decode(&createTaskParams)
	if err != nil {
		log.Printf("Error decoding request body: %s", err.Error())
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userId, err := tc.UserService.GetUserIdByEmail(createTaskParams.UserEmail)
	if err != nil {
		log.Printf("Error getting user id: %s", err.Error())
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	task, err := tc.TaskService.CreateTask(userId, createTaskParams)
	if err != nil {
		log.Printf("Error creating task: %s", err.Error())
		RespondWithError(w, http.StatusInternalServerError, "Error creating task")
		return
	}

	RespondWithJson(w, http.StatusCreated, task)
}

func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	tc.TaskService.SetContext(r.Context())
	taskIdStr := chi.URLParam(r, "taskId")
	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		log.Printf("Error parsing task id: %s", err.Error())
		RespondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	err = tc.TaskService.DeleteTask(taskId)
	if err != nil {
		log.Printf("Error deleting task: %s", err.Error())
		RespondWithError(w, http.StatusInternalServerError, "Error deleting task")
		return
	}
	RespondWithJson(w, http.StatusOK, fmt.Sprintf("Task with id %s successfully deleted", taskId))
}

func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	tc.TaskService.SetContext(r.Context())
	taskIdStr := chi.URLParam(r, "taskId")
	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		log.Printf("Error parsing task id: %s", err.Error())
		RespondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var updateTaskParams domain.UpdateTaskParams
	err = decoder.Decode(&updateTaskParams)
	if err != nil {
		log.Printf("Error decoding request body: %s", err.Error())
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	updateTaskParams.ID = taskId

	task, err := tc.TaskService.UpdateTask(updateTaskParams)
	if err != nil {
		log.Printf("Error updating task: %s", err.Error())
		RespondWithError(w, http.StatusInternalServerError, "Error updating task")
		return
	}

	RespondWithJson(w, http.StatusOK, task)
}

func (tc *TaskController) GetTasksByUserId(w http.ResponseWriter, r *http.Request) {
	tc.TaskService.SetContext(r.Context())
	decoder := json.NewDecoder(r.Body)
	var getTasksByUserIdParams domain.GetTasksByUserIdParams
	err := decoder.Decode(&getTasksByUserIdParams)
	if err != nil {
		log.Printf("Error decoding request body: %s", err.Error())
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	tasks, err := tc.TaskService.GetTasksByUserId(getTasksByUserIdParams.UserID)
	if err != nil {
		log.Printf("Error getting tasks: %s", err.Error())
		RespondWithError(w, http.StatusInternalServerError, "Error getting tasks")
		return
	}

	RespondWithJson(w, http.StatusOK, tasks)
}
