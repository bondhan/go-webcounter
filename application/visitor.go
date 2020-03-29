package application

import (
	"github.com/bondhan/go-webcounter/domain/repository"
)

// VisitorApp ...
type VisitorApp interface {
}

type newsApp struct {
	visitorRepo repository.VisitorRepository
}

// NewVisitorApp ...
func NewVisitorApp(visitorRepo repository.VisitorRepository) VisitorApp {
	return &newsApp{
		visitorRepo: visitorRepo,
	}
}
