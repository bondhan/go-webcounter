package main

import (
	"fmt"
	"os"

	"github.com/bondhan/go-webcounter/infrastructure/driver"
	"github.com/bondhan/go-webcounter/infrastructure/manager"
	"github.com/bondhan/go-webcounter/interfaces"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func main() {
	_, isProd := os.LookupEnv("PRODUCTION_ENV")
	if isProd {
		fmt.Println("PRODUCTION_ENV TRUE", os.Getenv("PRODUCTION_ENV"))
		driver.NewLogDriver(os.Getenv("LOG_NAME"), logrus.ErrorLevel).InitLog()
	} else {
		fmt.Println("PRODUCTION_ENV FALSE", os.Getenv("PRODUCTION_ENV"))
		driver.NewLogDriver(os.Getenv("LOG_NAME"), logrus.TraceLevel).InitLog()
	}

	interfaces.Init(os.Getenv("APPLICATION_PORT"), manager.New())
}
