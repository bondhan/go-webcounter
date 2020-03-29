package persistence

import (
	"github.com/bondhan/go-webcounter/config"
	"github.com/bondhan/go-webcounter/domain"
	"github.com/bondhan/go-webcounter/domain/repository"
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
	err := v.dbRead.Where("id = ?", 1).First(&visitor).Error

	return visitor, err
}

func (v *visitorRepository) WriteVisitor(visitor domain.Visitor) error {
	err := v.dbWrite.Save(&visitor).Error

	return err
}
