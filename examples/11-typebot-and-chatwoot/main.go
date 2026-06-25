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
	options := vzaps.InstanceOptions{InstanceToken: os.Getenv("VZAPS_INSTANCE_TOKEN")}
	typebots, err := client.Typebots.List(context.Background(), instanceID, options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("typebots: %#v\n", typebots)
	chatwoot, err := client.Chatwoot.Get(context.Background(), instanceID, options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("chatwoot: %#v\n", chatwoot)
}
