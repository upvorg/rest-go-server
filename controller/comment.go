package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"upv.life/server/common"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
)

func hasLikedPost(pid uint, uid uint) bool {
	var count int64
	db.Orm.Model(&model.Likes{}).Where("uid = ? and pid = ?", uid, pid).Find(&model.Likes{}).Count(&count)

	return count > 0
}

func hasCollectedPost(pid uint, uid uint) bool {
	var count int64
	db.Orm.Model(&model.Collects{}).Where("uid = ? and pid = ?", uid, pid).Find(&model.Collects{}).Count(&count)

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
		"err": db.Orm.Model(&model.Likes{}).Create(&model.Likes{
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
		"err": db.Orm.Model(&model.Likes{}).Where("uid = ? and pid = ?", uid, id).Delete(model.Likes{}).Error,
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
		"err": db.Orm.Model(&model.Collects{}).Create(&model.Collects{
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
		"err": db.Orm.Model(&model.Collects{}).Where("uid = ? and pid = ?", uid, id).Delete(model.Collects{}).Error,
	})
}

func GetCommentsByPostId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var comments []model.Comments
	db.Orm.Model(&model.Comments{}).Where("pid = ?", id).Find(&comments)
	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func CreateComment(c *gin.Context) {
	body := &model.Comments{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": common.Translate(err),
		})
		return
	}
	userID, _ := c.Get(middleware.CTX_AUTH_KEY)
	uid := uint(userID.(*middleware.AuthClaims).UserId)
	body.Uid = uid
	pid, _ := strconv.Atoi(c.Param("id"))
	body.Pid = uint(pid)
	if body.Content == "" {
		c.AbortWithStatusJSON(http.StatusPaymentRequired, gin.H{
			"msg": "Content is required.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Debug().Create(body).Error,
	})
}

func DeleteCommentById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Delete(&model.Comments{}, id).Error,
	})
}
