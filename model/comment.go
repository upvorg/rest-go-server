package model

import (
	"time"

	"gorm.io/gorm"
)

// TODO: add comment @reply #i
type Comment struct {
	ID        uint `gorm:"primaryKey"`
	Uid       uint
	Pid       uint
	Vid       uint   `gorm:"default:1"`
	Content   string `gorm:"size:200"`
	Color     string `gorm:"size:10"`
	CreatedAt *time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
