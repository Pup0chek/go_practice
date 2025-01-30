package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Маршруты для пользователей
	users := r.Group("/users")
	{
		users.GET("/:id", GetUser)
		users.POST("/", CreateUser)
	}

	// Маршруты для товаров
	products := r.Group("/products")
	{
		products.GET("/:id", GetProduct)
		products.POST("/", CreateProduct)
	}
}
