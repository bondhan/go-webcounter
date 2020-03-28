package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bondhan/go-webcounter/infrastructure/driver"
	"github.com/bondhan/go-webcounter/infrastructure/manager"
	"github.com/bondhan/gobareksa/interfaces"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	manager.GetContainer()
	logrus.Info("manager was called")

	logrus.Info("application started at port:", os.Getenv("APPLICATION_PORT"))
	interfaces.Init(os.Getenv("APPLICATION_PORT"))

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	//
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		webcounter := map[string]string{
			"status":          "go-webcounter alive",
			"server_datetime": time.Now().String(),
		}

		hw, _ := json.Marshal(webcounter)
		w.Write(hw)
	})

	//healthcheck
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/api/v1/visitorCount", getCounter)
	http.ListenAndServe(":8080", r)
}

// getCounter will
func getCounter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("count"))
}

func syncRedisMySql() {

}
