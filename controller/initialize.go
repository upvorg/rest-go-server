package controller

import (
	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/register", Register)
}
