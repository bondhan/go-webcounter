package handlers

import (
	"net/http"

	"github.com/bondhan/go-webcounter/application"
	"github.com/bondhan/go-webcounter/domain/repository"
	"github.com/bondhan/go-webcounter/interfaces/respond"
	"github.com/go-chi/chi"
)

// VisitorHandler ...
type VisitorHandler struct {
	visitorApp  application.VisitorApp
	visitorRepo repository.VisitorRepository
}

// VisitorRoutes ...
func VisitorRoutes(r *chi.Mux, vh *VisitorHandler) {
	r.Route("/api/v1/visitor", func(r chi.Router) {
		r.Get("/getCounter", vh.getCounter)
	})
}

// NewVisitorHandler ..
func NewVisitorHandler(visitorApp application.VisitorApp, visitorRepo repository.VisitorRepository) *VisitorHandler {
	return &VisitorHandler{
		visitorApp:  visitorApp,
		visitorRepo: visitorRepo,
	}
}

func (h *VisitorHandler) getCounter(w http.ResponseWriter, r *http.Request) {
	v, _ := h.visitorRepo.GetVisitor()

	respond.JSON(w, http.StatusOK, v)
}
