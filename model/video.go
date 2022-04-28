package model

import "time"

type Video struct {
	ID            uint   `gorm:"primaryKey"`
	Episode       int    `gorm:"default:1"`
	Cover         string `gorm:"size:200"`
	Title         string `gorm:"size:60;"`
	TitleJapanese string `gorm:"size:60"`
	TitleRomanji  string `gorm:"size:60"`
	Content       string `gorm:"size:200"`
	Uid           uint
	Pid           uint
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
