package controller

import (
	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/register", Register)
	api.POST("/login", Login)

	api.GET("/user/:id", GetUserInfo)
	api.GET("/users", GetUsers)
}
