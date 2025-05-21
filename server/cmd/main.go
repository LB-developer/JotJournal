package main

import (
	"log"

	"github.com/lb-developer/jotjournal/cmd/api"
	"github.com/lb-developer/jotjournal/db"
	"github.com/lb-developer/jotjournal/service/session"
)

func main() {
	dbPool, err := db.NewPgxPool()
	if err != nil {
		log.Fatal(err)
	}

	defer dbPool.Close()
	
	cache, err := session.NewCache()
	if err != nil {
		log.Fatal(err)
	}
	defer cache.Close()

	server := api.NewAPIServer(":8080", dbPool, cache)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
