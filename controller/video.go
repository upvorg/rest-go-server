package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"upv.life/server/common"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
	"upv.life/server/service"
)

func GetVideosByPostId(c *gin.Context) {
	var (
		postID uint64
		videos []*model.Video
	)
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "post id is invalid"})
		return
	}
	if err := db.Orm.Where("pid = ?", postID).Order("episode DESC").Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": videos})
}

func CreateVideo(c *gin.Context) {
	var (
		video model.Video
	)
	if err := c.BindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	postID := c.Param("id")
	if p, _ := service.GetSimplePostByID((postID)); p == nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "post not found"})
		return
	}
	video.Pid = postID
	video.Uid = uint(c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims).UserId)
	if err := db.Orm.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": video})
}

func DeleteVideoById(c *gin.Context) {
	videoID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if !isYourVideo(c, uint(videoID)) {
		return
	}

	if err := db.Orm.Where("id = ?", videoID).Delete(&model.Video{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": err})
	}
}

func UpdateVideoById(c *gin.Context) {
	var (
		videoID uint64
		video   = &model.Video{}
	)
	if err := c.BindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	videoID, _ = strconv.ParseUint(c.Param("id"), 10, 64)
	if !isYourVideo(c, uint(videoID)) {
		return
	}

	if err := db.Orm.Model(&model.Video{}).Where("id = ?", videoID).Updates(video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": video})
}

////////////////////////////////////////////////////////////////////////////////

func isYourVideo(c *gin.Context, vid uint) bool {
	var video model.Video
	if err := db.Orm.Model(&model.Video{}).Where("id = ?", vid).First(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return false
	}
	ctxUser := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	if !(common.IsRoot(ctxUser.Level) || common.IsAdmin(ctxUser.Level)) && (video.Uid != ctxUser.UserId) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"err": "Forbidden."})
		return false
	}
	return true
}
