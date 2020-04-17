package main

import (
	"govote/middlewares"
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
	app.Pre(middleware.RemoveTrailingSlash())

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status} ${latency_human}\n",
	}))
	app.Use(middleware.Recover())

	admin := app.Group("/api/admin")
	admin.POST("/login", admincontroller.Login)
	admin.GET("/votes", admincontroller.ListVotes, middlewares.AuthAdmin)
	admin.POST("/votes", admincontroller.CreateVotes, middlewares.AuthAdmin)
	admin.GET("/voters", admincontroller.ListVoters, middlewares.AuthAdmin)
	admin.POST("/voters", admincontroller.GenerateVoters, middlewares.AuthAdmin)

	voter := app.Group("/api/voter")
	voter.POST("/login", votercontroller.Login)
	voter.GET("/votes", votercontroller.ListVotes, middlewares.AuthVoter)
	voter.POST("/vote", votercontroller.Vote, middlewares.AuthVoter)

	app.Static("/static", "ui/static")
	app.File("/app", "ui/app.html")
	app.File("/app/*", "ui/app.html")

	app.Start(":8080")
}
