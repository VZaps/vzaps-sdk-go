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
	token, err := client.Auth.GetAccessToken(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JWT prefix:", token[:20]+"...")
	instances, err := client.Instances.List(context.Background(), vzaps.InstanceListRequest{Page: 1, PageSize: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", instances)
}
