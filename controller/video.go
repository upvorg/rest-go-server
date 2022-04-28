package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
	"upv.life/server/service"
)

func GetVideosByPostId(c *gin.Context) {
	var (
		postID uint64
		videos []*model.Videos
	)
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post id is invalid"})
		return
	}
	if err := db.Orm.Where("pid = ?", postID).Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": videos})
}

func CreateVideo(c *gin.Context) {
	var (
		postID uint64
		video  model.Videos
	)
	if err := c.BindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postID, _ = strconv.ParseUint(c.Param("id"), 10, 64)
	if p, _ := service.GetPostById(uint(postID)); p == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post not found"})
		return
	}
	userID, _ := c.Get(middleware.CTX_AUTH_KEY)
	video.Pid = uint(postID)
	video.Uid = uint(userID.(*middleware.AuthClaims).UserId)
	if err := db.Orm.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": video})
}

func DeleteVideoById(c *gin.Context) {
	videoID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := db.Orm.Where("id = ?", videoID).Delete(&model.Videos{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": err})
	}
}

func UpdateVideoById(c *gin.Context) {
	var (
		videoID uint64
		video   = &model.Videos{}
	)
	if err := c.BindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	videoID, _ = strconv.ParseUint(c.Param("id"), 10, 64)

	if err := db.Orm.Model(&model.Videos{}).Where("id = ?", videoID).Updates(video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": video})
}
