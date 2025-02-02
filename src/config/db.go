package config

import (
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	"github.com/rafialg11/rafi_BE_assesment/src/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
)

func InitDB() error {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	var DATABASE_URI string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	err = Database.AutoMigrate(
		&entities.Transaction{},
		&entities.Customer{},
		&entities.Account{},
	)
	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	fmt.Printf("Database Connected Successfully")
	return nil
}
