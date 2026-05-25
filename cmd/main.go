package main

import (
	"os"
	"workhub/config"
	"workhub/internal/models"
	"workhub/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	config.ConnectDatabase()

	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Job{},
		&models.Application{},
	)

	if err != nil {
		panic("Failed migration")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "WorkHub API Running",
		})
	})

	routes.SetupRoutes(r)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
