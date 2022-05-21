package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:16;unique_index"`
	Pwd       string         `json:"-"`
	Nickname  string         `gorm:"size:16"`
	Avatar    string         `gorm:"default:'https://q1.qlogo.cn/g?b=qq&nk=7619376472&s=640'"`
	Bio       string         `gorm:"default:这个人很懒，什么都没有留下"`
	Email     string         `gorm:"size:50;unique_index;default=null" binding:"omitempty,email" json:"Email,omitempty"`
	Level     uint           `gorm:"default:4" json:"Level,omitempty"`
	Status    uint           `gorm:"default:2" json:"Status,omitempty"`
	CreatedAt *time.Time     `gorm:"type:timestamp" json:"CreatedAt,omitempty"`
	UpdatedAt *time.Time     `gorm:"type:timestamp" json:"UpdatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type UserStat struct {
	LikesCount      uint `gorm:"<-:false"`
	CommentCount    uint `gorm:"<-:false"`
	CollectionCount uint `gorm:"<-:false"`
	Pits            uint `gorm:"<-:false"`
	Vits            uint `gorm:"<-:false"`
}
