package routes

import (
    "github.com/gin-gonic/gin"
    "backend/database"
    "backend/models"
    "backend/auth"
    "net/http"
    "encoding/json"
    "bytes"
    "os"
    "context"
    "go.mongodb.org/mongo-driver/bson"
)

func AI_TaskSuggestion(title string) string {
    body := map[string]string{
        "prompt": "Suggest a breakdown for this task: " + title,
    }
    jsonData, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "https://api.gemini.com/v1/generate", bytes.NewBuffer(jsonData))
    req.Header.Set("Authorization", "Bearer "+os.Getenv("GEMINI_API_KEY"))

    client := &http.Client{}
    res, _ := client.Do(req)
    defer res.Body.Close()

    var result map[string]interface{}
    json.NewDecoder(res.Body).Decode(&result)
    return result["text"].(string)
}

func RegisterRoutes(r *gin.Engine) {
    r.POST("/login", func(c *gin.Context) {
        var user models.User
        c.BindJSON(&user)
        token, _ := auth.GenerateToken(user.Email)
        c.JSON(http.StatusOK, gin.H{"token": token})
    })

    r.POST("/tasks", auth.AuthMiddleware(), func(c *gin.Context) {
        var task models.Task
        c.BindJSON(&task)
        task.Status = "Pending"
        aiSuggestion := AI_TaskSuggestion(task.Title)

        collection := database.DB.Collection("tasks")
        _, err := collection.InsertOne(context.TODO(), task)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"task": task, "ai_suggestion": aiSuggestion})
    })

    r.GET("/tasks", auth.AuthMiddleware(), func(c *gin.Context) {
        collection := database.DB.Collection("tasks")
        cursor, err := collection.Find(context.TODO(), bson.M{})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tasks"})
            return
        }

        var tasks []models.Task
        cursor.All(context.TODO(), &tasks)
        c.JSON(http.StatusOK, gin.H{"tasks": tasks})
    })
}