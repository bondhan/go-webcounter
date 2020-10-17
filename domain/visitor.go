package domain

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

// Visitor ...
type Visitor struct {
	gorm.Model
	Counter uint64 `gorm:"column:counter" json:"counter"`
}

// TableName sets the insert table name for this struct type
func (m *Visitor) TableName() string {
	return "m_visitor"
}
