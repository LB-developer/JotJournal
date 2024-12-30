package main

import (
	"log"

	"github.com/lb-developer/jotjournal/cmd/api"
	"github.com/lb-developer/jotjournal/db"
)

func main() {
	dbPool, err := db.NewPgxPool()
	if err != nil {
		log.Fatal(err)
	}

	defer dbPool.Close()
	server := api.NewAPIServer(":8080", nil)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
