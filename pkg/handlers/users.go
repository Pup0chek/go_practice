package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// Получаем информацию о пользователе
	c.JSON(http.StatusOK, gin.H{"message": "User details"})
}

func CreateUser(c *gin.Context) {
	// Логика для создания пользователя
	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}
