package main

import (
	"log"

	"github.com/AnuragChaubey2/JobManagerService/db"
	"github.com/AnuragChaubey2/JobManagerService/migration"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	pgConn, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("Error connecting to the PostgreSQL database: %v", err)
	}
	defer db.CloseDB()

	_, err = db.ConnectToRedis()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	defer db.CloseRedis()

	err = migration.MigrateDatabase(pgConn, "db/migrations/schema.sql")
	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	log.Println("Connections to PostgreSQL and Redis are established. Application is ready to use.")
}
