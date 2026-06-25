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
	req := vzaps.QueueRequest{InstanceScopedRequest: vzaps.InstanceScopedRequest{InstanceID: os.Getenv("VZAPS_INSTANCE_ID"), InstanceToken: os.Getenv("VZAPS_INSTANCE_TOKEN")}}
	messages, err := client.Queues.ListMessages(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("messages: %#v\n", messages)
	operations, err := client.Queues.ListOperations(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("operations: %#v\n", operations)
}
