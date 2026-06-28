# VZaps Go SDK Examples

Runnable programs that consume the published Go module (`github.com/vzaps/vzaps-sdk-go/vzaps`).

Each numbered folder is a standalone module with its own `go.mod`. You do **not** need to clone the full SDK repository to run one example.

## Prerequisites

- Go 1.22 or later

## Option A — one example folder (recommended)

Download only one example folder, for example [`07-send-text-message`](https://github.com/VZaps/vzaps-sdk-go/tree/main/examples/07-send-text-message):

1. Open the folder on GitHub and choose **Download ZIP**, or run:

```bash
npx --yes degit VZaps/vzaps-sdk-go/examples/07-send-text-message vzaps-go-send-text
cd vzaps-go-send-text
```

2. Set credentials:

```bash
export VZAPS_CLIENT_TOKEN=your-client-token
export VZAPS_CLIENT_SECRET=your-client-secret
export VZAPS_INSTANCE_ID=VZ...
export VZAPS_INSTANCE_TOKEN=your-instance-token
```

3. Run:

```bash
go run .
```

When developing against a local SDK checkout, add a replace directive to `go.mod`:

```go
replace github.com/vzaps/vzaps-sdk-go/vzaps => /path/to/vzaps-sdk-go
```

## Option B — sparse checkout

```bash
git clone --depth 1 --filter=blob:none --sparse https://github.com/VZaps/vzaps-sdk-go.git
cd vzaps-sdk-go
git sparse-checkout set examples/07-send-text-message
cd examples/07-send-text-message
go run .
```

## Option C — full repository clone

```bash
git clone https://github.com/VZaps/vzaps-sdk-go.git
cd vzaps-sdk-go/examples/07-send-text-message
go run .
```

## Examples

| Folder | Topic |
| --- | --- |
| `01-auth-and-list-instances` | Auth and instance listing |
| `02-create-instance` | Create instance |
| `03-instance-subscription` | Billing subscription |
| `04-session-and-pairing` | Session status, QR, and pairing code |
| `05-configure-webhook` | Webhook configuration |
| `06-realtime-subscribe` | Realtime WebSocket subscription |
| `07-send-text-message` | Send text message |
| `08-send-media-and-interactive` | Media, buttons, and list |
| `09-send-poll-reaction-and-chat-actions` | Poll, reaction, and chat actions |
| `10-queues` | Message and operation queues |
| `11-typebot-and-chatwoot` | TypeBot and Chatwoot |

## Coverage

- Auth and instance listing
- Instance creation and billing subscription checkout
- Session status, QR, and phone pairing code
- Webhook and realtime subscription
- Text, media, buttons, list, poll, reactions, presence
- Queue list/remove/purge examples
- TypeBot and Chatwoot integration examples
