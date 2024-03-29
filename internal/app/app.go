package app

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	external2 "gitlab.com/a5805/ondeu/ondeu-back/internal/client"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/handler"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/server"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/service"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func Run() {
	logrus.SetReportCaller(true)

	cfg := initConfigs()
	setLogLevel(cfg.LogLevel)

	db := repository.NewRedis(cfg.Cache)

	external := external2.NewExternalService(cfg)
	repo := repository.NewRepository(db)
	services := service.NewServices(cfg, external, repo)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.Port, handlers.Init()); err != nil {
			logrus.Errorf("error occured while running http server %s/n", err.Error())
		}
	}()

	logrus.Printf("server is starting at port: %s", cfg.Port)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	logrus.Printf("server is stopping at port: %s", cfg.Port)

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfigs() *models.AppConfigs {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Error loading .env file")
	}

	cache := &models.Redis{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	api1Percent, err := strconv.Atoi(os.Getenv("API1_PERCENT"))
	if err != nil {
		fmt.Printf("api1 percent is not set: %s\n", err.Error())
	}

	api2Percent, err := strconv.Atoi(os.Getenv("API2_PERCENT"))
	if err != nil {
		fmt.Printf("api2 percent is not set: %s\n", err.Error())
	}

	return &models.AppConfigs{
		Port:     os.Getenv("PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"),
		Cache:    cache,
		Balances: []models.Balance{
			{
				Url:    os.Getenv("URL1"),
				Weight: api1Percent,
			},
			{
				Url:    os.Getenv("URL2"),
				Weight: api2Percent,
			},
		},
	}
}

func setLogLevel(level string) {
	switch level {
	case "debug":
		{
			logrus.SetLevel(logrus.DebugLevel)
			break
		}
	case "info":
		{
			logrus.SetLevel(logrus.InfoLevel)
			break
		}
	case "warn":
		{
			logrus.SetLevel(logrus.WarnLevel)
			break
		}
	case "error":
		{
			logrus.SetLevel(logrus.ErrorLevel)
			break
		}
	case "fatal":
		{
			logrus.SetLevel(logrus.FatalLevel)
			break
		}
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}
