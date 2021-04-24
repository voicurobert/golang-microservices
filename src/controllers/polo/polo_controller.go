package polo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	polo = "polo"
)

func Polo(c *gin.Context) {
	fmt.Println("abc")
	c.String(http.StatusOK, polo)
}
