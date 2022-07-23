package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello from server"})
	})

	r.GET("/name", func(c *gin.Context) {
		name := c.DefaultQuery("name", "")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	return r
}
