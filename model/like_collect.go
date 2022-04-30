package model

import "time"

type Like struct {
	ID        uint `gorm:"primaryKey"`
	Uid       uint
	Pid       uint
	CreatedAt *time.Time
}

type Collection struct {
	ID        uint `gorm:"primaryKey"`
	Uid       uint
	Pid       uint
	CreatedAt *time.Time
}
