package product

import "net/http"

func Router() *http.ServeMux {
	productRouter := http.NewServeMux()

	productRouter.HandleFunc("GET /list", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("hola desde product/list"))
		return
	})

	return productRouter
}
