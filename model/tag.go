package model

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        uint           `gorm:"primaryKey"`
	Uid       uint           `json:"-"`
	Name      string         `gorm:"size:20"`
	Synopsis  string         `gorm:"size:200"`
	CreatedAt *time.Time     `gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
