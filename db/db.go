package db

import (
	"fmt"
	"github.com/Bravoezz/infra_srv_tic/config"
	"github.com/Bravoezz/infra_srv_tic/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var DB *sqlx.DB

func NewPgSqlStorage(pgSqlCnf *config.PgSqlConfig) (*sqlx.DB, error) {
	var err error
	DB, err = sqlx.Connect("postgres", pgSqlCnf.Format())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected db")

	count := 0
	assistance := models.Assistance{}
	rows, _ := DB.Queryx("SELECT * FROM asistencias")
	for rows.Next() {
		err := rows.StructScan(&assistance)
		if err != nil {
			log.Fatal(err)
		}
		count++
		log.Printf("%#v\n", assistance.HSalida.Format(time.ANSIC))
	}

	log.Println("Total de registros -> " + fmt.Sprint(count))

	return DB, nil
}
