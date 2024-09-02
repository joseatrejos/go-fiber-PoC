package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "go-fiber-PoC/models"
)

var DB *gorm.DB

func LoadEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitDB() {
	// Build the DSN connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	// Initialize the database connection
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Configure the database connection pooling
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get SQL database object")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	// We must add all of our models here for auto-migration to work
	err = DB.AutoMigrate(
		&models.User{},
		&models.Expediente{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations")
	}
}
