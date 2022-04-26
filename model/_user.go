// package model

// import (
// 	"log"
// 	"reflect"

// 	sq "github.com/Masterminds/squirrel"
// 	"upv.life/server/db"
// )

// type User struct {
// 	ID          uint64 `json:"id,omitempty"`
// 	Level       *uint8 `json:"level,omitempty"`  // sql default 包含 0 故 *
// 	Status      *uint8 `json:"status,omitempty"` // sql default
// 	Name        string `json:"name"`
// 	Pwd         string `json:"pwd,omitempty"`
// 	Nickname    string `json:"nickname,omitempty"`
// 	Avatar      string `json:"avatar,omitempty"` // sql default
// 	Bio         string `json:"bio,omitempty"`    // sql default
// 	Qq          string `json:"qq,omitempty"`
// 	Create_Time string `json:"create_time,omitempty"` // sql default
// 	Update_Time string `json:"update_time,omitempty"` // sql default
// 	UserMeta
// }

// type UserMeta struct {
// 	//获赞数目
// 	Like_Count uint64 `json:"like_count"`
// 	//动态数目
// 	Feed_Count uint64 `json:"feed_count"`
// 	//视频数目
// 	Video_Count uint64 `json:"video_count"`
// }

// func (User) TableName() string {
// 	return "users"
// }

// func (user *User) Create() (*User, error) {
// 	if result, err := db.DB.NamedExec("INSERT into users (name, nickname, pwd) values (:name, :nickname, :pwd)", user); err != nil {
// 		return user, err
// 	} else {
// 		id, _ := result.LastInsertId()
// 		user.ID = uint64(id)
// 		return user, nil
// 	}
// }

// func (user *User) GetUserByID() error {
// 	sql, arg, _ := sq.Select(`
// 		id, level, status, name, nickname, avatar, bio, qq, create_time, update_time,
// 		(SELECT COUNT(*) FROM likes l WHERE l.uid = users.id) AS like_count,
// 		(select COUNT(*) from posts p WHERE p.uid = users.id and p.type='post') AS feed_count,
// 		(select COUNT(*) from posts p WHERE p.uid = users.id and p.type='video') AS video_count
// 	`).
// 		From("users").
// 		Where(sq.Eq{"id": user.ID}).
// 		ToSql()
// 	return db.DB.Get(user, sql, arg...)
// }

// func (user *User) GetUserByName() error {
// 	return db.DB.Get(user, "SELECT id, level, status, name, nickname, avatar, bio, qq, create_time, update_time from users where name = ?", user.Name)
// }

// func (user *User) GetUserByNameWithPwd() error {
// 	return db.DB.Get(user, "SELECT id, level, status, name, pwd, nickname, avatar, bio, qq, create_time, update_time from users where name = ?", user.Name)
// }

// func (user *User) GetUsersByNickname() error {
// 	return db.DB.Select(user, "SELECT id, level, status, name, nickname, avatar, bio, qq, create_time, update_time from users where nickname like %?%", user.Nickname)
// }

// func (user *User) Update() error {
// 	if _, err := db.DB.NamedExec("UPDATE user set level = :level, status = :status, name = :name, nickname = :nickname, avatar = :avatar, bio = :bio, qq = :qq from users where id = :id", user); err != nil {
// 		return err
// 	} else {
// 		return nil
// 	}
// }

// func (user *User) Delete() error {
// 	if _, err := db.DB.Exec("DELETE FROM users where id = ?", user.ID); err != nil {
// 		return err
// 	} else {
// 		return nil
// 	}
// }

// // get list with pagination
// func (user *User) GetAll(page, pageSize uint64) ([]User, uint64, error) {
// 	var users []User
// 	var count uint64
// 	v := reflect.ValueOf(user)
// 	t := reflect.TypeOf(user)
// 	sqlBuilder := sq.StatementBuilder

// 	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
// 		v = v.Elem()
// 		t = t.Elem()
// 	}

// 	for i := 0; i < v.NumField(); i++ {
// 		switch v.Field(i).Kind() {
// 		case reflect.Uint:
// 		case reflect.Uint8:
// 		case reflect.Uint64:
// 			cv := v.Field(i).Convert(reflect.TypeOf(uint64(0)))
// 			if cv.Uint() != 0 {
// 				sqlBuilder = sqlBuilder.Where(sq.Eq{t.Field(i).Name: v.Field(i).Uint()})
// 			}
// 		case reflect.String:
// 			if v.Field(i).String() != "" {
// 				if t.Field(i).Name == "Name" {
// 					like := "%" + v.Field(i).String() + "%"
// 					sqlBuilder.Where(sq.Or{
// 						sq.Like{"name": like},
// 						sq.Like{"nickname": like},
// 					})
// 				} else {
// 					sqlBuilder = sqlBuilder.Where(sq.Eq{t.Field(i).Name: v.Field(i).String()})
// 				}
// 			}
// 		case reflect.Ptr:
// 			if v.Field(i).Elem().Kind() == reflect.Uint ||
// 				v.Field(i).Elem().Kind() == reflect.Uint8 ||
// 				v.Field(i).Elem().Kind() == reflect.Uint64 {
// 				// get *uint8 value
// 				cv := v.Field(i).Elem().Convert(reflect.TypeOf(uint64(0)))
// 				sqlBuilder = sqlBuilder.Where(sq.Eq{t.Field(i).Name: cv.Uint()})
// 			}

// 		}
// 	}
// 	sql := sqlBuilder.Select("id, level, status, name, nickname, avatar, bio, qq, create_time, update_time").From("users")
// 	csql := sqlBuilder.Select("count(id)").From("users").
// 		Offset(uint64(pageSize * (page - 1))).
// 		Limit(uint64(pageSize))

// 	query, arg, _ := sql.ToSql()
// 	cq, ca, _ := csql.ToSql()

// 	log.Print(query)

// 	if err := db.DB.Select(&users, query, arg...); err != nil {
// 		return nil, 0, err
// 	}
// 	if err := db.DB.QueryRowx(cq, ca...).Scan(&count); err != nil {
// 		return nil, 0, err
// 	}

// 	return users, count, nil
// }
