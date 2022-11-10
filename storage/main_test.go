package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	dbManager *DBManager

	// Postgres credentials
	PostgresUser     = "postgres"
	PostgresPassword = "postgres"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresDatabase = "blog_db"
)

func TestMain(m *testing.M) {
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

	dbManager = NewDBManager(db)
	os.Exit(m.Run())
}
