package model

import (
	"time"

	"gorm.io/gorm"
)

// TODO: add comment @reply #i
type Comment struct {
	ID        uint `gorm:"primaryKey"`
	ParentID  uint
	TargetID  uint
	Uid       uint   `gorm:"not null"`
	Pid       uint   `gorm:"not null"`
	Vid       uint   `gorm:"default:1"`
	Content   string `gorm:"size:200"`
	Color     string `gorm:"size:10"`
	CreatedAt *time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
	Creator   *User          `gorm:"foreignKey:Uid" json:"Creator,omitempty"`
	Children  []*Comment     `gorm:"foreignKey:parent_id" json:"Children,omitempty"`
}
