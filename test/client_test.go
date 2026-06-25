package test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	vzaps "github.com/vzaps/vzaps-sdk-go/vzaps"
)

func TestAuthCachesAccessToken(t *testing.T) {
	fake := &fakeHTTP{responses: []fakeResponse{{body: `{"access_token":"jwt-token","expires_in":3600}`}}}
	client := vzaps.MustNewClient(vzaps.ClientOptions{
		ClientToken:  "client-token",
		ClientSecret: "client-secret",
		BaseURL:      "https://api.test",
		HTTPClient:   fake,
	})

	token, err := client.Auth.GetAccessToken(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if token != "jwt-token" {
		t.Fatalf("expected jwt-token, got %q", token)
	}
	if _, err := client.Auth.GetAccessToken(context.Background()); err != nil {
		t.Fatal(err)
	}
	if len(fake.requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(fake.requests))
	}
	if fake.requests[0].URL.String() != "https://api.test/token" {
		t.Fatalf("unexpected url %s", fake.requests[0].URL.String())
	}
}

func TestSendTextUsesAuthAndInstanceHeaders(t *testing.T) {
	fake := &fakeHTTP{responses: []fakeResponse{
		{body: `{"access_token":"jwt-token","expires_in":3600}`},
		{body: `{"success":true}`},
	}}
	client := vzaps.MustNewClient(vzaps.ClientOptions{
		ClientToken:  "client-token",
		ClientSecret: "client-secret",
		BaseURL:      "https://api.test",
		HTTPClient:   fake,
	})

	_, err := client.Messages.SendText(context.Background(), vzaps.MessageSendTextRequest{
		MessageSendBaseRequest: vzaps.MessageSendBaseRequest{
			InstanceScopedRequest: vzaps.InstanceScopedRequest{InstanceID: "VZ123", InstanceToken: "instance-token"},
			Phone:                 "5511999999999",
		},
		Message: "Hello",
	})
	if err != nil {
		t.Fatal(err)
	}

	req := fake.requests[1]
	if req.URL.String() != "https://api.test/instances/VZ123/chat/send/text" {
		t.Fatalf("unexpected url %s", req.URL.String())
	}
	if got := req.Header.Get("Authorization"); got != "Bearer jwt-token" {
		t.Fatalf("unexpected authorization %q", got)
	}
	if got := req.Header.Get("X-Client-Token"); got != "client-token" {
		t.Fatalf("unexpected client token %q", got)
	}
	if got := req.Header.Get("X-Instance-Token"); got != "instance-token" {
		t.Fatalf("unexpected instance token %q", got)
	}
	if string(fake.bodies[1]) != `{"phone":"5511999999999","message":"Hello"}` {
		t.Fatalf("unexpected body %s", fake.bodies[1])
	}
}

func TestInstancesListMapsPageSizeAndSearch(t *testing.T) {
	fake := &fakeHTTP{responses: []fakeResponse{
		{body: `{"access_token":"jwt-token","expires_in":3600}`},
		{body: `{"page":1,"size":10,"total":0,"content":[]}`},
	}}
	client := vzaps.MustNewClient(vzaps.ClientOptions{
		ClientToken:  "client-token",
		ClientSecret: "client-secret",
		BaseURL:      "https://api.test",
		HTTPClient:   fake,
	})

	_, err := client.Instances.List(context.Background(), vzaps.InstanceListRequest{
		Page:     1,
		PageSize: 10,
		Search:   "atendimento",
	})
	if err != nil {
		t.Fatal(err)
	}

	var body map[string]any
	if err := json.Unmarshal(fake.bodies[1], &body); err != nil {
		t.Fatal(err)
	}
	if body["size"].(float64) != 10 {
		t.Fatalf("expected size 10, got %#v", body["size"])
	}
	filter := body["filter"].(map[string]any)
	if filter["query"] != "atendimento" {
		t.Fatalf("expected filter.query atendimento, got %#v", filter["query"])
	}
}

func TestInstancesGetUsesPublicEndpoint(t *testing.T) {
	fake := &fakeHTTP{responses: []fakeResponse{
		{body: `{"access_token":"jwt-token","expires_in":3600}`},
		{body: `{"id":"VZ123"}`},
	}}
	client := vzaps.MustNewClient(vzaps.ClientOptions{
		ClientToken:  "client-token",
		ClientSecret: "client-secret",
		BaseURL:      "https://api.test",
		HTTPClient:   fake,
	})

	if _, err := client.Instances.Get(context.Background(), "VZ123"); err != nil {
		t.Fatal(err)
	}
	if fake.requests[1].URL.String() != "https://api.test/instances/get" {
		t.Fatalf("unexpected url %s", fake.requests[1].URL.String())
	}
	if fake.requests[1].Method != http.MethodPost {
		t.Fatalf("unexpected method %s", fake.requests[1].Method)
	}
	if string(fake.bodies[1]) != `{"id":"VZ123"}` {
		t.Fatalf("unexpected body %s", fake.bodies[1])
	}
}
