package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"upv.life/server/db"
	"upv.life/server/model"
)

//TODO: get is_like & is_collect where likes.uid = uid
func GetPostById(id uint) (*model.Post, error) {
	var post model.Post
	if err := db.Orm.Model(&model.Post{}).
		Preload("Creator").
		Preload("Meta").
		Select(`
		posts.*,
		COUNT(likes.id) as LikesCount,
		COUNT(collects.id) as CollectsCount,
		COUNT(comments.id) as CommentsCount
		`).
		Joins("left join likes on likes.pid = posts.id").
		Joins("left join collects on collects.pid = posts.id").
		Joins("left join comments on comments.pid = posts.id").
		Where("posts.id = ?", id).
		Group("posts.id").
		First(&post).Error; err == gorm.ErrRecordNotFound {
		return nil, errors.New("post not found")
	} else if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetPostsByMetaType(m model.Meta, c *gin.Context) (*[]model.Post, error) {
	var posts []model.Post
	tx := db.Orm.Model(&model.Post{}).Scopes(model.Paginate(c)).
		Preload("Creator").
		Preload("Meta").
		Select(`
		posts.*,
		COUNT(likes.id) as LikesCount,
		COUNT(collects.id) as CollectsCount,
		COUNT(comments.id) as CommentsCount
		`).
		Joins("left join likes on likes.pid = posts.id").
		Joins("left join collects on collects.pid = posts.id").
		Joins("left join comments on comments.pid = posts.id").
		Joins("left join video_metas on video_metas.pid = posts.id").
		Where("posts.type = ?", m.Type).
		Order("posts.created_at desc").
		Group("posts.id")

	if m.KeyWord != "" {
		tx.Where("posts.title LIKE ? OR posts.content LIKE ?", "%"+m.KeyWord+"%", "%"+m.KeyWord+"%")
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
		Where("tag LIKE ?", "%"+tag+"%").
		Order("posts.created_at desc").
		Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func GetRecommendPosts() (*[]model.Post, error) {
	var posts []model.Post
	if err := db.Orm.Where("is_recommend = ?", 2).Order("posts.created_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}
