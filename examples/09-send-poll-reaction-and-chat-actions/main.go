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
	base := vzaps.MessageSendBaseRequest{InstanceScopedRequest: vzaps.InstanceScopedRequest{InstanceID: os.Getenv("VZAPS_INSTANCE_ID"), InstanceToken: os.Getenv("VZAPS_INSTANCE_TOKEN")}, Phone: "5511999999999"}
	poll, err := client.Messages.SendPoll(context.Background(), vzaps.MessageSendPollRequest{MessageSendBaseRequest: base, Name: "Which channel do you prefer?", Options: []string{"WhatsApp", "Email", "Phone"}, SelectableOptionsCount: 1})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("poll: %#v\n", poll)
	chats, err := client.Chats.List(context.Background(), vzaps.ChatListRequest{InstanceScopedRequest: base.InstanceScopedRequest, Page: 1, PageSize: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("chats: %#v\n", chats)
}
