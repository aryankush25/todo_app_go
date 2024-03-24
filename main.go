package main

import (
	"todo_app_go/internal/config"
	"todo_app_go/internal/handler"
	"todo_app_go/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	setupDatabase()
	// defer db.Close()

	router := gin.Default()
	setupRoutes(router)

	router.Run(":8080")
}

func setupDatabase() *gorm.DB {
	config.Load()
	dsn := config.DatabaseDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	return db
}

func setupRoutes(router *gin.Engine) {
	authHandler := handler.NewAuthHandler()
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
}
