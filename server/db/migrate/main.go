package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../../.env")

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

	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrate/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Migrations failed %v", err)
	}
	if err := m.Up(); err != nil {
		log.Fatalf("failed to apply up.sql files to db, err: %v", err)
	}
}
