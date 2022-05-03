package common

import (
	"regexp"
	"strings"
)

func CheckUserName(name string) bool {
	math, _ := regexp.Match(`^[a-zA-Z0-9_]{4,8}$`, []byte(name))
	return math
}

func CheckPassword(pwd string) bool {
	math, _ := regexp.Match(`^[a-zA-Z0-9./_]{8,16}$`, []byte(pwd))
	return math
}

func HasTag(tags string, tag string) bool {
	tags = strings.Trim(tags, " ")
	if tags == "" {
		return false
	}
	arr := strings.Split(tags, " ")
	for _, v := range arr {
		if v == tag {
			return true
		}
	}
	return false
}

func FixTags(target string, tags string) string {
	tags = strings.Trim(tags, " ")
	if tags == "" {
		return ""
	}
	arr := strings.Split(tags, " ")
	var result string = ""
	for _, v := range arr {
		if HasTag(target, v) {
			result += v + " "
		}
	}

	return strings.Trim(result, " ")
}
