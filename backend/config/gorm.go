package config

import (
	"fmt"
	"os"

	"github.com/Zaidannzzz/test/backend/httpserver/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// Get the database connection details from environment variables
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	// Construct the DSN string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Perform database migrations or other setup tasks here
	err = runMigrations(db)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")

	return db, nil
}

func runMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.UserModel{})
	if err != nil {
		return err
	}

	// Additional migrations and setup tasks can be added here

	fmt.Println("Database migrations completed")

	return nil
}
