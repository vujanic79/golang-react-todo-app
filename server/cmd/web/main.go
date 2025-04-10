package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/vujanic79/golang-react-todo-app/pkg/app"
	"github.com/vujanic79/golang-react-todo-app/pkg/db"
	"github.com/vujanic79/golang-react-todo-app/pkg/db/data"
	"github.com/vujanic79/golang-react-todo-app/pkg/http_rest"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	portString := os.Getenv("PORT")
	if portString == "" {
		slog.Error("PORT must be set")
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	dbQueries := db.GetPostgreSQLConnection()
	data.LoadDataToDatabase(dbQueries, "./pkg/db/data/task_statuses.csv")

	userRepository := db.NewUserRepository(dbQueries)
	taskRepository := db.NewTaskRepository(dbQueries)
	taskStatusRepository := db.NewTaskStatusRepository(dbQueries)

	userService := app.NewUserService(userRepository)
	taskService := app.NewTaskService(taskRepository)
	taskStatusService := app.NewTaskStatusService(taskStatusRepository)

	userController := http_rest.NewUserController(userService)
	taskController := http_rest.NewTaskController(taskService, userService)
	taskStatusController := http_rest.NewTaskStatusController(taskStatusService)

	router1 := chi.NewRouter()
	router1.Post("/users", userController.CreateUser)
	router1.Post("/tasks-by-user", taskController.GetTasksByUserId)
	router1.Post("/tasks", taskController.CreateTask)
	router1.Delete("/tasks/{taskId}", taskController.DeleteTask)
	router1.Put("/tasks/{taskId}", taskController.UpdateTask)
	router1.Post("/task-status", taskStatusController.CreateTaskStatus)
	router1.Get("/task-status", taskStatusController.GetTaskStatuses)
	router1.Get("/task-status/{taskStatus}", taskStatusController.GetTaskStatusByStatus)

	router.Mount("/todo", router1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start", err.Error())
	}
}
