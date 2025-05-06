package main

import (
	"calculation-service/internal/handler"
	Calculation "calculation-service/internal/transport/http/server"
	"calculation-service/repository"
	"calculation-service/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

//	@title Calculator APP REST API
//	@version 1.0
//	@description API server for Calculator application

//	@host localhost:8080
//	@BasePath /

//	@securityDefinitions.apikey ApiKeyAuth
//	@in header
//	@name Autorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("DataBasePostgress.host"),
		Port:     viper.GetString("DataBasePostgress.port"),
		Username: viper.GetString("DataBasePostgress.username"),
		Password: viper.GetString("DataBasePostgress.password"),
		DBName:   viper.GetString("DataBasePostgress.dbname"),
		SSLMode:  viper.GetString("DataBasePostgress.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	server := new(Calculation.Server)
	if err := server.RunServer(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
