package user

import (
	"net/http"
	"time"
)

func Router() *http.ServeMux {
	userRouter := http.NewServeMux()

	userRouter.HandleFunc("GET /list", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET de user /user"))
	})

	userRouter.HandleFunc("GET /find-by/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("userId")
		time.Sleep(time.Millisecond * 500)
		w.Write([]byte("GET de user /find by id " + userId))
	})

	return userRouter
}
