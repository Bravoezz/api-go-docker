package models

type Product struct {
	Id                 string `db:"id"`
	Nombre             string `db:"nombre"`
	Descripcion        string `db:"descripcion"`
	CantidadDisponible string `db:"cantidadDisponible"`
	PrecioUnitario     string `db:"precioUnitario"`
	Proveedor          string `db:"proveedor"`
}
