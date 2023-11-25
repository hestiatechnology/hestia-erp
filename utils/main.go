package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectDB establishes a connection to the PostgreSQL database using environment variables.
// It retrieves the database user, password, name, and host from the environment variables PGUSER, PGPASSWORD, PGDATABASE, and PGHOST respectively.
// If any of these environment variables are not set, the function logs a fatal error and shutdowns.
// The function returns a pointer to the sql.DB object representing the database connection and an error if any occurred during the connection process.
var db *pgxpool.Pool

func ConnectDB() (*pgxpool.Pool, error) {
	if db != nil {
		return db, nil
	}

	dbUser := os.Getenv("PGUSER")
	dbPass := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")
	dbHost := os.Getenv("PGHOST")

	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" {
		log.Fatal("Missing one or more environment variables for database connection")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&application_name=hestia_api", dbUser, dbPass, dbHost, dbName)

	// Open a connection to the database
	var err error
	db, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	fmt.Println("Connected to the database")

	return db, nil
}

// Get a company's timezone from the database.
// Returns the timezone in the format "Area/Location"
// Example: "Europe/Paris"
func GetCompanyTimezone(ctx context.Context, companyId string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}

	var timezone string
	err = db.QueryRow(ctx, "SELECT timezone FROM companies.company WHERE id = $1", companyId).Scan(&timezone)
	if err != nil {
		return "", err
	}

	return timezone, nil
}

// Function to grab the Authorization header and remove the Bearer prefix,
// returning the token only. No database connection is made.
func GetSessionId(authHeader string) string {
	token := strings.TrimPrefix(authHeader, "Bearer ")
	return token
}
func GetUserId(ctx context.Context, token string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}

	var userId string
	err = db.QueryRow(ctx, "SELECT user_id FROM users.users_session WHERE id = $1", token).Scan(&userId)
	if err != nil {
		return "", err
	}

	return userId, nil
}
