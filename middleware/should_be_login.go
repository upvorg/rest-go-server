package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShouldBeLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(GinCtxAuthKey); exists {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "You are Unauthorized",
			})
		}
	}
}
