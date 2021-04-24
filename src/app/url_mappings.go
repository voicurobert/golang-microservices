package app

import (
	"github.com/voicurobert/golang-microservices/src/controllers/polo"
	"github.com/voicurobert/golang-microservices/src/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
