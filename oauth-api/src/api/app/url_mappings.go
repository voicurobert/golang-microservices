package app

import (
	"github.com/voicurobert/golang-microservices/oauth-api/src/api/controllers/oauth"
	"github.com/voicurobert/golang-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
