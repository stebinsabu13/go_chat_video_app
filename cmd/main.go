package main

import (
	"log"

	"github.com/stebin13/go_chat_video_app/pkg/config"
	"github.com/stebin13/go_chat_video_app/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
