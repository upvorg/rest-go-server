package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	tag.Name = strings.Trim(tag.Name, " ")
	if IsTagExistByName(tag.Name) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "tag already exists",
		})
		return
	}

	tag.Uid = c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims).UserId
	c.JSON(http.StatusOK, gin.H{
		"err":  db.Orm.Create(tag).Error,
		"data": tag,
	})
}

func GetTags(c *gin.Context) {
	var tags []model.Tag
	err := db.Orm.Find(&tags).Error

	c.JSON(http.StatusOK, gin.H{
		"err":  err,
		"data": tags,
	})
}

func DeleteTag(c *gin.Context) {
	uid := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims).UserId
	tagId := c.Param("id")

	var tag model.Tag
	if err := db.Orm.Where("id = ?", tagId).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "tag not found",
		})
		return
	}

	if tag.Uid != uid {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "Forbidden",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Delete(&tag).Error,
	})
}

func IsTagExistByName(name string) bool {
	return db.Orm.Where("name = ?", name).First(&model.Tag{}).Error != gorm.ErrRecordNotFound
}

func getDBTags() (string, error) {
	var tags string
	if err := db.Orm.Model(&model.Tag{}).Select("group_concat(tags.name SEPARATOR ' ') as tags").Find(&tags).Error; err != nil {
		return "", err
	} else {
		return tags, nil
	}
}
