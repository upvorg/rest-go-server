package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VideoMeta struct {
	ID            uint           `gorm:"primaryKey" json:"-"`
	Pid           uuid.UUID      `gorm:"not null" json:"-"`
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
