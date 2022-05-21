package controller

import (
	"fmt"
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
	var post *model.Post
	var err error
	if user, exists := c.Get(middleware.CTX_AUTH_KEY); exists {
		post, err = service.GetPostById(uint(id), user.(*middleware.AuthClaims).UserId, user.(*middleware.AuthClaims).Level)
	} else {
		post, err = service.GetPostById(uint(id), 0, 0)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
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
			"err": common.Translate(err),
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
	size, _ := strconv.Atoi(c.Query("page_size"))
	posts, e := service.GetRecommendPosts(size)
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
		"err":  e,
	})
}

func CreatePost(c *gin.Context) {
	body := &model.Post{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": common.Translate(err),
		})
		return
	}

	user := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	body.Uid = uint(user.UserId)
	if !checkPostField(c, body, user) {
		return
	}

	var err error
	if body.Type == "video" && body.IsOriginal == 2 {
		err = db.Orm.Omit(clause.Associations).Create(body).Error
	} else {
		err = db.Orm.Create(body).Error
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
			"err": common.Translate(err),
		})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	body.ID = uint(id)

	if ok := isYourPost(c, body); !ok {
		return
	}

	user := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	if !checkPostField(c, body, user) {
		return
	}

	selects := []string{
		"Cover",
		"Title",
		"Content",
		"Tags",
	}
	tx := db.Orm.Model(&model.Post{}).Where("id = ?", body.ID)
	var err error

	if !common.IsAdmin(user.Level) || !common.IsRoot(user.Level) {
		if er := tx.Select(selects).Updates(body).Error; er != nil {
			err = er
		}
	} else {
		// TODO: GET Columns no ZERO value. "" | 0 | nil
		// eg: getNotZeroColumns(struct, []string{})
		//     getNotZeroColumns(struct)
		if body.Status != 0 {
			selects = append(selects, "Status")
		}
		if body.IsPined != 0 {
			selects = append(selects, "IsPined")
		}
		if body.IsRecommend != 0 {
			selects = append(selects, "IsRecommend")
		}
		if body.IsOriginal != 0 {
			selects = append(selects, "IsOriginal")
		}
		if er := tx.Select(selects).Updates(body).Error; er != nil {
			err = er
		}
	}

	if body.Type == "video" && body.IsOriginal != 2 {
		if er2 := db.Orm.Model(&model.VideoMeta{}).Where("pid = ?", body.ID).Updates(body.Meta).Error; er2 != nil {
			err = er2
		}
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

func UpdatePostPv(c *gin.Context) {
	pr := &model.PostRanking{}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, _ := service.GetSimplePostByID(uint(id))
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
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if err := db.Orm.Model(&model.Post{}).
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Meta").
		Select("posts.*, SUM(post_rankings.hits) as Hits").
		Joins("left join post_rankings on post_rankings.pid = posts.id and to_days(post_rankings.hits_at) = to_days(now())").
		Limit(size).
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
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "30"))
	if err := db.Orm.Model(&model.Post{}).
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Meta").
		Select("posts.*, SUM(post_rankings.hits) as Hits").
		Joins("left join post_rankings on post_rankings.pid = posts.id and DATE_FORMAT(post_rankings.hits_at,'%Y%m') = DATE_FORMAT(CURDATE(),'%Y%m')").
		Limit(size).
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
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))
	if err := db.Orm.Model(&model.Post{}).
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Meta").
		Select("posts.*, SUM(post_rankings.hits) as Hits").
		Joins("left join post_rankings on post_rankings.pid = posts.id").
		Limit(size).
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

func GetVideosUpdateOnWeek(c *gin.Context) {
	posts := []*model.Post{}
	if err := db.Orm.Model(&model.Post{}).
		Preload("Meta").
		Joins("left join video_metas on video_metas.pid = posts.id").
		Where("posts.type = 'video' AND posts.is_original <> 2 AND video_metas.is_end <> 2").
		Order("video_metas.updated_date desc").
		Find(&posts).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err,
		})
		return
	}
	var week []interface{}
	for i := 0; i < 7; i++ {
		week = append(week, []*model.Post{})
	}
	for _, post := range posts {
		if post.Meta.UpdatedDate != nil {
			fmt.Print(post.Meta.UpdatedDate.Weekday())
			week[post.Meta.UpdatedDate.Weekday()] = append(week[post.Meta.UpdatedDate.Weekday()].([]*model.Post), post)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": week,
	})
}

func ReviewPost(c *gin.Context) {
	id := c.Param("id")
	level := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims).Level

	if !common.IsAdmin(level) && !common.IsRoot(level) {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "Forbidden",
		})
		return
	}

	status, _ := strconv.Atoi(c.Query("status"))
	IsRecommend, _ := strconv.Atoi(c.Query("is_recommend"))

	fmt.Print(status, IsRecommend, "gfhjkl;")
	tx := db.Orm.Model(&model.Post{}).
		Where("id = ?", id).
		Updates(&model.Post{
			Status:      uint(status),
			IsRecommend: uint(IsRecommend),
		})

	c.JSON(http.StatusOK, gin.H{
		"err": tx.Error,
	})
}

////////////////////////////////////////////////////////////////////////////////

func isYourPost(c *gin.Context, body *model.Post) bool {
	post, err := service.GetSimplePostByID(body.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": err})
		return false
	}
	ctxUser := c.MustGet(middleware.CTX_AUTH_KEY).(*middleware.AuthClaims)
	if !(common.IsRoot(ctxUser.Level) || common.IsAdmin(ctxUser.Level)) && (post.Uid != ctxUser.UserId || body.Status != post.Status) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"err": "Forbidden."})
		return false
	}
	return true
}

func checkPostField(c *gin.Context, body *model.Post, user *middleware.AuthClaims) bool {
	if body.Type == "post" && len(body.Content) < 15 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "Content should not be empty or less than 15 characters.",
		})
		return false
	}

	if common.IsUser(user.Level) && body.Type == "video" {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "You have no permission.",
		})
		return false
	}

	if !common.IsRoot(user.Level) || !common.IsAdmin(user.Level) {
		body.Status = 0
		body.IsPined = 0
		body.IsRecommend = 0
	}

	if body.Type == "post" {
		body.Status = 4
	}

	tags, _ := getDBTags()
	body.Tags = common.FixTags(tags, body.Tags)
	return true
}
