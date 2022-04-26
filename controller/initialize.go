package controller

import (
	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/register", Register)
	api.POST("/login", Login)

	api.GET("/user/:id", GetUserById)
	api.GET("/users", GetUsers)

	api.GET("/post/:id", GetPostById)
	api.GET("/posts", GetPostsByMetaType)
	api.GET("/:tag/posts", GetPostByTag)
	api.GET("/posts/pined", GetPinedPosts)
}
