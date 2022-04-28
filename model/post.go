package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID            uint   `gorm:"primaryKey"`
	Cover         string `gorm:"size:200"`
	Title         string `gorm:"size:60;not null" binding:"required"`
	Content       string `gorm:"type:text"`
	Uid           uint   `json:"-"`
	Tag           string `gorm:"size:100"`
	Status        uint8  `gorm:"default:4"`
	Type          string `gorm:"default:post" binding:"required"`
	Pv            uint
	IsPined       uint `gorm:"default:1"`
	IsRecommend   uint `gorm:"default:1"`
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     gorm.DeletedAt `json:"-"`
	Meta          *VideoMeta     `gorm:"foreignKey:Pid" form:"meta,omitempty" json:"Meta,omitempty"`
	Creator       *User          `gorm:"foreignKey:Uid" form:"user,omitempty" json:"Creator,omitempty"`
	LikesCount    int            `gorm:"-"`
	CommentsCount int            `gorm:"-"`
	CollectsCount int            `gorm:"-"`
	IsLiked       bool           `gorm:"-"`
	IsCollected   bool           `gorm:"-"`
}

type VideoMeta struct {
	ID            uint   `gorm:"primaryKey" json:"-"`
	Pid           uint   `gorm:"not null" json:"-"`
	TitleJapanese string `gorm:"size:60"`
	TitleRomanji  string `gorm:"size:60"`
	Genre         string `gorm:"not null;size:10;default:新番" form:"genre,default=新番"`
	Region        string `gorm:"not null;size:10;default:其他" form:"region,default=其他"`
	Synopsis      string `gorm:"size:200"`
	IsEnd         uint8  `gorm:"default:1" form:"is_end,default=1"`
	PublishDate   *time.Time
	UpdatedDate   *time.Time
}

func (vm *VideoMeta) TableName() string {
	return "video_metas"
}

//api

type Meta struct {
	KeyWord string `form:"keyword,omitempty"`
	Type    string `form:"type,omitempty,default=post"`
	Region  string `form:"region,omitempty"`
	IsEnd   int    `form:"isend,omitempty"`
	Genre   string `form:"genre,omitempty"`
}
