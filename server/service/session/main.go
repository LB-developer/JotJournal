package session

import (
	"fmt"

	"github.com/valkey-io/valkey-glide/go/api"
)

func NewCache() (api.GlideClientCommands, error) {
	host := "localhost"
	port := 6379
	fmt.Println("START ER UP")

	config := api.NewGlideClientConfiguration().
		WithAddress(&api.NodeAddress{Host: host, Port: port})

	client, err := api.NewGlideClient(config)
	if err != nil {
		fmt.Println("There was an error: ", err)
		return nil, err
	}

	res, err := client.Ping()
	if err != nil {
		fmt.Println("There was an error: ", err)
		return nil, err
	}
	fmt.Println(res) // PONG

	return client, nil
}
