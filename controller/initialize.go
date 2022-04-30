package controller

import (
	"github.com/gin-gonic/gin"
	"upv.life/server/middleware"
)

func Initialize(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/register", Register)
	api.POST("/login", Login)

	api.GET("/users", GetUsers)
	api.GET("/user/:name", GetUserByName)
	api.GET("/user", middleware.ShouldBeLogin(), GetUser)
	api.PUT("/user/:name", middleware.ShouldBeLogin(), UpdateUserByName)

	api.GET("/post/:id", GetPostById)
	api.GET("/:tag/posts", GetPostByTag)
	api.GET("/posts", GetPostsByMetaType)
	api.GET("/post/:id/pv", UpdatePostPv)
	api.GET("/post/ranking", GetPostRanking)
	api.GET("/post/day/ranking", GetPostDayRanking)
	api.GET("/post/month/ranking", GetPostMonthRanking)
	api.GET("/posts/recommends", GetRecommendPosts)
	api.POST("/post", middleware.ShouldBeLogin(), CreatePost)
	api.PUT("/post/:id", middleware.ShouldBeLogin(), UpdatePost)
	api.DELETE("/post/:id", middleware.ShouldBeLogin(), DeletePostById)
	api.DELETE("/posts", middleware.ShouldBeLogin(), DeletePostsById)

	api.POST("/like/post/:id", middleware.ShouldBeLogin(), LikePostById)
	api.DELETE("/like/post/:id", middleware.ShouldBeLogin(), UnlikePostById)
	api.POST("/collect/post/:id", middleware.ShouldBeLogin(), CollectPostById)
	api.DELETE("/collect/post/:id", middleware.ShouldBeLogin(), UncollectPostById)

	api.GET("/post/:id/comments", GetCommentsByPostId)
	api.POST("/post/:id/comment", middleware.ShouldBeLogin(), CreateComment)
	// api.DELETE("/comment/:id", middleware.ShouldBeLogin(), DeleteCommentById)

	api.GET("/post/:id/videos", GetVideosByPostId)
	api.POST("/post/:id/video", middleware.ShouldBeLogin(), CreateVideo)
	api.DELETE("/video/:id", middleware.ShouldBeLogin(), DeleteVideoById)
	api.PUT("/post/:id/video", middleware.ShouldBeLogin(), UpdateVideoById)

	api.GET("/feedbacks", GetFeedbacks)
	api.POST("/feedback", CreateFeedback)
}
