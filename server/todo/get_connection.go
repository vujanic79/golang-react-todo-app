package todo

import (
	"database/sql"
	"github.com/vujanic79/golang-react-todo-app/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func GetDbConnection() (apiCfg *ApiConfig) {
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	return &ApiConfig{
		DB: dbQueries,
	}
}
