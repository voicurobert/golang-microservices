package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/voicurobert/golang-microservices/oauth-api/src/api/domain/oauth"
	"github.com/voicurobert/golang-microservices/oauth-api/src/api/services"
	"github.com/voicurobert/golang-microservices/src/api/utils/errors"
	"net/http"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, token)
}

func GetAccessToken(c *gin.Context) {
	tokenId := c.Param("token_id")
	token, err := services.OauthService.GetAccessToken(tokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}
