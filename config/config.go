package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var ExamCollection *mongo.Collection


// Gọi hàm này một lần trong main.go
func ConnectDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    connectionString := os.Getenv("DB_CONNECTIONSTRING")   // phải đúng tên biến
    dbName           := os.Getenv("DB_NAME")
    collectionName   := os.Getenv("DB_COLLECTIONNAME")

    fmt.Println("DB_CONNECTIONSTRING: ",connectionString)
    fmt.Println("DB_NAME: ",dbName)
    fmt.Println("DB_COLLECTIONNAME: ",collectionName)

    clientOptions := options.Client().ApplyURI(connectionString)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("MongoDB connection error:", err)
    }

    MongoClient = client
    ExamCollection = client.Database(dbName).Collection(collectionName)
}
