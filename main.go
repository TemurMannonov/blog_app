package main

import (
	"database/sql"
	"fmt"
	"log"

	"blog/api"
	"blog/storage"

	_ "github.com/lib/pq"
)

var (
	PostgresUser     = "postgres"
	PostgresPassword = "postgres"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresDatabase = "blog_db"
)

func main() {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPassword,
		PostgresDatabase,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	storage := storage.NewDBManager(db)

	server := api.NewServer(storage)

	err = server.Run(":8000")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
