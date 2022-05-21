package model

import (
	"time"

	"github.com/google/uuid"
)

type Video struct {
	ID            uint   `gorm:"primaryKey"`
	Episode       int    `gorm:"default:1"`
	Cover         string `gorm:"size:200"`
	Title         string `gorm:"size:60;"`
	TitleJapanese string `gorm:"size:60"`
	TitleRomanji  string `gorm:"size:60"`
	VideoUrl      string `gorm:"size:1024"`
	Synopsis      string `gorm:"size:200"`
	Uid           uint
	Pid           uuid.UUID
	CreatedAt     *time.Time `gorm:"type:timestamp"`
	UpdatedAt     *time.Time `gorm:"type:timestamp"`
}
