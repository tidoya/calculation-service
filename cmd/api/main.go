package main

import (
	"calculation-service/internal/handler"
	Calculation "calculation-service/internal/transport/http/server"
	"log"
)

//	@title Calculator APP REST API
//	@version 1.0
//	@description API server for Calculator application

//	@host localhost:8080
//	@BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Autorization

// Данные, которые будем отдавать в JSON
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {
	handlers := new(handler.Handler)
	server := new(Calculation.Server)

	if err := server.RunServer("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}

}
