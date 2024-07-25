package main

import (
	"linkmini/config"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const envPath = "../config/.env"

var AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

func init() {
	// load the environements file
	godotenv.Load(envPath)
}

func main() {
	// connect to the DB
	_, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	// create server
	server := gin.Default()

	// middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = AllowMethods
	server.Use(cors.New(config))

	// routes

	// run server
	server.Run(":" + os.Getenv("PORT"))
}
