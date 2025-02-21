package main

import (
    "backend/database"
    "backend/routes"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
)

func main() {
    // ✅ Load environment variables from .env with proper error handling
    if err := godotenv.Load(); err != nil {
        log.Fatalf("❌ Failed to load .env file: %v", err)
    }
    
    // ✅ Initialize database with error handling
    if err := database.InitDB(); err != nil {
        log.Fatalf("❌ Database initialization failed: %v", err)
    }

    // ✅ Initialize Gin router
    r := gin.Default()
    routes.RegisterRoutes(r)

    log.Println("🚀 Server running on http://localhost:8080")

    // ✅ Properly handle server startup errors
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("❌ Failed to start server: %v", err)
    }
}
