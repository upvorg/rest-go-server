package model

import (
	"upv.life/server/db"
)

type User struct {
	ID          int64  `json:"id,omitempty"`
	Level       int8   `json:"level,omitempty"`  // sql default
	Status      int8   `json:"status,omitempty"` // sql default
	Name        string `json:"name"`
	Pwd         string `json:"pwd,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	Avatar      string `json:"avatar,omitempty"` // sql default
	Bio         string `json:"bio,omitempty"`    // sql default
	Qq          string `json:"qq,omitempty"`
	Create_Time string `json:"create_time,omitempty"` // sql default
	Update_Time string `json:"update_time,omitempty"` // sql default
}

func (this *User) Create() (*User, error) {
	if result, err := db.Sqlx.NamedExec("INSERT into users (name, nickname, pwd) values (:name, :nickname, :pwd)", this); err != nil {
		return nil, err
	} else {
		this.ID, _ = result.LastInsertId()
		return this, nil
	}
}

func (this *User) GetUserByID() error {
	return db.Sqlx.Get(this, "SELECT id, level, status, name, nickname, avatar, bio, qq, create_time, update_time from users where id = ?", this.ID)
}

func (this *User) GetUserByName() error {
	return db.Sqlx.Get(this, "SELECT id, level, status, name, nickname, avatar, bio, qq, create_time, update_time from users where name = ?", this.Name)
}

func (this *User) GetUsersByNickname() error {
	return db.Sqlx.Select(this, "SELECT id, level, status, name, nickname, avatar, bio, qq, create_time, update_time from users where nickname like %?%", this.Nickname)
}

func (this *User) Update() error {
	if _, err := db.Sqlx.NamedExec("UPDATE user set level = :level, status = :status, name = :name, nickname = :nickname, avatar = :avatar, bio = :bio, qq = :qq from users where id = :id", this); err != nil {
		return err
	} else {
		return nil
	}
}

func (this *User) Delete() error {
	if _, err := db.Sqlx.Exec("DELETE FROM users where id = ?", this.ID); err != nil {
		return err
	} else {
		return nil
	}
}
