package service

import (
	"database/sql"
	"errors"

	"upv.life/server/common"
	"upv.life/server/db"
	"upv.life/server/model"
)

func Register(user *model.User) (*model.User, error) {

	valid, err := CheckUserAndPwd(user.Name, user.Pwd)
	if !valid {
		return nil, err
	}

	user.Pwd = common.HashAndSalt([]byte(user.Pwd))

	// create user
	result := db.Orm.Create(&model.User{
		Name:     user.Name,
		Nickname: user.Nickname,
		Pwd:      user.Pwd,
		Level:    4,
		Status:   2,
	})

	if result.Error != nil {
		return nil, result.Error
	} else {
		if baseuser, error := GetUserByName(user.Name); error != nil {
			return nil, error
		} else {
			return baseuser, nil
		}

	}
}

func IsUserExistByName(name string) bool {
	if user, err := GetUserByName(name); err != nil && err != sql.ErrNoRows {
		return false
	} else {
		return user.ID != 0
	}
}

func GetUserByName(name string) (*model.User, error) {
	var user model.User
	if err := db.Orm.Where("name = ?", name).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckUserName(name string) (bool, error) {
	if !common.CheckUserName(name) {
		return false, errors.New("the user name is invalid")
	}
	if IsUserExistByName(name) {
		return false, errors.New("the user name already taken")
	}
	return true, nil
}

func CheckUserAndPwd(name, pwd string) (bool, error) {
	if !common.CheckUserName(name) {
		return false, errors.New("the user name is invalid")
	}

	if !common.CheckPassword(pwd) {
		return false, errors.New("the password is invalid")
	}

	if IsUserExistByName(name) {
		return false, errors.New("the user name already taken")
	}
	return true, nil
}
