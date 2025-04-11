package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/pkg/errors"
	"github.com/vujanic79/golang-react-todo-app/pkg/app"
	"github.com/vujanic79/golang-react-todo-app/pkg/db"
	"github.com/vujanic79/golang-react-todo-app/pkg/db/data"
	"github.com/vujanic79/golang-react-todo-app/pkg/http_rest"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"net/http"
	"os"
)

const taskStatusesData = "./pkg/db/data/task_statuses.csv"

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

	userService := app.NewUserService(userRepository)
	taskService := app.NewTaskService(taskRepository)
	taskStatusService := app.NewTaskStatusService(taskStatusRepository)

	userController := http_rest.NewUserController(userService)
	taskController := http_rest.NewTaskController(taskService, userService)
	taskStatusController := http_rest.NewTaskStatusController(taskStatusService)

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
