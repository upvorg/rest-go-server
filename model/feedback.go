package model

import "time"

type Feedback struct {
	Id          int64      `gorm:"primary_key"`
	Ip          string     `gorm:"size:20"`
	Name        string     `gorm:"type:varchar(15);default:'佚名'"`
	DisplayName string     `gorm:"type:varchar(15);default:'佚名'"`
	Email       string     `gorm:"type:varchar(50);default:''"`
	Message     string     `gorm:"type:text;not null" binding:"required"`
	CreatedAt   *time.Time `gorm:"type:timestamp;"`
}
