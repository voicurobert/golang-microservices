package services

import (
	"github.com/voicurobert/golang-microservices/src/api/config"
	"github.com/voicurobert/golang-microservices/src/api/domain/github"
	repos "github.com/voicurobert/golang-microservices/src/api/domain/repositories"
	"github.com/voicurobert/golang-microservices/src/api/log/option_b"
	"github.com/voicurobert/golang-microservices/src/api/providers/github_provider"
	"github.com/voicurobert/golang-microservices/src/api/utils/errors"
	"net/http"
	"sync"
)

type reposService struct {
}

type reposServiceInterface interface {
	CreateRepo(clientId string, request repos.CreateRepoRequest) (*repos.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repos.CreateRepoRequest) (repos.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(clientId string, input repos.CreateRepoRequest) (*repos.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	option_b.Info("about to send request to external api",
		option_b.Field("client_id", clientId),
		option_b.Field("status", "pending"),
		option_b.Field("authenticated", clientId != ""))

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		option_b.Error("respond obtained from external api", err,
			option_b.Field("client_id", clientId),
			option_b.Field("status", "error"),
			option_b.Field("authenticated", clientId != ""))
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	option_b.Info("respond obtained from external api",
		option_b.Field("client_id", clientId),
		option_b.Field("status", "success"),
		option_b.Field("authenticated", clientId != ""))

	result := repos.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *reposService) CreateRepos(request []repos.CreateRepoRequest) (repos.CreateReposResponse, errors.ApiError) {
	input := make(chan repos.CreateRepositoriesResult)
	output := make(chan repos.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, currentRequest := range request {
		wg.Add(1)
		go s.createRepoConcurrent(currentRequest, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}
	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(request) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repos.CreateRepositoriesResult, output chan repos.CreateReposResponse) {
	var result repos.CreateReposResponse
	for res := range input {
		repoResult := repos.CreateRepositoriesResult{
			Response: res.Response,
			Error:    res.Error,
		}
		result.Results = append(result.Results, repoResult)
		wg.Done()
	}
	output <- result
}

func (s *reposService) createRepoConcurrent(input repos.CreateRepoRequest, output chan repos.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repos.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo("", input)
	if err != nil {
		output <- repos.CreateRepositoriesResult{Error: err}
		return
	}

	output <- repos.CreateRepositoriesResult{Response: result}
}
