package main

import (
	"app/pkg/server"
	"log"
)

const PORT = ":5000"

func main() {
	server := server.New()
	log.Fatal(server.Listen(PORT))
}
