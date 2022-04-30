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

	user := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	body.Uid = uint(user.UserId)
	if body.Type == "post" && body.Content == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "Content is required.",
		})
		return
	}

	if common.IsUser(user.Level) && body.Type == "video" {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "You can't create video post.",
		})
		return
	}

	if !common.IsRoot(user.Level) {
		body.Status = 0
		body.IsPined = 0
		body.IsRecommend = 0
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

	user := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	if !common.IsRoot(user.Level) {
		body.Status = 0
		body.IsPined = 0
		body.IsRecommend = 0
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
	err = db.Orm.Delete(&model.Post{}, id).Error
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
	c.ShouldBindJSON(&body)
	ids := (body["ids"]).([]int)
	ctxUser := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	err = db.Orm.Where("uid = ?", ctxUser.UserId).Delete(&model.Post{}, ids).Error
	if e := db.Orm.Where("pid in (?)", ids).
		Where("posts.deleted_at IS NOT NULL").
		Joins("left join posts on posts.id IN (?)", ids).
		Delete(&model.VideoMeta{}).Error; e != nil {
		err = e
	}
	c.JSON(http.StatusOK, gin.H{
		"err": err,
	})
}

func isYourPost(c *gin.Context, body *model.Post) bool {
	post, _ := service.GetPostById(body.ID)
	if post == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
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

func UpdatePostPv(c *gin.Context) {
	pr := &model.PostRanking{}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, _ := service.GetPostById(uint(id))
	if post == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"err": "Post not found.",
		})
		return
	}
	if err := db.Orm.Model(&model.PostRanking{}).
		Where("pid = ?", c.Param("id")).
		Where("to_days(hits_at) = to_days(now())").
		First(&pr).Error; err == gorm.ErrRecordNotFound {
		db.Orm.Model(&model.PostRanking{}).Create(&model.PostRanking{
			Pid:  uint(id),
			Hits: 1,
		})
		return
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err,
		})
		return
	}

	err := db.Orm.Model(&model.PostRanking{}).
		Where("pid = ?", c.Param("id")).
		Where("to_days(hits_at) = to_days(now())").
		Update("hits", gorm.Expr("hits + 1")).Error

	c.JSON(http.StatusOK, gin.H{
		"err": err,
	})

}

// https://www.cnblogs.com/yiyunkeji/p/7217738.html
func GetPostDayRanking(c *gin.Context) {
	posts := []*model.Post{}
	if err := db.Orm.Model(&model.Post{}).
		Preload("Creator").
		Preload("Meta").
		Select("posts.*, SUM(post_rankings.hits) as Hits").
		Joins("left join post_rankings on post_rankings.pid = posts.id and to_days(post_rankings.hits_at) = to_days(now())").
		Limit(20).
		Order("hits desc").
		Group("posts.id").
		Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func GetPostMonthRanking(c *gin.Context) {
	posts := []*model.Post{}
	if err := db.Orm.Model(&model.Post{}).
		Preload("Creator").
		Preload("Meta").
		Select("posts.*, SUM(post_rankings.hits) as Hits").
		Joins("left join post_rankings on post_rankings.pid = posts.id and DATE_FORMAT(post_rankings.hits_at,'%Y%m') = DATE_FORMAT(CURDATE(),'%Y%m')").
		Limit(30).
		Order("hits desc").
		Group("posts.id").
		Find(&posts).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func GetPostRanking(c *gin.Context) {
	posts := []*model.Post{}
	if err := db.Orm.Model(&model.Post{}).
		Preload("Creator").
		Preload("Meta").
		Select("posts.*, SUM(post_rankings.hits) as Hits").
		Joins("left join post_rankings on post_rankings.pid = posts.id").
		Limit(50).
		Order("hits desc").
		Group("posts.id").
		Find(&posts).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}
