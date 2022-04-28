package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"upv.life/server/common"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
	"upv.life/server/service"
)

func GetPostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := service.GetPostById(uint(id))

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

	userID, _ := c.Get(middleware.CTX_AUTH_KEY)
	body.Uid = uint(userID.(*middleware.AuthClaims).UserId)
	if body.Type == "post" && body.Content == "" {
		c.AbortWithStatusJSON(http.StatusPaymentRequired, gin.H{
			"msg": "Content is required.",
		})
		return
	}

	var err error
	if body.Type == "video" {
		err = db.Orm.Create(body).Error
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

	if ok := isYourPost(c, body); !ok {
		return
	}

	var err error
	if er := db.Orm.Model(&model.Post{}).Where("id =?", body.ID).Updates(body).Error; er != nil {
		err = er
	}
	if er := db.Orm.Model(&model.VideoMeta{}).Where("pid = ?", body.ID).Updates(body.Meta).Error; er != nil {
		err = er
	}

	c.JSON(http.StatusOK, gin.H{
		"err": err,
	})
}

func DeletePostById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if ok := isYourPost(c, &model.Post{ID: uint(id)}); !ok {
		return
	}
	var err error
	err = db.Orm.Delete(&model.Post{ID: uint(id)}).Error
	if e := db.Orm.Where("pid = ?", id).Delete(&model.VideoMeta{}).Error; e != nil {
		err = e
	}
	c.JSON(http.StatusOK, gin.H{
		"err": err,
	})
}

func DeletePostsById(c *gin.Context) {
	var (
		body map[string]interface{}
		err  error
	)
	//TODO: Is it your own post?
	c.ShouldBindJSON(&body)
	ids := (body["ids"]).([]int)
	err = db.Orm.Delete(&model.Post{}, ids).Error
	if e := db.Orm.Where("pid in (?)", ids).Delete(&model.VideoMeta{}).Error; e != nil {
		err = e
	}
	c.JSON(http.StatusOK, gin.H{
		"err": err,
	})
}

func isYourPost(c *gin.Context, body *model.Post) bool {
	post, _ := service.GetPostById(body.ID)
	if post == nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"err": "Post not found.",
		})
		return false
	}
	ctxUser := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	if !(common.IsRoot(ctxUser.Level) || common.IsAdmin(ctxUser.Level)) && (post.Uid != ctxUser.UserId || body.Status != 0) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"err": "Forbidden.",
		})
		return false
	}
	return true
}

//TODO: `UpdatedAt` should not be updated.
func UpdatePostPv(c *gin.Context) {
	err := db.Orm.Model(&model.Post{}).Where("id = ?", c.Param("id")).Update("pv", gorm.Expr("pv + 1")).Error
	c.JSON(http.StatusOK, gin.H{
		"err": err,
	})

}
