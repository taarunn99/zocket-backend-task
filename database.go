package database

import (
    "context"
    "errors"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var client *mongo.Client

func InitDB() error {
    if err := godotenv.Load(); err != nil {
        return errors.New("error loading .env file: " + err.Error())
    }

    clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
    var err error
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return errors.New("failed to connect to MongoDB: " + err.Error())
    }

    if err = client.Ping(context.TODO(), nil); err != nil {
        return errors.New("MongoDB connection failed: " + err.Error())
    }

    DB = client.Database("taskdb")
    log.Println("✅ Connected to MongoDB")
    return nil
}
