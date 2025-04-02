package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"log"
	"os"
)

func GetPostgreSQLConnection() (dbQueries *database.Queries) {
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries = database.New(db)
	return dbQueries
}
