package main

import (
	"github.com/AnuragChaubey2/JobManagerService/db"
	"github.com/AnuragChaubey2/JobManagerService/migration"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize zap logger")
	}
	defer logger.Sync()

	err = godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	pgConn, err := db.ConnectToDB()
	if err != nil {
		logger.Fatal("Error connecting to the PostgreSQL database", zap.Error(err))
	}
	defer db.CloseDB()

	_, err = db.ConnectToRedis()
	if err != nil {
		logger.Fatal("Error connecting to Redis", zap.Error(err))
	}
	defer db.CloseRedis()

	err = migration.MigrateDatabase(pgConn, "db/migrations/schema.sql")
	if err != nil {
		logger.Fatal("Migration error", zap.Error(err))
	}

	logger.Info("Connections to PostgreSQL and Redis are established. Application is ready to use.")
}
