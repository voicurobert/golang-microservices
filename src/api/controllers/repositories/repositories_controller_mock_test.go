package repositories

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	repos "github.com/voicurobert/golang-microservices/src/api/domain/repositories"
	"github.com/voicurobert/golang-microservices/src/api/services"
	"github.com/voicurobert/golang-microservices/src/api/utils/errors"
	"github.com/voicurobert/golang-microservices/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type repoServiceMock struct {
}

var (
	funcCreateRepo  func(request repos.CreateRepoRequest) (*repos.CreateRepoResponse, errors.ApiError)
	funcCreateRepos func(request []repos.CreateRepoRequest) (repos.CreateReposResponse, errors.ApiError)
)

func init() {

}

func (sm *repoServiceMock) CreateRepo(clientId string, request repos.CreateRepoRequest) (*repos.CreateRepoResponse, errors.ApiError) {
	return funcCreateRepo(request)
}

func (sm *repoServiceMock) CreateRepos(request []repos.CreateRepoRequest) (repos.CreateReposResponse, errors.ApiError) {
	return funcCreateRepos(request)
}

func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(request repos.CreateRepoRequest) (*repos.CreateRepoResponse, errors.ApiError) {
		return &repos.CreateRepoResponse{
			Name:  "mocked service",
			Owner: "golang",
			ID:    321,
		}, nil
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))

	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)
	var result repos.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 321, result.ID)
	assert.EqualValues(t, "mocked service", result.Name)
	assert.EqualValues(t, "golang", result.Owner)
}

func TestCreateRepoErrorFromGithubMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(request repos.CreateRepoRequest) (*repos.CreateRepoResponse, errors.ApiError) {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))

	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid repository name", apiErr.Message())
}
