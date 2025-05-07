package main

import (
	"log"

	"github.com/Code-Linx/Go-Ecommerce-Backend-Api/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	server.Run()
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
