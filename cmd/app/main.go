package main

import (
	"log"
	"lol/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Логирование запроса
		log.Println("Request received")

		// Переход к следующему обработчику
		c.Next()

		// Логирование после выполнения запроса
		log.Println("Response sent")
	}
}

func main() {
	router := gin.Default()
	router.Use(MyMiddleware())

	router.GET("/lol/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "lolololol!",
		})
	})

	router.GET("/smth", func(c *gin.Context) {
		query := c.DefaultQuery("q", "default_value")
		c.JSON(200, gin.H{
			"query": query,
		})
	})

	router.GET("/print/:smth", func(c *gin.Context) {
		smth := c.Param("smth")
		c.String(200, "Print %s", smth)
	})

	router.POST("/print/", func(c *gin.Context) {
		var print struct {
			Text string `json:"text"`
		}
		if err := c.ShouldBindJSON(&print); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		c.JSON(200, gin.H{
			"message": print.Text,
		})
	})

	handlers.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
