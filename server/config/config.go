package config

import (
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DBConfig = Config()

func Config() *pgxpool.Config {
	godotenv.Load()
	const (
		defaultMaxConns          = int32(4)
		defaultMinConns          = int32(0)
		defaultMaxConnLifetime   = time.Hour
		defaultMaxConnIdleTime   = time.Minute * 30
		defaultHealthCheckPeriod = time.Minute
		defaultConnectTimeout    = time.Second * 5
	)

	connectionURL, found := os.LookupEnv("DATABASE_URL")
	if !found {
		log.Fatalf("Unable to find database url from .env")
	}

	poolConfig, err := pgxpool.ParseConfig(connectionURL)
	if err != nil {
		log.Fatalf("Unable to parse pool config from given connection string, err:%v\n", err)
	}

	poolConfig.MaxConns = defaultMaxConns
	poolConfig.MinConns = defaultMinConns
	poolConfig.MaxConnLifetime = defaultMaxConnLifetime
	poolConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	poolConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	poolConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	poolConfig.ConnConfig.Port = 4000

	poolConfig.BeforeClose = func(c *pgx.Conn) {
		log.Printf("Closed the connection pool to the database")
	}

	return poolConfig
}
