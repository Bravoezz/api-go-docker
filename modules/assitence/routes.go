package assitence

import "net/http"

func Router() *http.ServeMux {
	assistanceRouter := http.NewServeMux()

	repository := NewAssistanceRepository()
	service := NewAssistanceService(repository)
	controller := NewAssistanceController(service)

	assistanceRouter.HandleFunc("GET /list", controller.ListAstc)
	assistanceRouter.HandleFunc("GET /find-by/{id}", controller.FindById)
	assistanceRouter.HandleFunc("POST /insert", controller.Insert)

	return assistanceRouter
}
