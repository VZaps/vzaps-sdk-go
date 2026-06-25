package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vzaps "github.com/vzaps/vzaps-sdk-go/vzaps"
)

func main() {
	client := vzaps.MustNewClient(vzaps.ClientOptions{ClientToken: os.Getenv("VZAPS_CLIENT_TOKEN"), ClientSecret: os.Getenv("VZAPS_CLIENT_SECRET")})
	instance, err := client.Instances.Create(context.Background(), vzaps.InstanceCreateRequest{Name: "Support", Webhook: "https://example.com/webhook", EventsSubscribe: []string{"Message", "Connected"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", instance)
}
