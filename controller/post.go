package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"upv.life/server/common"
	"upv.life/server/model"
	"upv.life/server/service"
)

func GetPostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := service.GetPostById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": post,
		})
	}
}

func GetPostsByMetaType(c *gin.Context) {
	body := model.Meta{}
	if err := c.ShouldBindQuery(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": common.Translate(err),
		})
		return
	}

	posts, e := service.GetPostsByMetaType(body, c)
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
		"err":  e,
	})
}

func GetPostByTag(c *gin.Context) {
	posts, e := service.GetPostsByTag(c.Param("tag"), c)
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
		"err":  e,
	})
}

func GetPinedPosts(c *gin.Context) {
	posts, e := service.GetPinedPosts()
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
		"err":  e,
	})
}
