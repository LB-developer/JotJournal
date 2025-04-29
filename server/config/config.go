package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Config struct {
	DBURL                  string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = InitConfig()

func InitConfig() *Config {
	if os.Getenv("ENV") != "prod" {
		_ = godotenv.Load()
	}
	return &Config{
		DBURL:                  getEnv("DATABASE_URL", "postgresql://localhost:5432/defaultdb"),
		JWTSecret:              getEnv("JWT_SECRET", "oh-no-we-are-exposed-please-dont-be-nefarious"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION", 3600*24*7),
	}
}

var DBConfig = InitDBConfig()

func InitDBConfig() *pgxpool.Config {
	if os.Getenv("ENV") != "prod" {
		_ = godotenv.Load()
	}
	const (
		defaultMaxConns          = int32(4)
		defaultMinConns          = int32(0)
		defaultMaxConnLifetime   = time.Hour
		defaultMaxConnIdleTime   = time.Minute * 30
		defaultHealthCheckPeriod = time.Minute
		defaultConnectTimeout    = time.Second * 5
	)

	connectionURL := Envs.DBURL

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

	// poolConfig.ConnConfig.Port = 4000

	poolConfig.BeforeClose = func(c *pgx.Conn) {
		log.Printf("Closed the connection pool to the database")
	}

	return poolConfig
}

func getEnv(key, fallback string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return fallback
	}

	return value
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, found := os.LookupEnv(key); found {
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return v
	}

	return fallback
}
