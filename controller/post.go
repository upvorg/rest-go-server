package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"upv.life/server/common"
	"upv.life/server/db"
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

func GetRecommendPosts(c *gin.Context) {
	posts, e := service.GetRecommendPosts()
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
		"err":  e,
	})
}

func CreatePost(c *gin.Context) {
	body := &model.Post{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": common.Translate(err),
		})
		return
	}
	// userID, _ := c.Get(middleware.CTX_AUTH_KEY)
	// body.Uid = uint(userID.(middleware.AuthClaims).UserId)

	if body.Type == "post" && body.Content == "" {
		c.AbortWithStatusJSON(http.StatusPaymentRequired, gin.H{
			"msg": "Content is required.",
		})
		return
	}

	var err error
	if body.Type == "video" {
		err = service.CreatePost(body)
	} else {
		err = db.Orm.Omit(clause.Associations).Create(body).Error
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": body,
		})
	}
}

func UpdatePost(c *gin.Context) {
	body := &model.Post{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": common.Translate(err),
		})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	body.ID = uint(id)

	var err error
	if er := db.Orm.Debug().Model(&model.Post{}).Where("id =?", body.ID).Updates(body).Error; er != nil {
		err = er
	}
	if er := db.Orm.Model(&model.VideoMetas{}).Where("pid = ?", body.ID).Updates(body.Meta).Error; er != nil {
		err = er
	}

	c.JSON(http.StatusOK, gin.H{
		"err": err,
	})
}

func DeletePostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Debug().Delete(&model.Post{ID: uint(id)}).Error,
	})
}

func DeletePostsById(c *gin.Context) {
	var body map[string]interface{}
	c.ShouldBindJSON(&body)
	ids := (body["ids"]).([]int)
	c.JSON(http.StatusOK, gin.H{
		"err": db.Orm.Delete(&model.Post{}, ids).Error,
	})
}
