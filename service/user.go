package service

import (
	"database/sql"
	"errors"

	"upv.life/server/common"
	"upv.life/server/db"
	"upv.life/server/model"
)

func Register(user *model.User) (*model.User, error) {
	// check if user exists
	if IsUserExist(user) {
		return nil, errors.New("the user name already taken")
	}

	user.Pwd = common.HashAndSalt([]byte(user.Pwd))

	// create user
	result := db.Orm.Create(user)
	if result.Error != nil {
		return nil, result.Error
	} else {
		if _, error := GetUserByName(user.Name); error != nil {
			return nil, error
		}
		user.Pwd = ""

		return user, nil
	}
}

func IsUserExist(user *model.User) bool {
	if _, err := GetUserByName(user.Name); err != nil && err != sql.ErrNoRows {
		return false
	}
	return user.ID != 0
}

func GetUserByName(name string) (*model.User, error) {
	var user model.User
	if err := db.Orm.Where("name = ?", name).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
