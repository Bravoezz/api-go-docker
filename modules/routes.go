package modules

import (
	"github.com/Bravoezz/infra_srv_tic/modules/assitence"
	"github.com/Bravoezz/infra_srv_tic/modules/product"
	"github.com/Bravoezz/infra_srv_tic/modules/user"
	"net/http"
)

func RegisterRoutes(v1Router *http.ServeMux) {
	v1Router.Handle("/user/", http.StripPrefix("/user", user.Router()))
	v1Router.Handle("/product/", http.StripPrefix("/product", product.Router()))
	v1Router.Handle("/assistance/", http.StripPrefix("/assistance", assitence.Router()))
}
