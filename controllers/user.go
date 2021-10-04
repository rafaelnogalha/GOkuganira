package controllers

import (
	"GOkuganira/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// GET /user
// Get all users
func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// func FindUserByName(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)
// 	var users []models.User
// 	db.Find(&users)
// 	c.JSON(http.StatusOK, gin.H{"data": users})
// }

// POST /users
// Create new Users
func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create User
	user := models.User{Username: input.Username, Password: input.Password}
	//if user.Username
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
