package main

import (
	"database/sql"
	"fmt"
	"log"

	"blog/api"
	"blog/config"
	"blog/storage"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")

	fmt.Println("Config: ", cfg)

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	storage := storage.NewDBManager(db)

	server := api.NewServer(storage)

	err = server.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
