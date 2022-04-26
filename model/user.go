package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Level     uint
	Status    uint
	Name      string `gorm:"size:16;unique_index"`
	Pwd       string
	Nickname  string `gorm:"size:16"`
	Avatar    string
	Bio       string
	Qq        string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	gorm.Model
}
