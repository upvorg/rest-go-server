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
	db.Orm.Model(&model.Comment{}).Where("pid = ?", id).Find(&comments)
	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func CreateComment(c *gin.Context) {
	body := &model.Comment{}
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
		"err": db.Orm.Create(body).Error,
	})
}

func DeleteCommentById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Delete(&model.Comment{}, id).Error,
	})
}
