package main

import (
	"github.com/abrorbeksoft/go_http_server/api"
	"github.com/gin-gonic/gin"
)

func main() {
	var route *gin.Engine
	route = api.Main()

	route.Run(":3000")
}
