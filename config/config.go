package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var ExamCollection *mongo.Collection

const (
    connectionString       = "mongodb://localhost:27017"
    dbName                 = "EXAMS_M"
    collectionName         = "exams"
)

// Gọi hàm này một lần trong main.go
func ConnectDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(connectionString)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("MongoDB connection error:", err)
    }

    MongoClient = client
    ExamCollection = client.Database(dbName).Collection(collectionName)
}
