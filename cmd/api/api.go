package api

import (
	"database/sql"
	"github.com/Bravoezz/infra_srv_tic/service"
	"log"
	"net/http"
	"time"
)

type ApiServer struct {
	address string
	db      *sql.DB
}

func NewApiServer(adrs string, db *sql.DB) *ApiServer {
	return &ApiServer{adrs, db}
}

func (s *ApiServer) Run() error {
	middlewareChain := MiddlewareChain(
		ReqLoggerMiddleware,
		ReqAllowOriginCorsMiddleware,
	)

	//config server
	v1 := http.NewServeMux()
	v1.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server online ✅"))
		return
	})

	//register routes in api/v1/*
	service.RegisterRoutes(v1)

	main := http.NewServeMux()
	main.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	server := &http.Server{
		Addr:    s.address,
		Handler: middlewareChain(main),
	}

	log.Printf("Server running on port :%s", s.address)
	return server.ListenAndServe()
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}
}

func ReqAllowOriginCorsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allowedOrigin := "http://localhost:3000"
		allowedOrigin := "*"
		origin := r.Header.Get("Origin")

		if allowedOrigin != "*" && origin != allowedOrigin {
			http.Error(w, "CORS policy: Not allowed", http.StatusForbidden)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	}
}

func ReqLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("method: %s, path: %s, time: %s", r.Method, r.URL.Path, time.Since(start))
	}
}
