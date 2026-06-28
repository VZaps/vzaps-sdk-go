package vzaps

import (
	"github.com/vzaps/vzaps-sdk-go/events"
	"github.com/vzaps/vzaps-sdk-go/resources"
)

type InstanceOptions = resources.InstanceOptions
type AuthTokenResponse = resources.AuthTokenResponse
type InstanceCreateRequest = resources.InstanceCreateRequest
type InstanceListRequest = resources.InstanceListRequest
type InstanceScopedRequest = resources.InstanceScopedRequest
type MessageSendBaseRequest = resources.MessageSendBaseRequest
type MessageSendTextRequest = resources.MessageSendTextRequest
type MessageSendImageRequest = resources.MessageSendImageRequest
type MessageSendAudioRequest = resources.MessageSendAudioRequest
type MessageSendDocumentRequest = resources.MessageSendDocumentRequest
type MessageSendVideoRequest = resources.MessageSendVideoRequest
type MessageSendStickerRequest = resources.MessageSendStickerRequest
type MessageSendGifRequest = resources.MessageSendGifRequest
type MessageSendLocationRequest = resources.MessageSendLocationRequest
type MessageSendContactRequest = resources.MessageSendContactRequest
type MessageButton = resources.MessageButton
type MessageSendButtonsRequest = resources.MessageSendButtonsRequest
type MessageListRow = resources.MessageListRow
type MessageListSection = resources.MessageListSection
type MessageSendListRequest = resources.MessageSendListRequest
type MessageSendLinkRequest = resources.MessageSendLinkRequest
type MessageSendPollRequest = resources.MessageSendPollRequest
type MessagePollVoteRequest = resources.MessagePollVoteRequest
type MessageReactRequest = resources.MessageReactRequest
type MessageReactRemoveRequest = resources.MessageReactRemoveRequest
type MessagePresenceRequest = resources.MessagePresenceRequest
type MessageMarkReadRequest = resources.MessageMarkReadRequest
type MessageDownloadRequest = resources.MessageDownloadRequest
type MessageEditRequest = resources.MessageEditRequest
type MessageDeleteRequest = resources.MessageDeleteRequest
type WebhookConfigRequest = resources.WebhookConfigRequest
type WebhookLogSearchRequest = resources.WebhookLogSearchRequest
type WebhookLogRequest = resources.WebhookLogRequest
type ContactAddRequest = resources.ContactAddRequest
type UserPhonesRequest = resources.UserPhonesRequest
type UserAvatarRequest = resources.UserAvatarRequest
type GroupListRequest = resources.GroupListRequest
type GroupInfoRequest = resources.GroupInfoRequest
type GroupInviteLinkRequest = resources.GroupInviteLinkRequest
type GroupMutationRequest = resources.GroupMutationRequest
type QueueRequest = resources.QueueRequest
type QueueMessageRequest = resources.QueueMessageRequest
type TypebotRequest = resources.TypebotRequest
type TypebotMutationRequest = resources.TypebotMutationRequest
type TypebotSessionRequest = resources.TypebotSessionRequest
type TypebotStartSessionRequest = resources.TypebotStartSessionRequest
type ChatwootRequest = resources.ChatwootRequest
type ChatwootImportRequest = resources.ChatwootImportRequest
type ChatRequest = resources.ChatRequest
type ChatListRequest = resources.ChatListRequest
type ChatDeleteRequest = resources.ChatDeleteRequest
type ChatMuteRequest = resources.ChatMuteRequest
type ChatClearRequest = resources.ChatClearRequest
type ChatExpirationRequest = resources.ChatExpirationRequest
type SessionBusinessCategory = resources.SessionBusinessCategory
type SessionBusinessProfile = resources.SessionBusinessProfile
type SessionStatusData = resources.SessionStatusData
type SessionStatusResponse = resources.SessionStatusResponse

type EventType = events.EventType
type Event = events.Event
type EventSubscribeRequest = events.SubscribeRequest
type EventHandler = events.EventHandler
type ErrorHandler = events.ErrorHandler
type LifecycleHandler = events.LifecycleHandler
type EventSubscription = events.Subscription
type WebSocketConn = events.WebSocketConn
type WebSocketDialer = events.WebSocketDialer

const (
	EventMessage                 = events.EventMessage
	EventReadReceipt             = events.EventReadReceipt
	EventPresence                = events.EventPresence
	EventHistorySync             = events.EventHistorySync
	EventChatPresence            = events.EventChatPresence
	EventConnected               = events.EventConnected
	EventDisconnected            = events.EventDisconnected
	EventGroupParticipantsAdd    = events.EventGroupParticipantsAdd
	EventGroupParticipantsRemove = events.EventGroupParticipantsRemove
	EventAll                     = events.EventAll
)
