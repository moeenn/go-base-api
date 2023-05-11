package main

import (
	"app/pkg/config"
	"app/pkg/server"
	"log"
)

func main() {
	server := server.New()

	err := server.Listen(config.AppConfig.Port)
	if err != nil {
		log.Fatal(err)
	}
}
