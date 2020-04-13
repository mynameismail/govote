package main

import (
	"govote/models"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	admincontroller "govote/controllers/admin"
	votercontroller "govote/controllers/voter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	models.Init()

	app := echo.New()

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status} ${latency_human}\n",
	}))
	app.Use(middleware.Recover())

	admin := app.Group("/api/admin")
	admin.POST("/login", admincontroller.Login)
	admin.GET("/votes", admincontroller.ListVotes)
	admin.POST("/votes", admincontroller.CreateVotes)
	admin.GET("/voters", admincontroller.ListVoters)
	admin.POST("/voters", admincontroller.CreateVoters)

	voter := app.Group("/api/voter")
	voter.POST("/login", votercontroller.Login)
	voter.POST("/vote", votercontroller.Vote)

	app.Start(":8080")
}
