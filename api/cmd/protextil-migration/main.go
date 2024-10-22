package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/rs/zerolog/log"
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
		log.Fatal().Msg("Database server, user and password must be provided")
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=1433", *dbServer, *dbUser, *dbPassword)
	log.Info().Str("connStr", connString).Msg("Connecting to database")
	// Connect to the database
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to database")
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to database")
	}

	fmt.Println("Successfully connected to the database!")

	// If export flag is provided, perform export functionality
	if *exportFlag {
		log.Info().Msg("Exporting data...")
	}
}
