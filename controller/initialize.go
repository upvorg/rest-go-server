package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"upv.life/server/middleware"
)

func Initialize(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		auth, _ := c.Get(middleware.GinCtxAuthKey)
		c.String(200, "Hello -> %v !\n", auth.(jwt.MapClaims)["user_id"])
	})
}
