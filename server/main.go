package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/vujanic79/golang-react-todo-app/http_handler"
	"github.com/vujanic79/golang-react-todo-app/sql/data"
	"github.com/vujanic79/golang-react-todo-app/todo"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT must be set")
	}

	apiCfg := todo.GetDbConnection()
	data.LoadDataToDatabase(apiCfg, "sql/data/task_statuses.csv")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	userController := http_handler.NewUserController(apiCfg.DB)
	taskController := http_handler.NewTaskController(apiCfg.DB)
	taskStatusController := http_handler.NewTaskStatusController(apiCfg.DB)

	router1 := chi.NewRouter()
	router1.Post("/users", userController.CreateUserHandler)
	router1.Post("/tasks-by-user", taskController.GetTasksByUserIdHandler)
	router1.Post("/tasks", taskController.CreateTaskHandler)
	router1.Delete("/tasks/{taskId}", taskController.DeleteTaskHandler)
	router1.Put("/tasks/{taskId}", taskController.UpdateTaskHandler)
	router1.Post("/task-statuses", taskStatusController.CreateTaskStatusHandler)
	router1.Get("/task-statuses", taskStatusController.GetTaskStatusesHandler)
	router1.Post("/task-statuses", taskStatusController.GetTaskStatusByStatusHandler)

	router.Mount("/todo", router1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start", err.Error())
	}
}
