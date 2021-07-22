package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello from server"})
	})

	return router

}
