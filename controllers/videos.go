package controllers

import (
	"fmt"
	"net/http"
	"os"
	"rabrill/fetchers"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func CommenterVideos(c *gin.Context) {
	key := os.Getenv("YT_API_KEY")
	url := c.Query("q")
	commenterVideos, err := fetchers.FetchCommenterVideos(url, key)

	var msg Message

	if err != nil {
		msg.Success = false
		msg.Data = fmt.Sprintf("%v", err)
		c.AbortWithStatusJSON(400, msg)
	} else {
		msg.Success = true
		msg.Data = commenterVideos
		c.JSON(http.StatusOK, msg)
	}
}
