package data

import (
	"context"
	"github.com/gocarina/gocsv"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"log"
	"os"
	"strings"
)

type TaskStatus struct {
	Status string `csv:"status"`
}

func LoadDataToDatabase(dbQueries *database.Queries, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var entries []TaskStatus
	err = gocsv.Unmarshal(file, &entries)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		_, err := dbQueries.CreateTaskStatus(context.Background(), entry.Status)
		if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			log.Fatal(err)
		}
	}
}
