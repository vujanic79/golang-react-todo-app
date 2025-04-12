package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"os"
)

func GetPostgreSQLConnection() (dbQueries *database.Queries) {
	l := logger.Get()

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
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("connectionParams", zerolog.Dict().
				Str("driver", dbDriver).
				Str("host", dbHost).
				Str("port", dbPort).
				Str("name", dbName).
				Str("sslmode", dbSslMode)).
			Msg("Database connection error")
	}
	dbQueries = database.New(db)
	return dbQueries
}
