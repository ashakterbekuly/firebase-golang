package api

var userState bool

func InitUserState() {
	userState = false
}

func GetUserState() bool {
	return userState
}

func SetUserState(status bool) {
	userState = status
}
