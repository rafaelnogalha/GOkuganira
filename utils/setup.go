package utils

import (
	"os"
	"log"
	"GOkuganira/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/joho/godotenv"
)

func SetupModels() *gorm.DB {
	//Load environmenatal variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("databaseUser")
	password := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	databaseHost := os.Getenv("databaseHost")

	//Define DB connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)
	fmt.Println(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.User{})
	fmt.Println("Successfully connected!", db)
	return db
}
