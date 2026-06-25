package events

import "time"

type EventType string

const (
	EventMessage                 EventType = "Message"
	EventReadReceipt             EventType = "ReadReceipt"
	EventPresence                EventType = "Presence"
	EventHistorySync             EventType = "HistorySync"
	EventChatPresence            EventType = "ChatPresence"
	EventConnected               EventType = "Connected"
	EventDisconnected            EventType = "Disconnected"
	EventGroupParticipantsAdd    EventType = "GroupParticipantsAdd"
	EventGroupParticipantsRemove EventType = "GroupParticipantsRemove"
	EventAll                     EventType = "All"
)

type Event struct {
	ID         string         `json:"id"`
	Type       EventType      `json:"type"`
	InstanceID string         `json:"instance_id"`
	CreatedAt  string         `json:"created_at"`
	Data       map[string]any `json:"data"`
}

type SubscribeRequest struct {
	InstanceID    string
	InstanceToken string
	Events        []EventType
	Reconnect     bool
	MaxRetries    int
	RetryDelay    time.Duration
	LastEventID   string
}
