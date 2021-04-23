package services

import (
	"github.com/voicurobert/golang-microservices/mvc/domain"
	"github.com/voicurobert/golang-microservices/mvc/utils"
	"net/http"
)

type itemsService struct {
}

var (
	ItemsService itemsService
)

func (i *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "not implemented",
		StatusCode: http.StatusInternalServerError,
		Code:       "",
	}
}
