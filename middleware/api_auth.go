package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/repo/models"
	"gorm.io/gorm"
)

func APIAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Auth-Token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		db := c.MustGet("db").(*gorm.DB)
		var accessToken models.AccessToken
		db.Where("token = ?", token).First(&accessToken)

		if accessToken.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		var user models.User
		db.First(&user, accessToken.UserID)

		c.Set("user", user)
		c.Next()
	}
}
