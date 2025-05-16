package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/dangvudev/HUBT_ExamS/controllers"
)

func RegisterExamRoutes(r *gin.Engine) {
    r.Static("/uploads", "./uploads")

    exams := r.Group("/exams")
    {
        exams.POST("/", controllers.UploadExam)
        exams.GET("/", controllers.GetAllExams)
        exams.GET("/:id", controllers.GetExamByID)
        exams.DELETE("/:id", controllers.DeleteExam)
    }
}
