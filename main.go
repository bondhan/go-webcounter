package main

import (
	"context"
	"github.com/bondhan/go-webcounter/infrastructure/config"
	"github.com/bondhan/go-webcounter/infrastructure/utils/redisclient"
	"github.com/bondhan/go-webcounter/interfaces/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bondhan/go-webcounter/application"
	"github.com/bondhan/go-webcounter/infrastructure/driver"
	"github.com/bondhan/go-webcounter/infrastructure/persistence"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func main() {
	_, isProd := os.LookupEnv("PRODUCTION_ENV")
	if isProd {
		driver.NewLogDriver(os.Getenv("LOG_NAME"), logrus.WarnLevel).InitLog()
	} else {
		driver.NewLogDriver(os.Getenv("LOG_NAME"), logrus.TraceLevel).InitLog()
	}

	//initialize instance of DB transaction and DB log
	dbWebCounter := config.NewDbConfig()

	defer dbWebCounter.DBWrite.Close()
	defer dbWebCounter.DBRead.Close()

	//initialize routes using gochi
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Duration(90 * time.Second)))

	_, err := redisclient.PingRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_EXPIRED_SECONDS"), os.Getenv("REDIS_PASSWORD"))
	if err != nil {
		logrus.Fatalf("Fail connecting to redisclient error:%s", err)
	}

	visitorRepo := persistence.NewVisitorRepository(dbWebCounter)
	visitorApp := application.NewVisitorApp(visitorRepo)
	visitorHandler := handlers.NewVisitorHandler(visitorApp)

	r.Mount("/web-counter", handlers.RouteVisitor(visitorHandler))

	server := &http.Server{Addr: ":" + os.Getenv("APPLICATION_PORT"), Handler: r}
	go func() {
		logrus.Info("application started at port:", os.Getenv("APPLICATION_PORT"))
		if err := server.ListenAndServe(); err != nil {
			logrus.Fatal(err)
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	<-stop
	logrus.Warnf("got signal: %v, closing DB connection gracefully", stop)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logrus.Warn("shutting down http server")
	if err := server.Shutdown(ctx); err != nil {
		logrus.Error(err)
	}
}
