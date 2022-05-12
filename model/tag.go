package model

type Tag struct {
	ID       uint   `gorm:"primaryKey"`
	Uid      uint   `json:"-"`
	Name     string `gorm:"size:20"`
	Synopsis string `gorm:"size:200"`
}
