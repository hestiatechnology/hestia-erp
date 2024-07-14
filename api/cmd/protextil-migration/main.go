package main

import (
	"database/sql"
	"flag"
	"fmt"
	"hestia/api/utils/logger"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	// Define command-line arguments
	dbServer := flag.String("server", "", "Database server IP/Domain")
	dbUser := flag.String("user", "", "Database user")
	dbPassword := flag.String("password", "", "Database password")
	exportFlag := flag.Bool("export", false, "Optional export flag, exports data to a json")

	// Define the usage message
	flag.Usage = func() {
		fmt.Println("Hestia ERP - Protextil Migration Tool")
		fmt.Println("This tool is used to migrate data from Protextil to Hestia ERP")
		fmt.Println("Usage: hst-protextil-migraton [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	// Parse the command-line arguments
	flag.Parse()

	// Check if mandatory arguments are provided
	if *dbServer == "" || *dbUser == "" || *dbPassword == "" {
		logger.ErrorLogger.Fatal("Database server, user and password must be provided")
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=1433", *dbServer, *dbUser, *dbPassword)
	logger.InfoLogger.Printf("Connecting to database: %s", connString)
	// Connect to the database
	db, err := sql.Open("mssql", connString)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		logger.ErrorLogger.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	// If export flag is provided, perform export functionality
	if *exportFlag {
		logger.InfoLogger.Println("Exporting data...")
	}
}
