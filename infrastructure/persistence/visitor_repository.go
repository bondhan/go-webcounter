package persistence

import (
	"github.com/bondhan/go-webcounter/domain"
	"github.com/bondhan/go-webcounter/domain/repository"
	"github.com/bondhan/go-webcounter/infrastructure/config"
	"github.com/jinzhu/gorm"
)

type visitorRepository struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
}

//NewVisitorRepository ...
func NewVisitorRepository(dbs config.DBStorage) repository.VisitorRepository {
	return &visitorRepository{
		dbWrite: dbs.DBWrite,
		dbRead:  dbs.DBRead,
	}
}

func (v *visitorRepository) GetVisitor() (domain.Visitor, error) {

	var visitor domain.Visitor
	err := v.dbRead.Order("id desc").First(&visitor).Error

	return visitor, err
}

func (v *visitorRepository) IncrementVisitorAndDeleteOlderVisitor() (domain.Visitor, error) {

	var OldVisitor, NewVisitor domain.Visitor

	tx := v.dbWrite.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// get the last visitor
	err := tx.Order("id desc").First(&OldVisitor).Error
	if err != nil {
		tx.Rollback()
		return NewVisitor, tx.Error
	}

	NewVisitor.Counter = OldVisitor.Counter+1

	err = tx.Create(&NewVisitor).Error
	if err != nil {
		tx.Rollback()
		return NewVisitor, tx.Error
	}

	err=tx.Delete(&OldVisitor).Error
	if err != nil {
		tx.Rollback()
		return NewVisitor, tx.Error
	}

	return NewVisitor, tx.Commit().Error
}
