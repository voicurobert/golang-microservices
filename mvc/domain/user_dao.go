package domain

import (
	"fmt"
	"github.com/voicurobert/golang-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: &User{
			ID:        123,
			FirstName: "Fede",
			LastName:  "Leo",
			Email:     "email@gmail.com",
		},
	}
)

type usersDaoInterface interface {
	GetUser(userId int64) (*User, *utils.ApplicationError)
}

type userDao struct {
}

var (
	UserDao usersDaoInterface = &userDao{}
)

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we are accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
