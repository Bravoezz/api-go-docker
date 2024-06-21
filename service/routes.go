package service

import (
	"github.com/Bravoezz/infra_srv_tic/service/product"
	"github.com/Bravoezz/infra_srv_tic/service/user"
	"net/http"
)

func RegisterRoutes(v1Router *http.ServeMux) {
	v1Router.Handle("/user/", http.StripPrefix("/user", user.Router()))
	v1Router.Handle("/product/", http.StripPrefix("/product", product.Router()))
}
