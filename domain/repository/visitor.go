package repository

import "github.com/bondhan/go-webcounter/domain"

// VisitorRepository ...
type VisitorRepository interface {
	IncrementVisitor(uint64) error
	GetVisitor() (domain.Visitor, error)
	IncrementVisitorNoParam() (domain.Visitor, error)
}
