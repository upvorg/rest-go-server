package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"upv.life/server/common"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
)

// tips: left join 1:n  repeat data
func GetSimplePostByID(id uint) (*model.Post, error) {
	var post model.Post
	if err := db.Orm.Model(&model.Post{}).
		Where("posts.id = ?", id).
		First(&post).Error; err == gorm.ErrRecordNotFound {
		return nil, errors.New("post not found")
	} else if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetPostById(id uint, uid uint, level uint) (*model.Post, error) {
	var post model.Post
	withUserQuery := ""
	tx := db.Orm.Model(&model.Post{})

	if uid != 0 {
		withUserQuery = fmt.Sprintf(`
		,IF((SELECT id FROM likes WHERE likes.pid = posts.id AND likes.uid = %d),2,1) as IsLiked,
		IF((SELECT id FROM collections WHERE collections.pid = posts.id AND collections.uid = %d),2,1) as IsCollected
		`, uid, uid)
	} else {
		withUserQuery = `,1 as IsLiked,1 as IsCollected`
	}

	err := tx.
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Meta").
		Select(`
		posts.*,
		(SELECT COUNT(id) FROM likes WHERE likes.pid = posts.id) as LikesCount,
		(SELECT COUNT(id) FROM collections WHERE collections.pid = posts.id) as CollectionCount,
		(SELECT COUNT(id) FROM comments WHERE comments.pid = posts.id) as CommentCount,
		(SELECT SUM(hits) FROM post_rankings WHERE pid = posts.id) as Hits
		`+withUserQuery).
		Where("posts.id = ?", id).
		Group("posts.id").
		First(&post).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("post not found")
	} else if err != nil {
		return nil, err
	}

	if post.Status != model.POST_STATUS_PUBLISHED {
		if uid == 0 {
			return nil, errors.New("post not found")
		} else {
			if post.Creator.ID != uid {
				return &post, nil
			}
			if !common.IsAdmin(level) && !common.IsRoot(level) {
				return nil, errors.New("post not found")
			}
		}
	}

	return &post, nil
}

//TODO: 前台不返回不为 POST_STATUS_PUBLISHED 的文章
func GetPostsByMetaType(m model.Meta, c *gin.Context) (*[]model.Post, error) {
	var posts []model.Post
	withUserQuery := ""
	tx := db.Orm.Model(&model.Post{}).Scopes(model.Paginate(c))
	user, exists := c.Get(middleware.CTX_AUTH_KEY)

	if exists {
		withUserQuery = fmt.Sprintf(`
		,IF((SELECT id FROM likes WHERE likes.pid = posts.id AND likes.uid = %d),2,1) as IsLiked,
		IF((SELECT id FROM collections WHERE collections.pid = posts.id AND collections.uid = %d),2,1) as IsCollected
		`, user.(*middleware.AuthClaims).UserId, user.(*middleware.AuthClaims).UserId)
	} else {
		withUserQuery = `,1 as IsLiked,1 as IsCollected`
	}

	tx.
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Meta").
		Select(`
		posts.*,
		(SELECT COUNT(id) FROM likes WHERE likes.pid = posts.id) as LikesCount,
		(SELECT COUNT(id) FROM collections WHERE collections.pid = posts.id) as CollectionCount,
		(SELECT COUNT(id) FROM comments WHERE comments.pid = posts.id) as CommentCount,
		(SELECT SUM(hits) FROM post_rankings WHERE pid = posts.id) as Hits
		`+withUserQuery).
		Joins("left join video_metas on video_metas.pid = posts.id").
		Where("posts.type = ?", m.Type).
		Order("posts.created_at desc").
		Group("posts.id")

	if m.Uid != "" {
		tx.Where("posts.uid = ?", m.Uid)
	}

	if m.Status != "" {
		tx.Where("posts.status = ?", m.Status)
	}

	if m.Tag != "" {
		tx.Where("posts.tags LIKE ?", "%"+m.Tag+"%")
	}

	if m.KeyWord != "" {
		tx.Where("posts.title LIKE ? OR posts.content LIKE ?", "%"+m.KeyWord+"%", "%"+m.KeyWord+"%")
	}

	if m.IsOriginal != 0 {
		tx.Where("posts.is_original = ?", m.IsOriginal)
	}

	if m.Genre != "" {
		tx.Where("video_metas.genre = ?", m.Genre)
	}

	if m.Region != "" {
		tx.Where("video_metas.region = ?", m.Region)
	}

	if m.IsEnd != 0 {
		tx.Where("video_metas.is_end = ?", m.IsEnd)
	}

	if err := tx.Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func GetPostsByTag(tag string, c *gin.Context) (*[]model.Post, error) {
	var posts []model.Post

	if err := db.Orm.Scopes(model.Paginate(c)).
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Meta").
		Where("tags LIKE ?", "%"+tag+"%").
		Order("posts.created_at desc").
		Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func GetRecommendPosts(size int) (*[]model.Post, error) {
	var posts []model.Post
	if err := db.Orm.Where("is_recommend = ?", 2).
		Scopes(model.PreloadCreatorOptinal()).
		Preload("Meta").
		Order("posts.created_at desc").
		Limit(size).
		Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}
