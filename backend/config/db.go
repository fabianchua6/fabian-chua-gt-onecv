package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // global variable

func DatabaseSetup() (*gorm.DB, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Read variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSLMODE")

	// Connect to the database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if connection successfully opened, then print success message
	if err == nil {
		fmt.Println("School DB successfully connected!")
	} else {
		panic(err)
	}

	return DB, err
}
