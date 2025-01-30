package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	// Получаем информацию о товаре
	c.JSON(http.StatusOK, gin.H{"message": "Product details"})
}

func CreateProduct(c *gin.Context) {
	var print struct {
		Task string `json:"task"`
	}
	if err := c.ShouldBindJSON(&print); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Task %s was successfully created", print.Task),
	})
}
