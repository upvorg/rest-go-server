package model

import "time"

type Post struct {
	ID          uint   `gorm:"primaryKey"`
	Cover       string `gorm:"size:200"`
	Title       string `gorm:"size:60;" binding:"required"`
	Content     string `gorm:"type:text"`
	Uid         uint   `json:"-"`
	Tag         string `gorm:"size:100"`
	Status      uint8  `gorm:"default:4"`
	Type        string `gorm:"default:post" form:"type,default=post"`
	Pv          uint
	IsPined     uint `gorm:"default:1"`
	IsRecommend uint `gorm:"default:1"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	Meta        *VideoMetas `json:"Meta,omitempty" gorm:"foreignKey:Pid"` // type video
	Creator     User        `gorm:"foreignKey:Uid;"`                      // type user
	// LikesCount    int
	// CommentsCount int
	// CollectsCount int
}

type VideoMetas struct {
	ID          uint `gorm:"primaryKey" json:"-"`
	Pid         uint
	Type        string     `gorm:"default:新番"`
	Region      string     `gorm:"size:8"`
	IsEnd       uint8      `gorm:"default:1"`
	PublishDate *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

type Likes struct {
	ID  uint `gorm:"primaryKey"`
	Uid uint
	Pid uint
}

//api

type Meta struct {
	KeyWord  string `form:"keyword,omitempty"`
	Type     string `form:"type,omitempty"`
	Region   string `form:"region,omitempty"`
	IsEnd    int    `form:"isend,omitempty"`
	PostType string `form:"pt,omitempty,default=video"`
	// Limit uint64 `json:"limit" form:"limit,default=15" binding:"min=1,max=30"`

}
