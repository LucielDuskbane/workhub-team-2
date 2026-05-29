package main

import (
	"os"
	"workhub/config"
	_ "workhub/docs"
	"workhub/internal/models"
	"workhub/internal/routes"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title WorkHub API
// @version 1.0
// @description WorkHub Final Project API
// @host workhub-team-2-production.up.railway.app
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	// Migration
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

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "WorkHub API Running",
		})
	})

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
