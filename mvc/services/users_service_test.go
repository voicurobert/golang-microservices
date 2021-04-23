package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/voicurobert/golang-microservices/mvc/domain"
	"github.com/voicurobert/golang-microservices/mvc/utils"
	"net/http"
	"testing"
)

var (
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct {
}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	fmt.Println("In usersDaoMockMethod")
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 was not found",
			StatusCode: http.StatusNotFound,
			Code:       "",
		}
	}

	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValuesf(t, http.StatusNotFound, err.StatusCode, "")
	assert.EqualValuesf(t, "user 0 was not found", err.Message, "")
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID:        123,
			FirstName: "",
			LastName:  "",
			Email:     "",
		}, nil
	}
	user, err := UsersService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValuesf(t, 123, user.ID, "")
}
