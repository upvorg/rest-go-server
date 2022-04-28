package common

func IsRoot(role uint) bool {
	return role == 1
}

func IsAdmin(role uint) bool {
	return role == 2
}

func IsCreator(role uint) bool {
	return role == 3
}

func IsUser(role uint) bool {
	return role == 4
}
