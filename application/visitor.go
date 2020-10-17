package application

import (
	"github.com/bondhan/go-webcounter/domain"
	"github.com/bondhan/go-webcounter/domain/repository"
)

// VisitorApp ...
type VisitorApp interface {
	GetLastCounterFromDB() (domain.Visitor, error)
	IncrementCounterDirectDB() (domain.Visitor, error)
}

type visitorApp struct {
	visitorRepo repository.VisitorRepository
}

// NewVisitorApp ...
func NewVisitorApp(visitorRepo repository.VisitorRepository) VisitorApp {
	return &visitorApp{
		visitorRepo: visitorRepo,
	}
}

func (va *visitorApp) GetLastCounterFromDB() (domain.Visitor, error) {

	visitorCounter, err := va.visitorRepo.GetVisitor()

	return visitorCounter, err
}

func (va *visitorApp) IncrementCounterDirectDB() (domain.Visitor, error) {

	writtenCounter, err := va.visitorRepo.IncrementVisitorAndDeleteOlderVisitor()

	return writtenCounter, err
}
