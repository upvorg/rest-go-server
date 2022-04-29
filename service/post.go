package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"upv.life/server/db"
	"upv.life/server/model"
)

//TODO: get like&collect&comment count & is_like & is_collect
func GetPostById(id uint) (*model.Post, error) {
	var post model.Post
	if err := db.Orm.Model(&model.Post{}).
		Preload("Creator").
		Preload("Meta").
		Where("posts.id = ?", id).
		First(&post).Error; err == gorm.ErrRecordNotFound {
		return nil, nil
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
		Joins("left join video_metas on video_metas.pid = posts.id").
		Where("posts.type = ?", m.Type)

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
		Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func GetRecommendPosts() (*[]model.Post, error) {
	var posts []model.Post
	if err := db.Orm.Where("is_recommend = ?", 2).Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}
