package controller

import (
	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/register", Register)
	api.POST("/login", Login)

	api.GET("/user/:name", GetUserByName)
	api.PUT("/user/:name", UpdateUserByName)
	api.GET("/users", GetUsers)

	api.POST("/post", CreatePost)
	api.GET("/post/:id", GetPostById)
	api.PUT("/post/:id", UpdatePost)
	api.DELETE("/post/:id", DeletePostById)
	api.GET("/posts", GetPostsByMetaType)
	api.GET("/:tag/posts", GetPostByTag)
	api.GET("/posts/recommends", GetRecommendPosts)
	api.DELETE("/posts", DeletePostsById)

	api.POST("/like/post", LikePostById)
	api.DELETE("/like/post", UnlikePostById)
	api.POST("/collect/post", CollectPostById)
	api.DELETE("/collect/post", UncollectPostById)

	api.GET("/post/:id/comments", GetCommentsByPostId)
	api.POST("/post/:id/comment", CreateComment)
}
