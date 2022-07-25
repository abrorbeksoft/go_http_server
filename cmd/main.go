package main

import (
	"github.com/go_http_server/api"
	"github.com/go_http_server/config"
)

func main() {

	cfg := config.Load()

	router := api.Main()

	router.Run(cfg.HttpPort)
}
