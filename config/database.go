package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/fachrezza/todo-api/internal/model"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",

		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db
	
	err = DB.AutoMigrate(
		&model.Task{},
		&model.User{},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Connected")
}