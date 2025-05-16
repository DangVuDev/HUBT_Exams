package main

import (
	"log"
	"os"

	"github.com/dangvudev/HUBT_ExamS/config"
	_ "github.com/dangvudev/HUBT_ExamS/docs"
	"github.com/dangvudev/HUBT_ExamS/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title HUBT Exam Storage API
// @version 1.0
// @description API quản lý đề thi của trường HUBT.
// @host localhost:8080
// @BasePath /

// @schemes http
func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
    config.ConnectDB()

    r := gin.Default()
    routes.RegisterExamRoutes(r)

    // Swagger UI
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
