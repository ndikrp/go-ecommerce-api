package main

import (
	"github.com/ndikrp/go-ecommerce-api/cmd/api"
	"log"
)

func main() {
	server := api.NewApiServer(":9000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
