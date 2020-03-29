package repository

import "github.com/bondhan/go-webcounter/domain"

// VisitorRepository ...
type VisitorRepository interface {
	GetVisitor() (domain.Visitor, error)
	WriteVisitor(visitor domain.Visitor) error
}
