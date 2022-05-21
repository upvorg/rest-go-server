package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
)

func IsNotLikedPost(pid uint, uid uint) bool {
	return db.Orm.Where("uid = ? and pid = ?", uid, pid).First(&model.Like{}).Error == gorm.ErrRecordNotFound
}

func IsNotCollectedPost(pid uint, uid uint) bool {
	return db.Orm.Where("uid = ? and pid = ?", uid, pid).First(&model.Collection{}).Error == gorm.ErrRecordNotFound
}

func LikePostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims).UserId

	if !IsNotLikedPost(uint(id), uid) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "You have liked this post.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Model(&model.Like{}).Create(&model.Like{
			Uid: uid,
			Pid: uint(id),
		}).Error,
	})

}

func UnlikePostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims).UserId
	if IsNotLikedPost(uint(id), uid) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "repeatedly",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Model(&model.Like{}).Where("uid = ? and pid = ?", uid, id).Delete(model.Like{}).Error,
	})
}

func CollectPostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims).UserId

	if !IsNotCollectedPost(uint(id), uid) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "You have collected this post.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Model(&model.Collection{}).Create(&model.Collection{
			Uid: uid,
			Pid: uint(id),
		}).Error,
	})
}

func UncollectPostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims).UserId

	if IsNotCollectedPost(uint(id), uid) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "You have collected this post.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Model(&model.Collection{}).Where("uid = ? and pid = ?", uid, id).Delete(model.Collection{}).Error,
	})
}

func GetLikesByUserId(c *gin.Context) {
	var posts []model.Post
	uid := c.Param("name")
	postType := c.Query("type")
	tx := db.Orm.Model(&model.Like{}).
		Select("posts.*").
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Meta").
		Where("likes.uid = ?", uid)

	if postType != "" {
		tx.Joins("left join posts on posts.id = likes.pid AND posts.type = ?", postType)
	} else {
		tx.Joins("left join posts on posts.id = likes.pid")
	}

	if err := tx.
		Joins("left join users on users.id = posts.uid").
		Joins("left join video_metas on video_metas.id = posts.id").
		Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func GetCollectionsByUserId(c *gin.Context) {
	var posts []model.Post
	uid := c.Param("name")
	postType := c.Query("type")
	tx := db.Orm.Model(&model.Collection{}).
		Select("posts.*").
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Meta").
		Where("collections.uid = ?", uid)

	if postType != "" {
		tx.Joins("left join posts on posts.id = collections.pid AND posts.type = ?", postType)
	} else {
		tx.Joins("left join posts on posts.id = collections.pid")
	}

	if err := tx.
		Joins("left join users on users.id = posts.uid").
		Joins("left join video_metas on video_metas.id = posts.id").
		Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}
