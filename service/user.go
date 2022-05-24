package service

import (
	"errors"

	"gorm.io/gorm"
	"upv.life/server/common"
	"upv.life/server/db"
	"upv.life/server/middleware"
	"upv.life/server/model"
)

func Register(user *model.User) (*model.User, string, error) {

	valid, err := CheckUserAndPwd(user.Name, user.Pwd)
	if !valid {
		return nil, "", err
	}

	user.Pwd = common.HashAndSalt([]byte(user.Pwd))
	result := db.Orm.Model(&model.User{}).Create(map[string]interface {
	}{
		"Name":     user.Name,
		"Nickname": user.Nickname,
		"Pwd":      user.Pwd,
		"Level":    4,
		"Status":   2,
	})

	if result.Error != nil {
		return nil, "", result.Error
	} else {
		if baseuser, error := GetUserByName(user.Name); error != nil {
			return nil, "", error
		} else {
			token, _ := middleware.GenerateJwtToken(baseuser.ID, baseuser.Name, baseuser.Level)
			return baseuser, token, nil
		}
	}
}

func IsUserExistByName(name string) bool {
	if _, err := GetUserByName(name); err != nil || err == gorm.ErrRecordNotFound {
		return false
	} else {
		return true
	}
}

func GetUserByName(name string) (*model.User, error) {
	var user model.User
	if err := db.Orm.Where("name = ?", name).First(&user).Error; err != nil {
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
