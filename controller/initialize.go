package controller

import (
	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			// auth, _ := c.Get(middleware.GinCtxAuthKey)
			c.JSON(200, gin.H{
				"message": "Hello World",
				// "auth":    auth.(middleware.AuthClaims),
			})
		})
	}

}
