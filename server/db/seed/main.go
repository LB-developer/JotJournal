package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
		"file://db",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Seeds failed %v", err)
	}

	var number int64
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("How many fake tasks to generate: ")

		scanner.Scan()

		text := scanner.Text()
		if len(text) != 0 {
			number, err = strconv.ParseInt(text, 10, 64)
			if err != nil {
				fmt.Println("invalid number")
				continue
			} else {
				break
			}
		} else {
			break
		}

	}

	for i := 0; i < int(number); i++ {
		var task types.Task
		createTask(&task)
		insertTask(db, task)
	}
}

func createTask(task *types.Task) {
	task.Weekly = true
	task.Monthly = false
	task.Daily = false
	task.IsCompleted = false

	theTime := time.Now().Add(time.Hour * 24 * 7).Truncate(time.Second)

	task.Deadline = theTime

	err := faker.FakeData(task)
	if err != nil {
		log.Fatalf("Couldn't create fake task, error: %v\n", err)
	}
}

func insertTask(db *sql.DB, task types.Task) int {
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
