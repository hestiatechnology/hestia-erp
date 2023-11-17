package utils

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Database
func ConnectDB() (*sql.DB, error) {
	connStr := "postgres://postgres:alexis27@localhost/erp?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	return db, err
}
