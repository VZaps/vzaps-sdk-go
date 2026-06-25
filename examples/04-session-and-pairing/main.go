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
	status, err := client.Sessions.Status(context.Background(), instanceID, options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status: %#v\n", status)
	pairCode, err := client.Sessions.PairCode(context.Background(), instanceID, "5511999999999", options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("pair code: %#v\n", pairCode)
}
