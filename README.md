# VZaps Go SDK

[![CI](https://github.com/VZaps/vzaps-sdk-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/VZaps/vzaps-sdk-go/actions/workflows/ci.yml) [![SDK Documentation](https://img.shields.io/badge/SDK-Documentation-blue)](https://docs.vzaps.com/en/sdk/go/installation) [![license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/vzaps/vzaps-sdk-go/vzaps.svg)](https://pkg.go.dev/github.com/vzaps/vzaps-sdk-go/vzaps)

Official Go client for the [VZaps public API](https://docs.vzaps.com). Send WhatsApp messages, manage instances, configure webhooks, and subscribe to realtime events with a resource-oriented, context-aware interface.

Works in **Go 1.22+**. HTTP uses the standard library; WebSocket realtime uses [`gorilla/websocket`](https://github.com/gorilla/websocket).

---

## Table of contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Quick start](#quick-start)
- [Authentication](#authentication)
- [Configuration](#configuration)
- [Resources](#resources)
- [Instance tokens](#instance-tokens)
- [Webhooks](#webhooks)
- [Realtime events](#realtime-events)
- [Error handling](#error-handling)
- [Go](#go)
- [Documentation](#documentation)

---

## Features

- **Automatic JWT handling** — exchanges `ClientToken` + `ClientSecret` for a bearer token and refreshes it before expiry.
- **Resource-oriented API** — `Instances`, `Messages`, `Webhooks`, `Contacts`, `Groups`, and `Events` mirror the public HTTP contract.
- **Realtime WebSocket client** — subscribe to instance events with reconnect, resume (`LastEventID`), and server-side ack.
- **Instance token support** — pass `InstanceToken` on each instance-scoped request.
- **Structured request types** — exported structs with JSON tags for the public API wire format.
- **Extensible transport** — inject a custom `HTTPClient` or WebSocket `Dialer` for tests or custom runtimes.

---

## Requirements

| Runtime | Minimum version |
| --- | --- |
| Go | 1.22+ |

The SDK uses `net/http` by default. No extra HTTP dependency is required.

---

## Installation

```bash
go get github.com/vzaps/vzaps-sdk-go/vzaps
```

---

## Quick start

Create credentials in the [VZaps dashboard](https://docs.vzaps.com) (`clientToken` and `clientSecret`), then send a text message:

```go
import (
	"context"
	"os"

	vzaps "github.com/vzaps/vzaps-sdk-go/vzaps"
)

client := vzaps.MustNewClient(vzaps.ClientOptions{
	ClientToken:  os.Getenv("VZAPS_CLIENT_TOKEN"),
	ClientSecret: os.Getenv("VZAPS_CLIENT_SECRET"),
})

_, err := client.Messages.SendText(context.Background(), vzaps.MessageSendTextRequest{
	MessageSendBaseRequest: vzaps.MessageSendBaseRequest{
		InstanceScopedRequest: vzaps.InstanceScopedRequest{
			InstanceID:    "VZKB8AU4S4CWY1SLXX4I5WJGRZQMDDFTV6",
			InstanceToken: os.Getenv("VZAPS_INSTANCE_TOKEN"),
		},
		Phone: "5511999999999",
	},
	Message: "Hello from VZaps",
})
```

---

## Authentication

VZaps uses a two-step model:

1. **Account credentials** — `ClientToken` and `ClientSecret` identify your integration. The SDK calls `POST /token` and caches the JWT.
2. **Instance token** — instance-scoped routes also require `X-Instance-Token`. Pass it on each instance-scoped request (see [Instance tokens](#instance-tokens)).

Every authenticated HTTP request sends:

| Header | Value |
| --- | --- |
| `Authorization` | `Bearer <jwt>` |
| `X-Client-Token` | Your client token |
| `X-Instance-Token` | Instance token, on instance-scoped requests |

You rarely need to call `Auth.GetAccessToken` directly — resources attach the token for you. Use it when integrating with custom HTTP logic:

```go
token, err := client.Auth.GetAccessToken(ctx)
```

---

## Configuration

The SDK connects to the VZaps production platform automatically:

| Service | Endpoint |
| --- | --- |
| REST API | `https://api.vzaps.com` |
| Realtime WebSocket | `wss://realtime.vzaps.com/events/ws` |

Pass options to `NewClient` or `MustNewClient`:

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `ClientToken` | `string` | — | **Required.** Public client token from the dashboard. |
| `ClientSecret` | `string` | — | **Required.** Client secret used to obtain JWTs. |
| `Timeout` | `time.Duration` | `30s` | HTTP request timeout. |
| `TokenSkew` | `time.Duration` | `1m` | Refresh JWT this long before expiry. |
| `HTTPClient` | `HTTPDoer` | `http.DefaultClient` | Custom HTTP client (tests, proxies, tracing). |
| `Dialer` | `WebSocketDialer` | gorilla default | Custom WebSocket dialer. |
| `UserAgent` | `string` | — | Optional `User-Agent` header on HTTP requests. |

No host configuration is required — install the module, pass your credentials, and the client targets the production API and realtime service.

---

## Resources

The client exposes namespaced resources. Generic response types let you align with your own structs or the [OpenAPI schema](https://docs.vzaps.com/api-reference).

All resource methods accept a `context.Context` as the first argument.

### `client.Instances`

| Method | HTTP | Description |
| --- | --- | --- |
| `Create(ctx, data)` | `PUT /instances/create` | Create a WhatsApp instance. |
| `List(ctx, data?)` | `POST /instances/list` | List instances (pagination, search, sort). |
| `Get(ctx, instanceID)` | `GET /instances/:id` | Get instance details. |
| `Update(ctx, instanceID, data, options?)` | `PATCH /instances/:id` | Update instance settings. |
| `Restart(ctx, instanceID, options?)` | `POST /instances/:id/restart` | Restart instance runtime. |

### `client.Messages`

`client.Messages` wraps the public WhatsApp send and chat endpoints. The most common calls are shown below; the SDK also exposes the other public message operations documented in the API reference, including media, interactive messages, reactions, polls, downloads, edits, deletes, presence, and read receipts.

```go
_, err := client.Messages.SendText(ctx, vzaps.MessageSendTextRequest{
	MessageSendBaseRequest: vzaps.MessageSendBaseRequest{
		InstanceScopedRequest: vzaps.InstanceScopedRequest{
			InstanceID:    "VZ...",
			InstanceToken: "instance-token",
		},
		Phone: "5511999999999",
	},
	Message: "Hello",
})

_, err = client.Messages.SendImage(ctx, vzaps.MessageSendImageRequest{
	MessageSendBaseRequest: base,
	Image:   "https://example.com/photo.jpg",
	Caption: "Check this out",
})
```

Available send helpers include `SendText`, `SendImage`, `SendAudio`, `SendDocument`, `SendVideo`, `SendSticker`, `SendGif`, `SendLocation`, `SendContact`, `SendButtons`, `SendList`, `SendLink`, and `SendPoll`. See the API documentation for complete payload examples.

### `client.Webhooks`

| Method | HTTP | Description |
| --- | --- | --- |
| `Get(ctx, instanceID, options?)` | `GET /instances/:id/webhook` | Read current webhook configuration. |
| `Set(ctx, request)` | `POST /instances/:id/webhook` | Configure webhook URL and subscribed events. |

### `client.Contacts`

| Method | HTTP | Description |
| --- | --- | --- |
| `List(ctx, instanceID, options?)` | `GET /instances/:id/contact/list` | List contacts for the instance. |
| `Add(ctx, request)` | `POST /instances/:id/contact/add` | Add a contact. |

### `client.Groups`

| Method | HTTP | Description |
| --- | --- | --- |
| `List(ctx, request)` | `GET /instances/:id/group/list` | List groups (paginated). |
| `Get(ctx, request)` | `GET /instances/:id/group/info` | Get group metadata by `GroupID`. |

### `client.Sessions`

| Method | HTTP | Description |
| --- | --- | --- |
| `Status(ctx, instanceID, options)` | `GET /instances/:id/session/status` | Check WhatsApp login state and, when connected, live profile fields. |

`GET /instances/{id}/session/status` returns `SessionStatusResponse`. When `Data.Connected` is `true`, `Data` includes (in order) `Phone`, `WhatsAppJID`, `PushName`, `BusinessName`, `BusinessProfile`, `ProfilePictureID`, `ProfilePictureURL`, `ProfileURL`, and optional `VerifiedName`, `About`, `Website`. When disconnected, `Data` only has `Connected: false`.

Other public namespaces are available as first-class resources too: `Sessions`, `Users`, `Queues`, `Typebots`, `Chatwoot`, and `Chats`.

### `client.Request(ctx, method, path, options?, &out)`

Escape hatch for advanced calls or newly released endpoints:

```go
var instance map[string]any
err := client.Request(ctx, http.MethodPost, "/instances/get", vzaps.RequestOptions{
	Body: map[string]string{"id": "VZ..."},
}, &instance)
```

---

## Instance tokens

Instance-scoped routes require the instance token in addition to account credentials. Pass it on each request that targets an instance:

```go
_, err := client.Messages.SendText(ctx, vzaps.MessageSendTextRequest{
	MessageSendBaseRequest: vzaps.MessageSendBaseRequest{
		InstanceScopedRequest: vzaps.InstanceScopedRequest{
			InstanceID:    "VZ...",
			InstanceToken: "instance-token",
		},
		Phone: "5511999999999",
	},
	Message: "Hello",
})
```

---

## Webhooks

Configure HTTP callbacks for instance events (same payload shape as realtime `data`, delivered to your URL):

```go
_, err := client.Webhooks.Set(ctx, vzaps.WebhookConfigRequest{
	InstanceScopedRequest: vzaps.InstanceScopedRequest{
		InstanceID:    "VZ...",
		InstanceToken: "instance-token",
	},
	WebhookURL: "https://example.com/webhooks/vzaps",
	Events:     []string{"Message", "Connected", "Disconnected"},
})
```

Common event types: `Message`, `ReadReceipt`, `Connected`, `Disconnected`, `Presence`, `ChatPresence`, `HistorySync`, `GroupParticipantsAdd`, `GroupParticipantsRemove`, or `All`.

Event payloads (webhook and realtime) use **snake_case**, matching the platform. Incoming media events include `media_url` inside `data` when platform storage is available.

---

## Realtime events

Subscribe to the same events over WebSocket at **`wss://realtime.vzaps.com`**. This is the recommended path for in-app notifications, bots, and dashboards that need low-latency delivery without exposing a public webhook URL.

### Subscribe

```go
sub, err := client.Events.Subscribe(ctx, vzaps.EventSubscribeRequest{
	InstanceID:    "VZ...",
	InstanceToken: "instance-token",
	Events:        []vzaps.EventType{vzaps.EventMessage, vzaps.EventConnected, vzaps.EventDisconnected},
	Reconnect:     true,
	LastEventID:   "evt_previous_id", // optional resume after disconnect
})

sub.OnOpen(func() {
	log.Println("Connected to realtime")
})

sub.On(vzaps.EventMessage, func(event vzaps.Event) {
	log.Println(event.Data)
})

sub.OnError(func(err error) {
	log.Println(err)
})

// Graceful shutdown
sub.Close()
```

### Event envelope

Each WebSocket message keeps the platform shape (`snake_case`):

```json
{
  "id": "evt_…",
  "type": "Message",
  "instance_id": "VZ…",
  "created_at": "2026-06-23T22:57:17.000Z",
  "data": {
    "type": "Message",
    "event": { },
    "media_url": "https://…"
  }
}
```

- **`data`** — same payload as webhook delivery (`snake_case`).
- **`media_url`** — present on incoming media messages when platform storage is available.

### Delivery and ack

Delivery is **at-least-once**. After your handler runs, the SDK sends an ack automatically on the WebSocket connection. Use `LastEventID` when reconnecting if you need to reduce gaps. Deduplicate on `event.ID` in your application if you process events idempotently.

### Subscribe options

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `InstanceID` | `string` | — | **Required.** Instance to watch. |
| `Events` | `[]EventType` | all subscribed | Comma-filtered event types. |
| `InstanceToken` | `string` | — | **Required.** Instance token for authorization. |
| `Reconnect` | `bool` | `true` | Reconnect after socket close. |
| `MaxRetries` | `int` | unlimited | Max reconnect attempts. |
| `RetryDelay` | `time.Duration` | exponential backoff | Delay between reconnects. |
| `LastEventID` | `string` | — | Resume cursor after disconnect. |

### Handler registration

| Event name | When it fires |
| --- | --- |
| `OnOpen` | WebSocket connected. |
| `OnClose` | WebSocket closed. |
| `OnError` | Handler or transport error. |
| `On(EventMessage, …)`, `On(EventConnected, …)` | Matching realtime event type. |
| `On(EventAll, …)` | Every event type. |

---

## Error handling

The SDK returns typed errors you can inspect with `errors.As`:

| Type | When |
| --- | --- |
| `*vzaps.Error` | Base type; HTTP errors include `Status`, `Code`, and `Details`. |
| `*vzaps.AuthenticationError` | Invalid `ClientToken` / `ClientSecret` (401). |
| `*vzaps.TimeoutError` | Request exceeded `Timeout`. |

```go
import (
	"errors"
	"fmt"

	vzaps "github.com/vzaps/vzaps-sdk-go/vzaps"
)

_, err := client.Messages.SendText(ctx, req)
if err != nil {
	var authErr *vzaps.AuthenticationError
	var timeoutErr *vzaps.TimeoutError
	var apiErr *vzaps.Error

	switch {
	case errors.As(err, &authErr):
		fmt.Println("Check client credentials")
	case errors.As(err, &timeoutErr):
		fmt.Println("Request timed out")
	case errors.As(err, &apiErr):
		fmt.Println(apiErr.Status, apiErr.Message, apiErr.Details)
	}
	return err
}
```

---

## Go

The package uses **PascalCase** for exported struct fields and **snake_case JSON tags** on HTTP request types and API responses. **Realtime and webhook event payloads stay in snake_case** so both delivery channels match.

Exported types include options, events, and requests:

```go
import vzaps "github.com/vzaps/vzaps-sdk-go/vzaps"

var _ vzaps.ClientOptions
var _ vzaps.Event
var _ vzaps.EventType
var _ vzaps.MessageSendTextRequest
var _ vzaps.WebhookConfigRequest
var _ vzaps.EventSubscribeRequest
```

Decode responses into your own struct when you want strongly typed API responses:

```go
type InstanceListResponse struct {
	Content []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"content"`
}

var page InstanceListResponse
err := client.Request(ctx, http.MethodPost, "/instances/list", vzaps.RequestOptions{
	Body: map[string]any{"page": 1, "size": 20},
}, &page)
```

---

## Documentation

- [VZaps docs](https://docs.vzaps.com)
- [API reference (OpenAPI)](https://docs.vzaps.com/api-reference)
- [Postman collections](https://docs.vzaps.com/postman/)
- [Report an issue](https://github.com/vzaps/vzaps-sdk-go/issues)

---

## License

MIT © VZaps
