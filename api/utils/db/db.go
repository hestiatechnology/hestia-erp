package db

import (
	"context"
	"fmt"
	"os"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

var db *pgxpool.Pool

func GetDBPoolConn() (*pgxpool.Pool, error) {
	if db != nil {
		return db, nil
	}

	dbUser := os.Getenv("PGUSER")
	dbPass := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")
	dbHost := os.Getenv("PGHOST")

	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" {
		log.Fatal().Msg("Missing one or more environment variables for database connection")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&application_name=hestia-erp", dbUser, dbPass, dbHost, dbName)

	// Open a connection to the database
	dbconfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal().Err(err).Str("connStr", connStr).Msg("Wrong connetion string")
		// logger.ErrorLogger.Fatal("Wrong connetion string", err)
	}

	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// Allows to use Google's UUIDs
		pgxUUID.Register(conn.TypeMap())
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}

	db, err = pgxpool.NewWithConfig(context.Background(), dbconfig)
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to database")
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	log.Info().Msg("First connection to the database successful")
	return db, nil
}
