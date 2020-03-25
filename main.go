package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("go-webcounter"))
	})
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Get("/count", getCounter)
	http.ListenAndServe(":8080", r)
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("count"))
}

func syncRedisMySql() {

}
