package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConstructor func(driverName, dataSourceName string) (*sql.DB, error)

func Connect(dbConstructor DBConstructor, envPath ...string) (*sql.DB, error) {
	if dbConstructor == nil {
		dbConstructor = sql.Open
	}

	// Load .env file if a path is provided
	if len(envPath) > 0 {
		if err := godotenv.Load(envPath[0]); err != nil {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	// Check if all required environment variables are set
	requiredEnvVars := []string{"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			return nil, fmt.Errorf("missing required environment variable: %s", envVar)
		}
	}

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := dbConstructor("mysql", dns)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database successfully!")

	return db, nil
}
