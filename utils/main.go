package utils

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// Database
func ConnectDB() (*sql.DB, error) {
	connStr := "postgres://" + os.Getenv("") + "postgres:alexis27@localhost/erp?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	return db, err
}
