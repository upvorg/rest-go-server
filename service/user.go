package service

import (
	"database/sql"
	"errors"

	"upv.life/server/model"
)

func Register(user *model.User) (*model.User, error) {
	// check if user exists
	if err := user.GetUserByName(); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if user.ID != 0 {
		return nil, errors.New("The user name already taken.")
	}

	//TODO: password encryption

	// create user
	if _, error := user.Create(); error != nil {
		return nil, error
	} else {
		if user.GetUserByID(); error != nil {
			return nil, error
		}
		user.Pwd = ""

		return user, nil
	}
}
