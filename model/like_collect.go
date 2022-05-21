package model

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
	ID        uint `gorm:"primaryKey"`
	Uid       uint
	Pid       uuid.UUID
	CreatedAt *time.Time `gorm:"type:timestamp"`
}

type Collection struct {
	ID        uint `gorm:"primaryKey"`
	Uid       uint
	Pid       uuid.UUID
	CreatedAt *time.Time `gorm:"type:timestamp"`
}
