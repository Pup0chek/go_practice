package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	// Получаем информацию о товаре
	c.JSON(http.StatusOK, gin.H{"message": "Product details"})
}

func CreateProduct(c *gin.Context) {
	// Логика для создания товара
	c.JSON(http.StatusOK, gin.H{"message": "Product created"})
}
