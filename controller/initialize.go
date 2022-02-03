package controller

import (
	"github.com/gin-gonic/gin"
	"upv.life/server/middleware"
)

func Initialize(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		auth, _ := c.Get(middleware.GinCtxAuthKey)
		c.JSON(200, gin.H{
			"message": "Hello World",
			"auth":    auth,
		})
	})
}
