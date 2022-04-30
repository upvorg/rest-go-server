package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
)

func hasLikedPost(pid uint, uid uint) bool {
	var count int64
	db.Orm.Model(&model.Like{}).Where("uid = ? and pid = ?", uid, pid).Find(&model.Like{}).Count(&count)

	return count > 0
}

func hasCollectedPost(pid uint, uid uint) bool {
	var count int64
	db.Orm.Model(&model.Collection{}).Where("uid = ? and pid = ?", uid, pid).Find(&model.Collection{}).Count(&count)

	return count > 0
}

func LikePostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID, _ := c.Get(middleware.CTX_AUTH_KEY)
	uid := uint(userID.(*middleware.AuthClaims).UserId)

	if hasLikedPost(uint(id), uid) {
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
	userID, _ := c.Get(middleware.CTX_AUTH_KEY)
	uid := uint(userID.(*middleware.AuthClaims).UserId)
	if !hasLikedPost(uint(id), uid) {
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
	userID, _ := c.Get(middleware.CTX_AUTH_KEY)
	uid := uint(userID.(*middleware.AuthClaims).UserId)
	if hasCollectedPost(uint(id), uid) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "You have collected this post.",
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
	userID, _ := c.Get(middleware.CTX_AUTH_KEY)
	uid := uint(userID.(*middleware.AuthClaims).UserId)
	if !hasCollectedPost(uint(id), uid) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "You have collected this post.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Model(&model.Collection{}).Where("uid = ? and pid = ?", uid, id).Delete(model.Collection{}).Error,
	})
}
