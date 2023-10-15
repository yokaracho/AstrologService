package main

import (
	astrologservice "AstrologService"
	handler2 "AstrologService/pkg/handler"
	postgres "AstrologService/pkg/repository/postgres"
	service2 "AstrologService/pkg/service"
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := GetLogger()

	if err := godotenv.Load("/home/dev/AstrologService/.env"); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}
	logger.Infof("env variables successfully loaded")

	postgresCfg := postgres.Config{
		Username: "postgres",
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("POST_PORT"),
		DBName:   os.Getenv("DATABASE"),
	}

	postgresPool, err := postgres.NewConnectionPool(context.Background(), postgresCfg)
	if err != nil {
		logger.Fatalf("database connection error: %s", err.Error())
	}
	logger.Infof("database connection successfully ")

	repository := postgres.NewRepository(postgresPool)
	service := service2.NewService(repository)
	handler := handler2.NewHandler(service)

	server := new(astrologservice.Server)

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			logrus.Infof("server shutdown error: %s", err.Error())
		}
	}()

	if err := server.Run(os.Getenv("PORT"), handler.GetRouter()); err != nil {
		log.Fatalf("server running error: %s", err.Error())
	} else {
		log.Printf("Server is running on port %s", os.Getenv("PORT"))
	}
}

func GetLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.SetReportCaller(true)

	return logger
}
