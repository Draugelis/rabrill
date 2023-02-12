package main

import (
	"log"
	"os"
	"rabrill/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	router := gin.Default()
	router.GET("/", controllers.CommenterVideos)
	router.Run("localhost:8585")
}
