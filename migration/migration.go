package migration

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jackc/pgx/v5"
)

// MigrateDatabase checks if the tables exist and runs the schema file if needed.
func MigrateDatabase(conn *pgx.Conn, filePath string) error {
	if !tablesExist(conn) {
		if err := runSQLFile(conn, filePath); err != nil {
			return fmt.Errorf("failed to execute schema.sql: %v", err)
		}
		fmt.Println("Database tables created successfully!")
	} else {
		fmt.Println("Tables already exist. Skipping creation.")
	}
	return nil
}

// tablesExist checks if any required table exists in the database.
func tablesExist(conn *pgx.Conn) bool {
	var exists bool
	err := conn.QueryRow(context.Background(),
		"SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'users');",
	).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if tables exist: %v", err)
	}
	return exists
}

// runSQLFile executes a .sql file to set up the database schema.
func runSQLFile(conn *pgx.Conn, filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("unable to read file %s: %v", filePath, err)
	}

	_, err = conn.Exec(context.Background(), string(file))
	if err != nil {
		return fmt.Errorf("failed to execute SQL script: %v", err)
	}
	return nil
}
