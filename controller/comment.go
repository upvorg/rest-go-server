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

func GetCommentsByPostId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var comments []model.Comment
	err := db.Orm.Model(&model.Comment{}).
		Preload("Creator").
		Preload("Children").
		Preload("Children.Creator").
		Where("pid = ? AND parent_id IS NULL", id).
		Order("created_at DESC").Find(&comments).Error
	c.JSON(http.StatusOK, gin.H{
		"data": comments,
		"err":  err,
	})
}

func CreateComment(c *gin.Context) {
	body := &model.Comment{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": common.Translate(err),
		})
		return
	}
	body.Uid = uint(c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims).UserId)
	pid, _ := strconv.Atoi(c.Param("id"))
	body.Pid = uint(pid)
	if body.Content == "" {
		c.AbortWithStatusJSON(http.StatusPaymentRequired, gin.H{
			"err": "Content is required.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Create(body).Error,
	})
}

func DeleteCommentById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	comment := &model.Comment{}
	if err := db.Orm.Where("id = ?", id).First(comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	ctxUser := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	if !(common.IsRoot(ctxUser.Level) || common.IsAdmin(ctxUser.Level)) && (comment.Uid != ctxUser.UserId) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"err": "Forbidden."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Delete(&model.Comment{}, id).Error,
	})
}
