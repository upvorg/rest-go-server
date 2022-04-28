package controller

import (
	"github.com/gin-gonic/gin"
	"upv.life/server/middleware"
)

func Initialize(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/register", Register)
	api.POST("/login", Login)

	api.GET("/user/:name", GetUserByName)
	api.PUT("/user/:name", middleware.ShouldBeLogin(), UpdateUserByName)
	api.GET("/users", GetUsers)

	api.POST("/post", middleware.ShouldBeLogin(), CreatePost)
	api.GET("/post/:id", GetPostById)
	api.PUT("/post/:id", middleware.ShouldBeLogin(), UpdatePost)
	api.DELETE("/post/:id", middleware.ShouldBeLogin(), DeletePostById)
	api.GET("/posts", GetPostsByMetaType)
	api.GET("/:tag/posts", GetPostByTag)
	api.GET("/posts/recommends", GetRecommendPosts)
	api.DELETE("/posts", middleware.ShouldBeLogin(), DeletePostsById)

	api.POST("/like/post/:id", middleware.ShouldBeLogin(), LikePostById)
	api.DELETE("/like/post/:id", middleware.ShouldBeLogin(), UnlikePostById)
	api.POST("/collect/post/:id", middleware.ShouldBeLogin(), CollectPostById)
	api.DELETE("/collect/post/:id", middleware.ShouldBeLogin(), UncollectPostById)

	api.GET("/post/:id/comments", GetCommentsByPostId)
	api.POST("/post/:id/comment", middleware.ShouldBeLogin(), CreateComment)
	api.DELETE("/comment/:id", middleware.ShouldBeLogin(), DeleteCommentById)

	api.GET("/post/:id/videos", GetVideosByPostId)
	api.POST("/post/:id/video", middleware.ShouldBeLogin(), CreateVideo)
	api.DELETE("/video/:id", middleware.ShouldBeLogin(), DeleteVideoById)
}
