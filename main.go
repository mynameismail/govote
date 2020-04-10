package main

import "log"
import "os"
import "github.com/gin-gonic/gin"
import "github.com/joho/godotenv"

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }
    
    username := os.Getenv("APP_USERNAME")
    password := os.Getenv("APP_PASSWORD")
    
    log.Println(username + " " + password)
    
    app := gin.Default()
    app.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "OK"})
    })
    app.Run()
}
