package models

import "time"

type Assistance struct {
	Id            int       `db:"id" json:"id"`
	Dni           string    `db:"dni" json:"dni"`
	Nombre        string    `db:"nombre" json:"nombre"`
	Apellido      string    `db:"apellido" json:"apellido"`
	Cargo         string    `db:"cargo" json:"cargo"`
	HIngreso      time.Time `db:"h_ingreso" json:"h_ingreso"`
	IniRefrigerio time.Time `db:"ini_refrigerio" json:"ini_refrigerio"`
	FinRefrigerio time.Time `db:"fin_refrigerio" json:"fin_refrigerio"`
	HSalida       time.Time `db:"h_salida" json:"h_salida"`
}
