package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"log/slog"
	"os"
)

func GetPostgreSQLConnection() (dbQueries *database.Queries) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	dbUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName, dbSslMode)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Failed to connect to database",
			slog.Group("connectionParams",
				slog.String("driver", dbDriver), slog.String("host", dbHost), slog.String("port", dbPort),
				slog.String("name", dbName), slog.String("sslmode", dbSslMode)))
	}
	dbQueries = database.New(db)
	return dbQueries
}
