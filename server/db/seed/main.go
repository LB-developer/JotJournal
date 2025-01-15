package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-faker/faker/v4"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/lb-developer/jotjournal/types"
)

func main() {
	godotenv.Load()

	connectionURL, found := os.LookupEnv("DATABASE_URL")
	if !found {
		log.Println(connectionURL)
		log.Fatalf("Unable to find database url from .env")
	}

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		log.Fatalf("Couldn't connect to database during seeding, error: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Couldn't get driver during seeding, error: %v", err)
	}

	_, err = migrate.NewWithDatabaseInstance(
		"file://db/seed/seeds",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Seeds failed %v", err)
	}

	var task types.Task
	createTask(&task)
	insertUser(db, task)
}

func createTask(task *types.Task) {
	task.Weekly = true
	task.Monthly = false
	task.Daily = false
	task.IsCompleted = false

	err := faker.FakeData(task)
	if err != nil {
		log.Fatalf("Couldn't create fake task, error: %v\n", err)
	}
}

func insertUser(db *sql.DB, task types.Task) int {
	query := `
	INSERT INTO tasks (monthly, weekly, daily, deadline, description, is_completed, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, 1)
	RETURNING id
	`

	lastInsertId := 0
	err := db.QueryRow(
		query,
		task.Monthly,
		task.Weekly,
		task.Daily,
		task.Deadline,
		task.Description,
		task.IsCompleted,
	).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}

	return lastInsertId
}
