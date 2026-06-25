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
	instanceID := os.Getenv("VZAPS_INSTANCE_ID")
	instanceToken := os.Getenv("VZAPS_INSTANCE_TOKEN")
	result, err := client.Instances.Subscribe(context.Background(), instanceID, map[string]any{"events": []string{"Message", "Connected"}}, vzaps.InstanceOptions{InstanceToken: instanceToken})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", result)
}
