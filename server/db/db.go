package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lb-developer/jotjournal/config"
)

func NewPgxPool() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.NewWithConfig(ctx, config.DBConfig)
	if err != nil {
		log.Println("Unable to create connection pool")
		return nil, err
	}

	connection, err := dbpool.Acquire(ctx)
	if err != nil {
		log.Println("Unable to acquire connection from the database pool")
		return nil, err
	}

	if err := connection.Ping(ctx); err != nil {
		log.Println("Unable to ping database")
		return nil, err
	}

	return dbpool, nil
}
