package main

import (
	"database/sql"
	"flag"
	"fmt"
	"hestia/api/utils/logger"
	"log"
)

func main() {
	// Define command-line arguments
	dbUser := flag.String("user", "", "Database user")
	dbPassword := flag.String("password", "", "Database password")
	dbName := flag.String("database", "", "Database name")
	exportFlag := flag.Bool("export", false, "Optional export flag")

	// Parse the command-line arguments
	flag.Parse()

	// Check if mandatory arguments are provided
	if *dbUser == "" || *dbPassword == "" || *dbName == "" {
		logger.ErrorLogger.Fatal("Database user, password, and name must be provided")
	}

	// Build the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@/%s", *dbUser, *dbPassword, *dbName)

	// Connect to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	// If export flag is provided, perform export functionality
	if *exportFlag {
		logger.InfoLogger.Println("Exporting data...")
	}
}
