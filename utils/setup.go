package utils

import (
	"GOkuganira/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

func SetupModels() *gorm.DB {
	//Load environmenatal variables
	//err := godotenv.Load()
	dbURI := "postgres://wkzyteeoeezzow:a1be4d471410b30b7840b9c5ed4498e93404767dc4b58190a7fea2bd166fa6cc@ec2-54-204-148-110.compute-1.amazonaws.com:5432/dbh1jqp64bufnj"
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// username := os.Getenv("DATABASE_USER")
	// password := os.Getenv("DATABASE_PASSWORD")
	// databaseName := os.Getenv("DATABASE_NAME")
	// databaseHost := os.Getenv("DATABASE_HOST")

	//Define DB connection string
	//dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)
	//fmt.Println(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.User{})
	fmt.Println("Successfully connected!", db)
	return db
}
