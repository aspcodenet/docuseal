package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Landing(c *gin.Context) {
	c.HTML(http.StatusOK, "landing.html", nil)
}
