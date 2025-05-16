package controllers

import (
    "net/http"

    "github.com/dangvudev/HUBT_ExamS/models"
    "github.com/dangvudev/HUBT_ExamS/services"
    "github.com/gin-gonic/gin"
)

// UploadExam godoc
// @Summary Tải lên đề thi kèm câu hỏi
// @Description Tải lên đề thi với đầy đủ thông tin: môn học, năm học, tác giả, khoa và danh sách câu hỏi (có thể kèm hình ảnh)
// @Tags Exams
// @Accept json
// @Produce json
// @Param exam body models.Exam true "Dữ liệu đề thi đầy đủ"
// @Success 200 {object} models.Exam "Thông tin đề thi đã lưu"
// @Failure 400 {object} map[string]string "Dữ liệu gửi lên không hợp lệ hoặc thiếu trường bắt buộc"
// @Failure 500 {object} map[string]string "Lỗi server trong quá trình lưu trữ"
// @Router /exams [post]
func UploadExam(c *gin.Context) {
    var exam models.Exam

    // Bước 1: Bind JSON từ request body vào struct Exam
    if err := c.ShouldBindJSON(&exam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu gửi lên không hợp lệ: " + err.Error()})
        return
    }

    // Bước 2: Kiểm tra trường bắt buộc (ví dụ Subject và Year)
    if exam.Subject == "" || exam.Year == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Thiếu trường bắt buộc: subject hoặc year"})
        return
    }

    // Bước 3: Gọi service lưu đề thi vào database
    createdExam, err := services.CreateExam(exam)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi lưu dữ liệu: " + err.Error()})
        return
    }

    // Bước 4: Trả về kết quả đề thi đã được tạo thành công
    c.JSON(http.StatusOK, createdExam)
}

// GetAllExams godoc
// @Summary Lấy danh sách đề thi
// @Description Lấy tất cả các đề thi đã lưu
// @Tags Exams
// @Produce json
// @Success 200 {array} models.Exam
// @Failure 500 {object} map[string]string
// @Router /exams [get]
func GetAllExams(c *gin.Context) {
    exams, err := services.GetAllExams()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, exams)
}

// GetExamByID godoc
// @Summary Lấy đề thi theo ID
// @Description Lấy chi tiết đề thi từ ID MongoDB
// @Tags Exams
// @Produce json
// @Param id path string true "ID của đề thi"
// @Success 200 {object} models.Exam
// @Failure 404 {object} map[string]string
// @Router /exams/{id} [get]
func GetExamByID(c *gin.Context) {
    id := c.Param("id")
    exam, err := services.GetExamByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy đề thi"})
        return
    }
    c.JSON(http.StatusOK, exam)
}

// DeleteExam godoc
// @Summary Xóa đề thi
// @Description Xóa một đề thi dựa trên ID
// @Tags Exams
// @Produce json
// @Param id path string true "ID của đề thi"
// @Success 200 {object} map[string]string "Đã xóa thành công"
// @Failure 404 {object} map[string]string "Không tìm thấy đề thi để xóa"
// @Router /exams/{id} [delete]
func DeleteExam(c *gin.Context) {
    id := c.Param("id")
    err := services.DeleteExam(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không xóa được. Có thể đề thi không tồn tại"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Đã xóa thành công"})
}
