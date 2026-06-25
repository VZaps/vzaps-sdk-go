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
	image, err := client.Messages.SendImage(context.Background(), vzaps.MessageSendImageRequest{MessageSendBaseRequest: base, Image: "https://picsum.photos/800/600.jpg", Caption: "Image from VZaps Go SDK"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("image: %#v\n", image)
	buttons, err := client.Messages.SendButtons(context.Background(), vzaps.MessageSendButtonsRequest{MessageSendBaseRequest: base, Message: "Choose an option", Footer: "VZaps SDK", Buttons: []vzaps.MessageButton{{ID: "sales", Text: "Sales"}, {ID: "support", Text: "Support"}}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("buttons: %#v\n", buttons)
}
