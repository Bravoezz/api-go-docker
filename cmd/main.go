package main

import (
	"github.com/Bravoezz/infra_srv_tic/cmd/api"
	"log"
)

func main() {
	server := api.NewApiServer(":8000", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
