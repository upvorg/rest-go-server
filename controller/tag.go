package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"upv.life/server/common"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
)

func CreateTag(c *gin.Context) {
	tag := &model.Tag{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": common.Translate(err),
		})
		return
	}
	tags, _ := getDBTags()

	if common.HasTag(tags, tag.Name) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "tag already exists",
		})
		return
	}

	tag.Uid = c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims).UserId
	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Create(tag).Error,
	})
}

func GetTags(c *gin.Context) {
	tags, err := getDBTags()
	c.JSON(http.StatusOK, gin.H{
		"err":  err,
		"data": tags,
	})
}

////////////////////////////////////////////////////////////////////////////////

func getDBTags() (string, error) {
	var tags string
	if err := db.Orm.Model(&model.Tag{}).Select("group_concat(tags.name SEPARATOR ' ') as tags").Find(&tags).Error; err != nil {
		return "", err
	} else {
		return tags, nil
	}
}
