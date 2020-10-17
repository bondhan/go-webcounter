package config

import (
	"os"

	"github.com/bondhan/go-webcounter/infrastructure/driver"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type DBStorage struct {
	DBWrite *gorm.DB
	DBRead  *gorm.DB
}

// NewDbConfig ...
func NewDbConfig() DBStorage {
	//init postgresql database
	postgreDsnWrite := "host=" + os.Getenv("DB_HOST_WRITE") + " port=" + os.Getenv("DB_PORT_WRITE") +
		" user=" + os.Getenv("DB_USER_WRITE") + " dbname=" + os.Getenv("DB_NAME") +
		" password=" + os.Getenv("DB_PASSWORD_WRITE") + " sslmode=" + os.Getenv("DB_SSLMODE_WRITE")
	postgreWrite := driver.NewDbDriver(postgreDsnWrite, "postgres")
	dbWrite, err := postgreWrite.ConnectDatabase()
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}

	postgreDsnRead := "host=" + os.Getenv("DB_HOST_READ") + " port=" + os.Getenv("DB_PORT_READ") +
		" user=" + os.Getenv("DB_USER_READ") + " dbname=" + os.Getenv("DB_NAME") +
		" password=" + os.Getenv("DB_PASSWORD_READ") + " sslmode=" + os.Getenv("DB_SSLMODE_READ")
	postgreRead := driver.NewDbDriver(postgreDsnRead, "postgres")
	dbRead, err := postgreRead.ConnectDatabase()
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}

	_, isProd := os.LookupEnv("PRODUCTION_ENV")
	if isProd {
		dbWrite.LogMode(false)
		dbRead.LogMode(false)
	} else {
		dbWrite.LogMode(true)
		dbRead.LogMode(true)
	}

	return DBStorage{dbWrite, dbRead}
}
