package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "prod" {
		_ = godotenv.Load()
	}

	connectionURL, found := os.LookupEnv("DATABASE_URL")
	if !found {
		log.Println(connectionURL)
		log.Fatalf("Unable to find database url from .env")
	}

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		log.Fatalf("Couldn't connect to database during migrations, error: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Couldn't get driver during migrations, error: %v", err)
	}

	dbName, found := os.LookupEnv("POSTGRES_DB")
	if !found {
		log.Println(dbName)
		log.Fatalf("Unable to find db from .env")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrate/migrations",
		dbName, driver)
	if err != nil {
		log.Fatalf("Migrations failed %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to apply up.sql files to db, err: %v", err)
	}

	// if cmd == "down" {
	// 	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
	// 		log.Fatalf("failed to apply down.sql files to db, err: %v", err)
	// 	}
	// }

	version, dirty, err := m.Version()
	fmt.Printf("Current migration version: %v, dirty: %v, err: %v\n", version, dirty, err)
}
