package utils

import (
	"GOkuganira/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

func SetupModels() *gorm.DB {
	dbURI := "postgres://wkzyteeoeezzow:a1be4d471410b30b7840b9c5ed4498e93404767dc4b58190a7fea2bd166fa6cc@ec2-54-204-148-110.compute-1.amazonaws.com:5432/dbh1jqp64bufnj"

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.User{})
	return db
}
