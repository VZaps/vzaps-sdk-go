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
	result, err := client.Messages.SendText(context.Background(), vzaps.MessageSendTextRequest{
		MessageSendBaseRequest: vzaps.MessageSendBaseRequest{InstanceScopedRequest: vzaps.InstanceScopedRequest{InstanceID: os.Getenv("VZAPS_INSTANCE_ID"), InstanceToken: os.Getenv("VZAPS_INSTANCE_TOKEN")}, Phone: "5511999999999"},
		Message:                "Hello from VZaps Go SDK",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", result)
}
