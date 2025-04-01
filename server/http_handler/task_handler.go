package http_handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/vujanic79/golang-react-todo-app/http_response"
	"github.com/vujanic79/golang-react-todo-app/internal/database"
	"github.com/vujanic79/golang-react-todo-app/sql"
	"github.com/vujanic79/golang-react-todo-app/todo"
	"log"
	"net/http"
)

type TaskController struct {
	DB *database.Queries
}

var _ todo.TaskController = (*TaskController)(nil)

func NewTaskController(DB *database.Queries) *TaskController { return &TaskController{} }

func (taskController *TaskController) CreateTaskHandler(
	w http.ResponseWriter,
	r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var createTaskParams todo.CreateTaskParams
	err := decoder.Decode(&createTaskParams)
	if err != nil {
		log.Printf("Error decoding request body: %s", err.Error())
		http_response.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userService := sql.NewUserService()
	userId, err := userService.GetUserIdByEmail(r.Context(), createTaskParams.UserEmail, taskController.DB)
	if err != nil {
		log.Printf("Error getting user id: %s", err.Error())
		http_response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	taskService := sql.NewTaskService()
	dbTask, err := taskService.CreateTask(r.Context(), createTaskParams, taskController.DB, userId)
	if err != nil {
		log.Printf("Error creating task: %s", err.Error())
		http_response.RespondWithError(w, http.StatusInternalServerError, "Error creating task")
		return
	}

	http_response.RespondWithJson(w, http.StatusCreated, todo.MapDbTaskToTask(dbTask))
}

func (taskController *TaskController) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIdStr := chi.URLParam(r, "taskId")
	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		log.Printf("Error parsing task id: %s", err.Error())
		http_response.RespondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	taskService := sql.NewTaskService()
	err = taskService.DeleteTask(r.Context(), taskId, taskController.DB)
	if err != nil {
		log.Printf("Error deleting task: %s", err.Error())
		http_response.RespondWithError(w, http.StatusInternalServerError, "Error deleting task")
		return
	}
	http_response.RespondWithJson(w, http.StatusOK, fmt.Sprintf("Task with id %s successfully deleted", taskId))
}

func (taskController *TaskController) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIdStr := chi.URLParam(r, "taskId")
	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		log.Printf("Error parsing task id: %s", err.Error())
		http_response.RespondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var updateTaskParams todo.UpdateTaskParams
	err = decoder.Decode(&updateTaskParams)
	if err != nil {
		log.Printf("Error decoding request body: %s", err.Error())
		http_response.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	updateTaskParams.ID = taskId

	taskService := sql.NewTaskService()
	dbTask, err := taskService.UpdateTask(r.Context(), updateTaskParams, taskController.DB)
	if err != nil {
		log.Printf("Error updating task: %s", err.Error())
		http_response.RespondWithError(w, http.StatusInternalServerError, "Error updating task")
		return
	}

	http_response.RespondWithJson(w, http.StatusOK, todo.MapDbTaskToTask(dbTask))
}

func (taskController *TaskController) GetTasksByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var getTasksByUserIdParams todo.GetTasksByUserIdParams
	err := decoder.Decode(&getTasksByUserIdParams)
	if err != nil {
		log.Printf("Error decoding request body: %s", err.Error())
		http_response.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	taskService := sql.NewTaskService()
	tasks, err := taskService.GetTasksByUserId(r.Context(), getTasksByUserIdParams.UserID, taskController.DB)
	if err != nil {
		log.Printf("Error getting tasks: %s", err.Error())
		http_response.RespondWithError(w, http.StatusInternalServerError, "Error getting tasks")
		return
	}

	http_response.RespondWithJson(w, http.StatusOK, todo.MapDbTasksToTasks(tasks))
}
