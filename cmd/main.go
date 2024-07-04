package main

import (
	"github.com/Bravoezz/infra_srv_tic/cmd/api"
	"github.com/Bravoezz/infra_srv_tic/config"
	"github.com/Bravoezz/infra_srv_tic/db"

	//"github.com/Bravoezz/infra_srv_tic/db"
	"log"
)

func main() {
	var port string
	if config.GetEnv("PORT") == "" {
		port = "8000"
	} else {
		port = config.GetEnv("PORT")
	}

	dbCon, _ := db.NewPgSqlStorage(&config.PgSqlConfig{
		Host:     config.GetEnv("DB_HOST"),
		Password: config.GetEnv("DB_PASSWORD"),
		DbName:   config.GetEnv("DB_NAME"),
		User:     config.GetEnv("DB_USER_NAME"),
	})
	defer dbCon.Close()
	server := api.NewApiServer(":"+port, dbCon)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
