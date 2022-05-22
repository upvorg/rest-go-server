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

func GetComments(c *gin.Context) {
	var comments []model.Comment

	ctxUser := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	tx := db.Orm.Model(&model.Comment{}).
		Scopes(model.Paginate(c)).
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Post").
		Preload("Children").
		Preload("Children.Creator").
		Where("(parent_id IS NULL OR parent_id = 0)")

	if !(common.IsAdmin(ctxUser.Level) && !common.IsRoot(ctxUser.Level)) {
		tx = tx.Joins("LEFT JOIN posts ON posts.id = comments.pid").
			Where("posts.uid = ?", ctxUser.UserId)
	}

	err := tx.Order("created_at DESC").Find(&comments).Error
	c.JSON(http.StatusOK, gin.H{
		"err":  err,
		"data": comments,
	})
}

func GetCommentsByPostId(c *gin.Context) {
	pid := c.Param("id")
	var comments []model.Comment
	err := db.Orm.Model(&model.Comment{}).
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Children").
		Preload("Children.Creator").
		Where("pid = ? AND (parent_id IS NULL OR parent_id = 0)", pid).
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
	body.Pid = c.Param("id")
	if body.Content == "" {
		c.AbortWithStatusJSON(http.StatusPaymentRequired, gin.H{
			"err": "Content is required.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err":  db.Orm.Create(body).Error,
		"data": body,
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
