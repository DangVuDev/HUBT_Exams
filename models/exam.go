package models

type Exam struct {
    ID        string    `json:"id" bson:"_id,omitempty"`
    Subject   string    `json:"subject" bson:"subject"`           // Tên môn học
    Author    string    `json:"author" bson:"author"`             // Tác giả
    Year      int       `json:"year" bson:"year"`                 // Năm học
    Faculty   string    `json:"faculty" bson:"faculty"`           // Tên khoa
    Questions []Question `json:"questions" bson:"questions"`      // Danh sách câu hỏi
}

type Question struct {
    Title       string     `json:"title" bson:"title"`             // Tên câu hỏi
    Content     string     `json:"content" bson:"content"`         // Đề bài
    Options     []string   `json:"options" bson:"options"`         // Bộ đáp án
    Answer      string     `json:"answer" bson:"answer"`           // Đáp án đúng
    Image       *ImageData `json:"image,omitempty" bson:"image,omitempty"` // Hình ảnh nếu có
}

type ImageData struct {
    Base64   string `json:"base64" bson:"base64"`   // Dữ liệu base64
    FileType string `json:"fileType" bson:"fileType"` // Đuôi file, ví dụ: jpg, png
}