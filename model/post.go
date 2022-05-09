package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID              uint           `gorm:"primaryKey"`
	Cover           string         `gorm:"size:200"`
	Title           string         `gorm:"size:60;not null" binding:"required"`
	Content         string         `gorm:"type:text"`
	Uid             uint           `json:"-"`
	Tags            string         `gorm:"size:100;column:tags"`
	Status          uint8          `gorm:"default:3"`
	Type            string         `gorm:"default:post" binding:"required"`
	IsPined         uint           `gorm:"default:1"`
	IsRecommend     uint           `gorm:"default:1"`
	IsOriginal      uint           `gorm:"default:1"`
	CreatedAt       *time.Time     `gorm:"type:timestamp"`
	UpdatedAt       *time.Time     `gorm:"type:timestamp"`
	DeletedAt       gorm.DeletedAt `json:"-"`
	Meta            *VideoMeta     `gorm:"foreignKey:Pid" form:"meta,omitempty" json:"Meta,omitempty"`
	Creator         *User          `gorm:"foreignKey:Uid" form:"user,omitempty" json:"Creator,omitempty"`
	Hits            uint           `gorm:"<-:false" json:"Hits,omitempty"`
	LikesCount      uint           `gorm:"<-:false" json:"LikesCount,omitempty"`
	CommentCount    uint           `gorm:"<-:false" json:"CommentCount,omitempty"`
	CollectionCount uint           `gorm:"<-:false" json:"CollectionCount,omitempty"`
	IsLiked         uint           `gorm:"<-:false" json:"IsLiked,omitempty"`
	IsCollected     uint           `gorm:"<-:false" json:"IsCollected,omitempty"`
}

type PostRanking struct {
	ID     uint       `gorm:"primaryKey"`
	Pid    uint       `gorm:"not null;"`
	Hits   uint       `gorm:"not null;"`
	HitsAt *time.Time `gorm:"type:timestamp;not null;default:now()"`
}

type VideoMeta struct {
	ID            uint           `gorm:"primaryKey" json:"-"`
	Pid           uint           `gorm:"not null" json:"-"`
	TitleJapanese string         `gorm:"size:60"`
	TitleRomanji  string         `gorm:"size:60"`
	Genre         string         `gorm:"not null;size:10;default:新番" form:"genre,default=新番"`
	Region        string         `gorm:"not null;size:10;default:其他" form:"region,default=其他"`
	Episodes      int            `gorm:"default:0"`
	IsEnd         uint8          `gorm:"default:1" form:"is_end,default=1"`
	PublishDate   *time.Time     `gorm:"type:timestamp"`
	UpdatedDate   *time.Time     `gorm:"type:timestamp"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}

func (vm *VideoMeta) TableName() string {
	return "video_metas"
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
}
