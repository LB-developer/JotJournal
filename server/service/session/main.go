package session

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/valkey-io/valkey-glide/go/api"
)

func NewValkeyClient() (api.GlideClientCommands, error) {
	if os.Getenv("ENV") != "prod" {
		_ = godotenv.Load()
	}
	host := os.Getenv("VALKEY_HOST")
	if host == "" {
		host = "localhost" // fallback
	}
	portStr := os.Getenv("VALKEY_PORT")
	if portStr == "" {
		portStr = "6379"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid VALKEY_PORT: %v", err)
	}

	config := api.NewGlideClientConfiguration().
		WithAddress(&api.NodeAddress{Host: host, Port: port})

	client, err := api.NewGlideClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Glide client: %w", err)
	}

	if pong, err := client.Ping(); err != nil {
		return nil, fmt.Errorf("ping to Valkey failed: %w", err)
	} else {
		fmt.Println("Valkey ping response:", pong)
	}

	return client, nil
}
