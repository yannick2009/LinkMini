package main

import (
	"linkmini/config"
	"linkmini/handlers"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/surrealdb/surrealdb.go"
)

var (
	envPath      = "../.env"
	AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	DB           *surrealdb.DB
)

func init() {
	//------------- LOAD THE ENV FILE ------------- //
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	//------------- DATABASE ------------- //
	err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	//------------- INIT SERVER ------------- //
	server := gin.Default()
	//------------- MIDDLEWARE ------------- //
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = AllowMethods
	server.Use(cors.New(config))

	//------------- ROUTES ------------- //
	server.POST("/create", handlers.CreateShortURLHandler)
	server.GET("/:hash", handlers.RedirectURL)
	server.GET("/stats/:hash", handlers.GetURLStats)

	//------------- RUN SERVER ------------- //
	server.Run(":" + os.Getenv("PORT"))
}
