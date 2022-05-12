package controller

import (
	"github.com/gin-gonic/gin"
	"upv.life/server/config"
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
	// :name is id!!
	api.GET("/user/:name/likes", middleware.ShouldBeLogin(), GetLikesByUserId)
	api.GET("/user/:name/collections", middleware.ShouldBeLogin(), GetCollectionsByUserId)
	api.GET("/user/stat", middleware.ShouldBeLogin(), GetUserStat)
	api.GET("/user/post/activity", middleware.ShouldBeLogin(), GetUserPostActivity)

	api.GET("/post/:id", GetPostById)
	api.GET("/posts", GetPostsByMetaType)
	api.GET("/post/:id/pv", UpdatePostPv)
	api.GET("/tag/:tag/posts", GetPostByTag)
	api.GET("/post/ranking", GetPostRanking)
	api.GET("/post/day/ranking", GetPostDayRanking)
	api.GET("/post/month/ranking", GetPostMonthRanking)
	api.GET("/posts/recommends", GetRecommendPosts)
	api.POST("/post", middleware.ShouldBeLogin(), CreatePost)
	api.PUT("/post/:id", middleware.ShouldBeLogin(), UpdatePost)
	api.DELETE("/post/:id", middleware.ShouldBeLogin(), DeletePostById)
	api.DELETE("/posts", middleware.ShouldBeLogin(), DeletePostsById)
	api.GET("/post/week", GetVideosUpdateOnWeek)
	api.POST("/post/:id/review", middleware.ShouldBeLogin(), ReviewPost)

	api.POST("/like/post/:id", middleware.ShouldBeLogin(), LikePostById)
	api.DELETE("/like/post/:id", middleware.ShouldBeLogin(), UnlikePostById)
	api.POST("/collect/post/:id", middleware.ShouldBeLogin(), CollectPostById)
	api.DELETE("/collect/post/:id", middleware.ShouldBeLogin(), UncollectPostById)

	api.GET("/comments", GetComments)
	api.GET("/post/:id/comments", GetCommentsByPostId)
	api.POST("/post/:id/comment", middleware.ShouldBeLogin(), CreateComment)
	api.DELETE("/comment/:id", middleware.ShouldBeLogin(), DeleteCommentById)

	api.GET("/post/:id/videos", GetVideosByPostId)
	api.POST("/post/:id/video", middleware.ShouldBeLogin(), CreateVideo)
	api.DELETE("/video/:id", middleware.ShouldBeLogin(), DeleteVideoById)
	api.PUT("/video/:id", middleware.ShouldBeLogin(), UpdateVideoById)

	api.GET("/tags", GetTags)
	api.POST("/tag", middleware.ShouldBeLogin(), CreateTag)
	api.DELETE("/tag/:id", middleware.ShouldBeLogin(), DeleteTag)

	api.GET("/feedbacks", GetFeedbacks)
	api.POST("/feedback", CreateFeedback)

	api.POST("/upload", FileUploader)
	api.POST("/upload/image", SMMSImageUploder)

	if config.AppMode == "debug" {
		r.Static("/uploads", "./uploads")
	}
}
