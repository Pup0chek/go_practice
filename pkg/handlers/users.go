package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// Получаем информацию о пользователе
	c.JSON(http.StatusOK, gin.H{"message": "User details"})
}

func CreateUser(c *gin.Context) {
	var print struct {
		Login string `json:"login"`
	}
	if err := c.ShouldBindJSON(&print); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("User %s was successfully created", print.Login),
	})
}
