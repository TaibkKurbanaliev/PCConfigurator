package main

import (
	"log"
	"pcbuilder/internal/transport/rest"
	"pcbuilder/internal/transport/rest/handler"
)

func main() {
	handlers := new(handler.Handler)
	handlers.InitRoutes()
	srv := new(rest.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
