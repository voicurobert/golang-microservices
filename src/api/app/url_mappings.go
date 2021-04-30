package app

import (
	polo "github.com/voicurobert/golang-microservices/src/api/controllers/polo"
	repositories "github.com/voicurobert/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
