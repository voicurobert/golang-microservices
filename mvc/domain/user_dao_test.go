package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization

	// Execution
	user, err := GetUser(0)

	// Validation

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValuesf(t, http.StatusNotFound, err.StatusCode, "we were expecting 404 when user is not found")
	assert.EqualValuesf(t, "not_found", err.Code, "we were expecting 'not_found' code")
	assert.EqualValuesf(t, "user 0 was not found", err.Message, "")

	//if user != nil {
	//	t.Error("we were not expecting a user with id 0")
	//}
	//
	//if err == nil {
	//	t.Error("we were expecting an error when user is 0")
	//}
	//
	//if err.StatusCode != http.StatusNotFound {
	//	t.Error("we were expecting 404 when user is not found")
	//}
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValuesf(t, 123, user.ID, "expecting user_id 123")
	assert.EqualValuesf(t, "Fede", user.FirstName, "")
	assert.EqualValuesf(t, "Leo", user.LastName, "")
	assert.EqualValuesf(t, "email@gmail.com", user.Email, "")
}
