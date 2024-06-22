package main

import (
	"github.com/Bravoezz/infra_srv_tic/cmd/api"
	"github.com/Bravoezz/infra_srv_tic/cmd/config"
	"log"
)

func main() {
	var port string
	if config.GetEnv("PORT") == "" {
		port = "8000"
	} else {
		port = config.GetEnv("PORT")
	}

	server := api.NewApiServer(":"+port, nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
