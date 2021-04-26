package app

import (
	polo2 "github.com/voicurobert/golang-microservices/src/api/controllers/polo"
	repositories2 "github.com/voicurobert/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo2.Polo)
	router.POST("/repositories", repositories2.CreateRepo)
}
