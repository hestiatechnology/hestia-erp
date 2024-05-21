package db

import (
	"context"
	"fmt"
	"hestia/api/utils/logger"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

var db *pgxpool.Pool

func GetDbPoolConn() (*pgxpool.Pool, error) {
	if db != nil {
		return db, nil
	}

	dbUser := os.Getenv("PGUSER")
	dbPass := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")
	dbHost := os.Getenv("PGHOST")

	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" {
		logger.ErrorLogger.Fatal("Missing one or more environment variables for database connection")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&application_name=hestia_api", dbUser, dbPass, dbHost, dbName)

	// Open a connection to the database
	dbconfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.ErrorLogger.Fatal("Wrong connetion string", err)
	}

	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// Allows to use Google's UUIDs
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	db, err = pgxpool.NewWithConfig(context.Background(), dbconfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	logger.InfoLogger.Println("First connection to the database successful")
	return db, nil
}
