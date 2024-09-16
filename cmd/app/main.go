package main

import (
	"log"
	"pcbuilder/config"
	"pcbuilder/internal/service"
	"pcbuilder/internal/storage"
	"pcbuilder/internal/transport/rest"
	"pcbuilder/internal/transport/rest/handler"
)

func main() {
	repo := storage.NewRepository()
	service := service.NewService(repo)
	handlers := handler.NewHandler(service)

	handlers.InitRoutes()

	conf := config.NewConfig("config/configuration.json")
	port, err := conf.Get("Port")
	if err != nil {
		log.Panic(err)
	}

	srv := new(rest.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
