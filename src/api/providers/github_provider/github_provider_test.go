package github_provider

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/voicurobert/golang-microservices/src/api/clients/rest_client"
	"github.com/voicurobert/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	rest_client.StartMockups()
	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	assert.EqualValuesf(t, "Authorization", headerAuthorization, "")
	assert.EqualValuesf(t, "token %s", headerAuthorizationFormat, "")
	assert.EqualValuesf(t, "https://api.github.com/user/repos", urlCreateRepo, "")
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	//rest_client.StartMockups() no need to call this anymore because we call it in TestMain()
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response:   nil,
		Err:        errors.New("invalid rest client response"),
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValuesf(t, http.StatusInternalServerError, err.StatusCode, "")
	assert.EqualValuesf(t, "invalid rest client response", err.Message, "")
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	//rest_client.StartMockups() no need to call this anymore because we call it in TestMain()
	rest_client.FlushMockups()

	invalidCLoser, _ := os.Open("-asf3")
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCLoser,
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValuesf(t, http.StatusInternalServerError, err.StatusCode, "")
	assert.EqualValuesf(t, "invalid argument", err.Message, "")
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	//rest_client.StartMockups() no need to call this anymore because we call it in TestMain()
	rest_client.FlushMockups()

	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":1}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValuesf(t, http.StatusInternalServerError, err.StatusCode, "")
	assert.EqualValuesf(t, "invalid json response body", err.Message, "")
}

func TestCreateRepoUnauthorized(t *testing.T) {
	//rest_client.StartMockups() no need to call this anymore because we call it in TestMain()
	rest_client.FlushMockups()

	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires authentication", "documentation_url":"https://developer.github.com/"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValuesf(t, http.StatusUnauthorized, err.StatusCode, "")
	assert.EqualValuesf(t, "Requires authentication", err.Message, "")
}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	//rest_client.StartMockups() no need to call this anymore because we call it in TestMain()
	rest_client.FlushMockups()

	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":"123"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValuesf(t, http.StatusInternalServerError, err.StatusCode, "")
	assert.EqualValuesf(t, "error when trying to unmarshalling github create repo response", err.Message, "")
}

func TestCreateRepoNoError(t *testing.T) {
	//rest_client.StartMockups() no need to call this anymore because we call it in TestMain()
	rest_client.FlushMockups()

	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123, "name": "golang-tutorial", "full_name":"voicurobert"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValuesf(t, 123, response.ID, "")
	assert.EqualValuesf(t, "golang-tutorial", response.Name, "")
	assert.EqualValuesf(t, "voicurobert", response.FullName, "")
}
