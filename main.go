package main

import (
	"log"

	"github.com/CelmarPA/api-students/api"
)

func main() {
	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}
