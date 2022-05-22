package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	POST_STATUS_DELETED   uint = 1
	POST_STATUS_REJECT    uint = 2
	POST_STATUS_PENDING   uint = 3
	POST_STATUS_PUBLISHED uint = 4
	POST_STATUS_DRAFT     uint = 5
)

//[]byte
type Post struct {
	ID          string         `gorm:"primary_key;unique;type:varchar(36);not null"`
	Cover       string         `gorm:"size:200"`
	Title       string         `gorm:"size:60;not null" binding:"required"`
	Content     string         `gorm:"type:text"`
	Uid         uint           `json:"-"`
	Tags        string         `gorm:"size:100;column:tags"`
	Status      uint           `gorm:"default:3"`
	Type        string         `gorm:"default:post" binding:"required"`
	IsPined     uint           `gorm:"default:1"`
	IsRecommend uint           `gorm:"default:1"`
	IsOriginal  uint           `gorm:"default:1"`
	CreatedAt   *time.Time     `gorm:"type:timestamp"`
	UpdatedAt   *time.Time     `gorm:"type:timestamp"`
	DeletedAt   gorm.DeletedAt `json:"-"`
	Meta        *VideoMeta     `gorm:"foreignKey:Pid" form:"meta,omitempty" json:"Meta,omitempty"`
	Creator     *User          `gorm:"foreignKey:Uid" form:"user,omitempty" json:"Creator,omitempty"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	post.ID = uuid.NewString()
	return nil
}

type PostInFeed struct {
	Post       `gorm:"embedded"`
	IsLiked    uint `json:"IsLiked"`
	LikesCount uint `json:"LikesCount"`
}

type FullPost struct {
	Post            `gorm:"embedded"`
	Hits            uint `json:"Hits"`
	IsLiked         uint `json:"IsLiked"`
	LikesCount      uint `json:"LikesCount"`
	IsCollected     uint `json:"IsCollected"`
	CollectionCount uint `json:"CollectionCount"`
	CommentCount    uint `json:"CommentCount"`
}

func PreloadCreatorOptinal() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Creator", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, nickname, avatar, bio")
		})
	}
}

//api

type Meta struct {
	KeyWord    string `form:"keyword,omitempty"`
	Type       string `form:"type,omitempty,default=video"`
	Region     string `form:"region,omitempty"`
	IsEnd      int    `form:"isend,omitempty"`
	Genre      string `form:"genre,omitempty"`
	Tag        string `form:"tag,omitempty"`
	IsOriginal uint   `form:"is_original,omitempty"`
	Uid        string `form:"uid,omitempty"`
	Status     string `form:"status,omitempty"`
}

type PostRanking struct {
	ID     uint       `gorm:"primaryKey"`
	Pid    string     `gorm:"not null;"`
	Hits   uint       `gorm:"not null;"`
	HitsAt *time.Time `gorm:"type:timestamp;not null;default:now()"`
}

type PostActicity struct {
	Type         string // like, comment, collection
	PostType     string // post, video
	Pid          string
	Uid          uint
	PostTitle    string
	UserName     string
	UserNickname string
	UserAvatar   string
	Comment      string `json:"Comment,omitempty"`
	CreatedAt    *time.Time
}
