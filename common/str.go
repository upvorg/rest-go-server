package common

import "regexp"

func CheckUserName(name string) bool {
	math, _ := regexp.Match(`^[a-zA-Z0-9_]{4,8}$`, []byte(name))
	return math
}

func CheckPassword(pwd string) bool {
	math, _ := regexp.Match(`^[a-zA-Z0-9./_]{8,16}$`, []byte(pwd))
	return math
}
