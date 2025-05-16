package services

import (
    "context"
    "errors"
    "time"

    "github.com/dangvudev/HUBT_ExamS/config"
    "github.com/dangvudev/HUBT_ExamS/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateExam inserts a new exam into MongoDB
func CreateExam(exam models.Exam) (*models.Exam, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Generate new ObjectID
    newID := primitive.NewObjectID()
    exam.ID = newID.Hex()

    // Insert with manual ID
    _, err := config.ExamCollection.InsertOne(ctx, exam)
    if err != nil {
        return nil, err
    }

    return &exam, nil
}

// GetAllExams retrieves all exams from MongoDB
func GetAllExams() ([]models.Exam, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := config.ExamCollection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var exams []models.Exam
    for cursor.Next(ctx) {
        var exam models.Exam
        if err := cursor.Decode(&exam); err != nil {
            return nil, err
        }
        exams = append(exams, exam)
    }

    return exams, nil
}

// GetExamByID finds one exam by ID
func GetExamByID(id string) (*models.Exam, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var exam models.Exam
    err := config.ExamCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&exam)
    if err != nil {
        return nil, err
    }

    return &exam, nil
}

// DeleteExam deletes an exam by ID
func DeleteExam(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    result, err := config.ExamCollection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        return err
    }

    if result.DeletedCount == 0 {
        return errors.New("exam not found")
    }

    return nil
}
