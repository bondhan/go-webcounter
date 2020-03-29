package manager

import (
	"github.com/bondhan/go-webcounter/application"
	"github.com/bondhan/go-webcounter/config"
	"github.com/bondhan/go-webcounter/infrastructure/persistence"
	"github.com/bondhan/go-webcounter/interfaces/handlers"
	"github.com/bondhan/go-webcounter/internal/utils"
	"go.uber.org/dig"
)

// Manager ...
type Manager struct {
	Container *dig.Container
}

// New ...
func New() *Manager {
	container := dig.New()
	utils.PanicErr(container.Provide(config.NewDbConfig))
	utils.PanicErr(container.Provide(persistence.NewVisitorRepository))
	utils.PanicErr(container.Provide(application.NewVisitorApp))
	utils.PanicErr(container.Provide(handlers.NewCommonHandler))
	utils.PanicErr(container.Provide(handlers.NewVisitorHandler))

	return &Manager{Container: container}
}
