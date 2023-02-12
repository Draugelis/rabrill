package controllers

import (
	"net/http"
	"os"
	"rabrill/fetchers"

	"github.com/gin-gonic/gin"
)

func CommenterVideos(c *gin.Context) {
	key := os.Getenv("YT_API_KEY")
	url := c.Query("q")
	commenterVideos := fetchers.FetchCommenterVideos(url, key)

	c.IndentedJSON(http.StatusOK, commenterVideos)
}
