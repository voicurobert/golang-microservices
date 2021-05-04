package oauth

import (
	"github.com/voicurobert/golang-microservices/src/api/utils/errors"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"fede": {
			ID:       123,
			Username: "fede",
		},
	}
)

func GetUserByUsernameAndPassword(username, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundApiError("no user found with given username")
	}
	return user, nil
}
