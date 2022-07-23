package main

import (
	"github.com/abrorbeksoft/go_http_server/api"
)

func main() {
	router := api.Main()

	router.Run(":3000")
}
