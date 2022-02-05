package model

import (
	sq "github.com/Masterminds/squirrel"
	"upv.life/server/db"
)

type User struct {
	ID          uint64 `json:"id,omitempty"`
	Level       uint8  `json:"level,omitempty"`  // sql default
	Status      uint8  `json:"status,omitempty"` // sql default
	Name        string `json:"name"`
	Pwd         string `json:"pwd,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	Avatar      string `json:"avatar,omitempty"` // sql default
	Bio         string `json:"bio,omitempty"`    // sql default
	Qq          string `json:"qq,omitempty"`
	Create_Time string `json:"create_time,omitempty"` // sql default
	Update_Time string `json:"update_time,omitempty"` // sql default
	UserMeta
}

type UserMeta struct {
	//获赞数目
	Like_Count uint64 `json:"like_count"`
	//动态数目
	Feed_Count uint64 `json:"feed_count"`
	//视频数目
	Video_Count uint64 `json:"video_count"`
}

func (this *User) Create() (*User, error) {
	if result, err := db.Sqlx.NamedExec("INSERT into users (name, nickname, pwd) values (:name, :nickname, :pwd)", this); err != nil {
		return nil, err
	} else {
		id, _ := result.LastInsertId()
		this.ID = uint64(id)
		return this, nil
	}
}

func (this *User) GetUserByID() error {
	sql, arg, _ := sq.Select(`
		id, level, status, name, nickname, avatar, bio, qq, create_time, update_time,
		(SELECT COUNT(*) FROM likes l WHERE l.uid = users.id) AS like_count,
		(select COUNT(*) from posts p WHERE p.uid = users.id and p.type='post') AS feed_count,
		(select COUNT(*) from posts p WHERE p.uid = users.id and p.type='video') AS video_count
	`).
		From("users").
		Where(sq.Eq{"id": this.ID}).
		ToSql()
	return db.Sqlx.Get(this, sql, arg...)
}

func (this *User) GetUserByName() error {
	return db.Sqlx.Get(this, "SELECT id, level, status, name, nickname, avatar, bio, qq, create_time, update_time from users where name = ?", this.Name)
}

func (this *User) GetUserByNameWithPwd() error {
	return db.Sqlx.Get(this, "SELECT id, level, status, name, pwd, nickname, avatar, bio, qq, create_time, update_time from users where name = ?", this.Name)
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
