package services

import (
	"github.com/voicurobert/golang-microservices/mvc/domain"
	"github.com/voicurobert/golang-microservices/mvc/utils"
)

type usersService struct {
}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
