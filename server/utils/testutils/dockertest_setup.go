package testutils

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
)

func runMigrations(dbURL string) error {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for filepath.Base(wd) != "server" {
		parent := filepath.Dir(wd)
		if parent == wd {
			panic("Could not find server directory")
		}
		wd = parent
	}

	wd = "file://" + filepath.Join(wd, "db", "migrate", "migrations")
	m, err := migrate.New(
		wd,
		dbURL)
	if err != nil {
		log.Fatalf("Migrations failed %s", err)
	}
	return m.Up()
}

func SetupDockerTest() (*pgxpool.Pool, func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Couldn't construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Couldn't connect to Docker: %s", err)
	}

	resource, err := pool.Run("postgres", "15", []string{
		"POSTGRES_USER=postgres",
		"POSTGRES_PASSWORD=secret",
		"POSTGRES_DB=mydb",
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	connString := fmt.Sprintf("postgres://postgres:secret@localhost:%s/mydb?sslmode=disable", resource.GetPort("5432/tcp"))

	var dbpool *pgxpool.Pool
	if err := pool.Retry(func() error {
		var err error

		dbpool, err = pgxpool.New(context.Background(), connString)
		if err != nil {
			return err
		}
		return dbpool.Ping(context.Background())
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	if err := runMigrations(connString); err != nil {
		log.Fatalf("Failed to run migrations: %s", err)
	}

	var once sync.Once

	return dbpool, func() {
		once.Do(func() {
			if err := pool.Purge(resource); err != nil {
				log.Fatalf("Could not purge resource: %s", err)
			}
		})
	}
}
