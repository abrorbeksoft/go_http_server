package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Main() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowCredentials = true
	corsCfg.AllowHeaders = append(corsCfg.AllowHeaders, "*")

	r.Use(cors.New(corsCfg))

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
