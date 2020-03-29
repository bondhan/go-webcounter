package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

// CommonRoutes ..
func CommonRoutes(r *chi.Mux, ch *CommonHandler) {

	r.Get("/", ch.Hello)

	// healtcheck
	r.Get("/ping", ch.Ping)
}

type CommonHandler struct {
}

// NewCommonHandler ..
func NewCommonHandler() *CommonHandler {
	return &CommonHandler{}
}

func (c *CommonHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (c *CommonHandler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("go-webcounter is alive"))
}
