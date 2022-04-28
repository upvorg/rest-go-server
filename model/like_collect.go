package model

import "time"

type Like struct {
	ID        uint `gorm:"primaryKey"`
	Uid       uint
	Pid       uint
	CreatedAt *time.Time
}

type Collect struct {
	ID        uint `gorm:"primaryKey"`
	Uid       uint
	Pid       uint
	CreatedAt *time.Time
}
