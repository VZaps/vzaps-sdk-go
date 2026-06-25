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
	sub, err := client.Events.Subscribe(context.Background(), vzaps.EventSubscribeRequest{InstanceID: os.Getenv("VZAPS_INSTANCE_ID"), InstanceToken: os.Getenv("VZAPS_INSTANCE_TOKEN"), Events: []vzaps.EventType{vzaps.EventMessage, vzaps.EventConnected}, Reconnect: true})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Close()
	sub.On(vzaps.EventMessage, func(event vzaps.Event) { fmt.Printf("message event: %#v\n", event) })
	sub.OnError(func(err error) { log.Println("realtime error:", err) })
	select {}
}
