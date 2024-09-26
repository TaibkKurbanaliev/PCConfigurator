package main

import (
	"context"
	"log"
	"pcbuilder/config"
	"pcbuilder/internal/service"
	"pcbuilder/internal/storage"
	"pcbuilder/internal/transport/rest"
	"pcbuilder/internal/transport/rest/handler"
	"time"
)

// @title PCConfigurator
// @version @0.1
// @descripton Api server for PCConfigurator WebSite

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	conf := config.NewConfig("config/TestConfiguration.json")
	if conf == nil {
		log.Panic("err")
		return
	}

	port, err := conf.Get("Port")
	if err != nil {
		log.Panic(err)
	}

	dbPort, err := conf.Get("DB.Port")
	if err != nil {
		log.Panic(err)
	}

	host, err := conf.Get("DB.Host")
	if err != nil {
		log.Panic(err)
	}

	user, err := conf.Get("DB.User")
	if err != nil {
		log.Panic(err)
	}

	password, err := conf.Get("DB.Password")
	if err != nil {
		log.Panic(err)
	}

	dbName, err := conf.Get("DB.Name")
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbConfiguration := storage.ConfigDB{Host: host, Port: dbPort, User: user, Password: password, DBName: dbName}

	client, err := storage.NewMongoClient(&ctx, dbConfiguration)
	if err != nil {
		log.Panic(err)
	}

	repo := storage.NewRepository(client.Database(dbName))
	newService := service.NewService(repo)
	handlers := handler.NewHandler(newService)

	handlers.InitRoutes()

	srv := new(rest.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
