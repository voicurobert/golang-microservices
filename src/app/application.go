package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApplication() {
	fmt.Println("ABC")
	mapUrls()

	//http.HandleFunc("/users", controllers.GetUser)
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	panic(err)
	//}
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
