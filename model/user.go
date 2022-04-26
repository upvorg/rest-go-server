package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Level     uint
	Status    uint
	Name      string `gorm:"size:16;unique_index"`
	Pwd       string `json:"-"`
	Nickname  string `gorm:"size:16"`
	Avatar    string
	Bio       string
	QQ        string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	// Posts     []Post `gorm:"foreignKey:Uid"`
	// gorm.Model
}
