package model

import (
	"log"
	"reflect"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"upv.life/server/db"
)

type User struct {
	ID          uint64 `json:"id,omitempty"`
	Level       *uint8 `json:"level,omitempty"`  // sql default 包含 0 故 *
	Status      *uint8 `json:"status,omitempty"` // sql default
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

func (User) TableName() string {
	return "users"
}

func (this *User) Create() (*User, error) {
	if result, err := db.DB.NamedExec("INSERT into users (name, nickname, pwd) values (:name, :nickname, :pwd)", this); err != nil {
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
	return db.DB.Get(this, sql, arg...)
}

func (this *User) GetUserByName() error {
	return db.DB.Get(this, "SELECT id, level, status, name, nickname, avatar, bio, qq, create_time, update_time from users where name = ?", this.Name)
}

func (this *User) GetUserByNameWithPwd() error {
	return db.DB.Get(this, "SELECT id, level, status, name, pwd, nickname, avatar, bio, qq, create_time, update_time from users where name = ?", this.Name)
}

func (this *User) GetUsersByNickname() error {
	return db.DB.Select(this, "SELECT id, level, status, name, nickname, avatar, bio, qq, create_time, update_time from users where nickname like %?%", this.Nickname)
}

func (this *User) Update() error {
	if _, err := db.DB.NamedExec("UPDATE user set level = :level, status = :status, name = :name, nickname = :nickname, avatar = :avatar, bio = :bio, qq = :qq from users where id = :id", this); err != nil {
		return err
	} else {
		return nil
	}
}

func (this *User) Delete() error {
	if _, err := db.DB.Exec("DELETE FROM users where id = ?", this.ID); err != nil {
		return err
	} else {
		return nil
	}
}

// get list with pagination
func (this *User) GetAll(page, pageSize uint64) ([]User, uint64, error) {
	var users []User
	var count uint64
	v := reflect.ValueOf(this)
	t := reflect.TypeOf(this)
	countSql := sq.Select("COUNT(id)")
	sql := sq.Select("id, level, status, name, nickname, avatar, bio, qq, create_time, update_time")

	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		log.Print(t.Field(i).Name, ":", v.Field(i))
		switch v.Field(i).Kind() {
		case reflect.Int:
		case reflect.Uint:
		case reflect.Uint8:
		case reflect.Uint64:
			// cv := v.Field(i).Convert(reflect.TypeOf(uint64(0)))
			// if cv.Uint() != 0 {
			sql = sql.Where(squirrel.Eq{t.Field(i).Name: v.Field(i).Uint()})
			countSql = countSql.Where(squirrel.Eq{t.Field(i).Name: v.Field(i).Uint()})
			// }
		case reflect.String:
			if v.Field(i).String() != "" {
				if t.Field(i).Name == "Name" {
					like := "%" + v.Field(i).String() + "%"
					sql = sql.Where(squirrel.Or{
						squirrel.Like{"name": like},
						squirrel.Like{"nickname": like},
					})
					countSql = countSql.Where(squirrel.Or{
						squirrel.Like{"name": like},
						squirrel.Like{"nickname": like},
					})
				} else {
					sql = sql.Where(squirrel.Eq{t.Field(i).Name: v.Field(i).String()})
					countSql = countSql.Where(squirrel.Eq{t.Field(i).Name: v.Field(i).String()})
				}
			}
		}
	}
	sql = sql.
		From("users").
		Offset(uint64(pageSize * (page - 1))).
		Limit(uint64(pageSize))
	countSql = countSql.From("users")

	query, arg, _ := sql.ToSql()
	cq, ca, _ := countSql.ToSql()
	err := db.DB.Select(&users, query, arg...)
	err = db.DB.QueryRowx(cq, ca...).Scan(&count)

	return users, count, err
}

func (this *User) Test() {
	log.Println(this.TableName())
	this.Test()
}
