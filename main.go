package main

import (
	"log"
	"net/http"
	"os"
	"rabrill/fetchers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getStuff(c *gin.Context) {
	key := os.Getenv("YT_API_KEY")
	url := c.Query("q")
	commenterVideos := fetchers.FetchCommenterVideos(url, key)

	c.IndentedJSON(http.StatusOK, commenterVideos)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	router := gin.Default()
	router.GET("/", getStuff)
	router.Run("localhost:8585")
}
