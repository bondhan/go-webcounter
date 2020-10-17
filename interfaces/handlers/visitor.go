package handlers

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/bondhan/go-webcounter/application"
	"github.com/bondhan/go-webcounter/interfaces/respond"
)

// RouteVisitor ...
func RouteVisitor(vh VisitorHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", vh.GetLastCounter)
	r.Get("/increment", vh.IncrementCounter)
	r.Get("/db/", vh.GetLastCounterDB)
	r.Get("/db/increment", vh.IncrementCounterDB)

	return r
}

// VisitorHandler ...
type visitorHandler struct {
	visitorApp application.VisitorApp
}

type VisitorHandler interface {
	GetLastCounter(w http.ResponseWriter, r *http.Request)
	IncrementCounter(w http.ResponseWriter, r *http.Request)
	GetLastCounterDB(w http.ResponseWriter, r *http.Request)
	IncrementCounterDB(w http.ResponseWriter, r *http.Request)
}

// NewVisitorHandler ..
func NewVisitorHandler(visitorApp application.VisitorApp) VisitorHandler {
	return &visitorHandler{
		visitorApp: visitorApp,
	}
}

func (h *visitorHandler) GetLastCounter(w http.ResponseWriter, r *http.Request) {

	respond.JSON(w, http.StatusOK, nil)
}

func (h *visitorHandler) IncrementCounter(w http.ResponseWriter, r *http.Request) {

	respond.JSON(w, http.StatusOK, nil)
}

func (h *visitorHandler) GetLastCounterDB(w http.ResponseWriter, r *http.Request) {

	visitor, err := h.visitorApp.GetLastCounterFromDB()
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	respond.JSON(w, http.StatusOK, visitor)
}

func (h *visitorHandler) IncrementCounterDB(w http.ResponseWriter, r *http.Request) {

	visitor, err := h.visitorApp.IncrementCounterDirectDB()
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	respond.JSON(w, http.StatusOK, visitor)
}
