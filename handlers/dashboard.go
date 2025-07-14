package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	email, _ := c.Get("email")
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Email": email,
	})
}
