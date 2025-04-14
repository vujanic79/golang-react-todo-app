package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/pkg/errors"
	"github.com/vujanic79/golang-react-todo-app/pkg/controller"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"github.com/vujanic79/golang-react-todo-app/pkg/repository"
	"github.com/vujanic79/golang-react-todo-app/pkg/repository/data"
	"github.com/vujanic79/golang-react-todo-app/pkg/service"
	"net/http"
	"os"
)

const taskStatusesData = "./pkg/repository/data/task_statuses.csv"

func main() {
	l := logger.Get()
	p := os.Getenv("PORT")
	if p == "" {
		err := errors.New("PORT environment variable not set")
		l.Fatal().Stack().Err(errors.WithStack(err)).
			Msg("Setting PORT environment variable error")
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	dbQueries := db.GetPostgreSQLConnection()
	data.LoadDataToDatabase(dbQueries, taskStatusesData)

	userRepository := db.NewUserRepository(dbQueries)
	taskRepository := db.NewTaskRepository(dbQueries)
	taskStatusRepository := db.NewTaskStatusRepository(dbQueries)

	userService := service.NewUserService(userRepository)
	taskService := service.NewTaskService(taskRepository)
	taskStatusService := service.NewTaskStatusService(taskStatusRepository)

	userController := controller.NewUserController(userService)
	taskController := controller.NewTaskController(taskService, userService)
	taskStatusController := controller.NewTaskStatusController(taskStatusService)

	subR := chi.NewRouter()
	subR.Post("/users", userController.CreateUser)
	subR.Post("/tasks-by-user", taskController.GetTasksByUserId)
	subR.Post("/tasks", taskController.CreateTask)
	subR.Delete("/tasks/{taskId}", taskController.DeleteTask)
	subR.Put("/tasks/{taskId}", taskController.UpdateTask)
	subR.Post("/task-status", taskStatusController.CreateTaskStatus)
	subR.Get("/task-status", taskStatusController.GetTaskStatuses)
	subR.Get("/task-status/{taskStatus}", taskStatusController.GetTaskStatusByStatus)

	r.Mount("/todo", subR)

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + p,
	}

	err := srv.ListenAndServe()
	if err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).
			Msg("Starting HTTP server error")
	}
}
