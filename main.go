package main

import (
	"log"

	"github.com/MarceloLima11/VirtualWallet/db"
	"github.com/MarceloLima11/VirtualWallet/routes"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	routes.Init()
}
