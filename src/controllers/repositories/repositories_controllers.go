package repositories

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/voicurobert/golang-microservices/src/api/domain/repositories"
	"github.com/voicurobert/golang-microservices/src/services"
	"github.com/voicurobert/golang-microservices/src/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	fmt.Println(c.Request.Response.Body)
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
