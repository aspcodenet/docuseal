package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.HTML(http.StatusOK, "ping.html", gin.H{
		"message": "pong",
	})
}
