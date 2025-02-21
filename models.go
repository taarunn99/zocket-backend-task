package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User model
type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Email    string             `bson:"email"`
    Password string             `bson:"password"`
}

// Task model
type Task struct {
    ID         primitive.ObjectID `bson:"_id,omitempty"`
    Title      string             `bson:"title"`
    Status     string             `bson:"status,omitempty"`
    AssignedTo string             `bson:"assigned_to,omitempty"`
}
