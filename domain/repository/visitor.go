package repository

import "github.com/bondhan/go-webcounter/domain"

// VisitorRepository ...
type VisitorRepository interface {
	GetVisitor() (domain.Visitor, error)
	IncrementVisitorAndDeleteOlderVisitor() (domain.Visitor, error)
}
