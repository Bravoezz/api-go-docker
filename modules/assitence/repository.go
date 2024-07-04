package assitence

import (
	"fmt"
	"github.com/Bravoezz/infra_srv_tic/db"
	"github.com/Bravoezz/infra_srv_tic/models"
)

type AssistanceRepository struct{}

func NewAssistanceRepository() *AssistanceRepository {
	return new(AssistanceRepository)
}

func (rp AssistanceRepository) List() (*[]models.Assistance, error) {
	assitances := []models.Assistance{}
	err := db.DB.Select(&assitances, "SELECT * FROM asistencias")
	return &assitances, err
}

func (rp AssistanceRepository) FindById(id int) (*models.Assistance, error) {
	assistance := models.Assistance{}
	err := db.DB.Get(&assistance, "SELECT * FROM asistencias where id=$1", id)
	return &assistance, err
}

func (rp AssistanceRepository) Insert(assistance *models.Assistance) (*models.Assistance, error) {

	var query = `
		INSERT INTO asistencias (dni, nombre, apellido, cargo, h_ingreso, ini_refrigerio, fin_refrigerio, h_salida)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`

	fmt.Printf("%v\n", *assistance)

	var id int
	err := db.DB.QueryRowx(query,
		assistance.Dni,
		assistance.Nombre,
		assistance.Apellido,
		assistance.Cargo,
		assistance.HIngreso,
		assistance.IniRefrigerio,
		assistance.FinRefrigerio,
		assistance.HSalida,
	).Scan(&id)

	if err != nil {
		return nil, err
	}
	return rp.FindById(id)
}
