package services

import (
	"github.com/voicurobert/golang-microservices/src/api/domain/github"
	repos "github.com/voicurobert/golang-microservices/src/api/domain/repositories"
	"github.com/voicurobert/golang-microservices/src/api/providers/github_provider"
	"github.com/voicurobert/golang-microservices/src/config"
	"github.com/voicurobert/golang-microservices/src/utils/errors"
	"strings"
)

type reposService struct {
}

type reposServiceInterface interface {
	CreateRepo(request repos.CreateRepoRequest) (*repos.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repos.CreateRepoRequest) (*repos.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	result := repos.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
