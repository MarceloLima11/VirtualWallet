package main

import (
	"log"

	"github.com/MarceloLima11/VirtualWallet/postgres"
	"github.com/MarceloLima11/VirtualWallet/routes"
)

func main() {
	if err := postgres.Init(); err != nil {
		log.Fatal(err)
	}

	if err := routes.Init(); err != nil {
		log.Fatal(err)
	}
}
