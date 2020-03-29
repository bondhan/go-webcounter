package interfaces

import (
	"net/http"
	"os"

	"github.com/bondhan/go-webcounter/infrastructure/manager"
	"github.com/bondhan/go-webcounter/interfaces/handlers"
	"github.com/bondhan/go-webcounter/internal/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

// Init ...
func Init(port string, manager *manager.Manager) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	err := manager.Container.Invoke(func(ch *handlers.CommonHandler) {
		handlers.CommonRoutes(r, ch)
	})

	if err != nil {
		panic(err)
	}

	//bind the router and function for common
	utils.PanicErr(manager.Container.Invoke(func(ch *handlers.CommonHandler) {
		handlers.CommonRoutes(r, ch)
	}))

	//bind the router and function for visitor
	utils.PanicErr(manager.Container.Invoke(func(vh *handlers.VisitorHandler) {
		handlers.VisitorRoutes(r, vh)
	}))

	logrus.Info("application started at port:", os.Getenv("APPLICATION_PORT"))
	utils.PanicErr(http.ListenAndServe(":"+port, r))
}
