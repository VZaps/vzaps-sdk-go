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
	result, err := client.Webhooks.Set(context.Background(), vzaps.WebhookConfigRequest{
		InstanceScopedRequest: vzaps.InstanceScopedRequest{InstanceID: os.Getenv("VZAPS_INSTANCE_ID"), InstanceToken: os.Getenv("VZAPS_INSTANCE_TOKEN")},
		WebhookURL:            "https://example.com/webhook",
		Events:                []string{"Message", "Connected", "Disconnected"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", result)
}
