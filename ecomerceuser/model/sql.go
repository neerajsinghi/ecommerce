package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err = gorm.Open(postgres.Open(dsn),
		&gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// AutoMigrate the User model
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate the database")
	}

	fmt.Println("Successfully connected to the database!")
}

func GetDB() *gorm.DB {
	return db
}
