package main

import (
    "github.com/gin-gonic/gin"
    "github.com/dangvudev/HUBT_ExamS/config"
    "github.com/dangvudev/HUBT_ExamS/routes"

    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "github.com/dangvudev/HUBT_ExamS/docs"
)

// @title HUBT Exam Storage API
// @version 1.0
// @description API quản lý đề thi của trường HUBT.
// @host localhost:8080
// @BasePath /

// @schemes http
func main() {
    config.ConnectDB()

    r := gin.Default()
    routes.RegisterExamRoutes(r)

    // Swagger UI
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run(":8080")
}
