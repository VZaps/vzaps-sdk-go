package resources

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/vzaps/vzaps-sdk-go/internal/transport"
)

type AuthResource struct{ http *transport.Client }

func NewAuth(httpClient *transport.Client) *AuthResource {
	return &AuthResource{http: httpClient}
}

func (r *AuthResource) GetAccessToken(ctx context.Context) (string, error) {
	return r.http.GetAccessToken(ctx)
}

type InstancesResource struct{ http *transport.Client }

func NewInstances(httpClient *transport.Client) *InstancesResource {
	return &InstancesResource{http: httpClient}
}

func (r *InstancesResource) Create(ctx context.Context, req InstanceCreateRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPut, "/instances/create", req, "")
}

func (r *InstancesResource) List(ctx context.Context, req InstanceListRequest) (map[string]any, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = req.PageSize
	}
	if req.Size == 0 {
		req.Size = 20
	}
	if req.Filter == nil {
		req.Filter = map[string]any{}
	}
	if strings.TrimSpace(req.Search) != "" {
		req.Filter["query"] = strings.TrimSpace(req.Search)
	}
	return requestMap(ctx, r.http, http.MethodPost, "/instances/list", req, "")
}

func (r *InstancesResource) Get(ctx context.Context, instanceID string) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/get", map[string]string{"id": instanceID}, "")
}

func (r *InstancesResource) Update(ctx context.Context, instanceID string, body any, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPatch, "/instances/"+url.PathEscape(instanceID), body, options.InstanceToken)
}

func (r *InstancesResource) Restart(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/"+url.PathEscape(instanceID)+"/restart", nil, options.InstanceToken)
}

func (r *InstancesResource) Delete(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodDelete, "/instances/"+url.PathEscape(instanceID), nil, options.InstanceToken)
}

func (r *InstancesResource) Provision(ctx context.Context, req InstanceCreateRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPut, "/instances/provision", req, "")
}

func (r *InstancesResource) Search(ctx context.Context, body any) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/search", body, "")
}

func (r *InstancesResource) Subscribe(ctx context.Context, instanceID string, body any, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/"+url.PathEscape(instanceID)+"/subscribe", body, options.InstanceToken)
}

func (r *InstancesResource) ResumeSubscription(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/"+url.PathEscape(instanceID)+"/resume-subscription", nil, options.InstanceToken)
}

func (r *InstancesResource) Cancel(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPut, "/instances/"+url.PathEscape(instanceID)+"/cancel", nil, options.InstanceToken)
}

type MessagesResource struct{ http *transport.Client }

func NewMessages(httpClient *transport.Client) *MessagesResource {
	return &MessagesResource{http: httpClient}
}

func (r *MessagesResource) SendText(ctx context.Context, req MessageSendTextRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/text", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendImage(ctx context.Context, req MessageSendImageRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/image", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendAudio(ctx context.Context, req MessageSendAudioRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/audio", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendDocument(ctx context.Context, req MessageSendDocumentRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/document", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendVideo(ctx context.Context, req MessageSendVideoRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/video", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendSticker(ctx context.Context, req MessageSendStickerRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/sticker", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendGif(ctx context.Context, req MessageSendGifRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/gif", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendLocation(ctx context.Context, req MessageSendLocationRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/location", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendContact(ctx context.Context, req MessageSendContactRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/contact", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendButtons(ctx context.Context, req MessageSendButtonsRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/buttons", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendList(ctx context.Context, req MessageSendListRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/list", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendLink(ctx context.Context, req MessageSendLinkRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/link", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) SendPoll(ctx context.Context, req MessageSendPollRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/send/poll", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) PollVote(ctx context.Context, req MessagePollVoteRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/poll/vote", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) React(ctx context.Context, req MessageReactRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/react", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) RemoveReaction(ctx context.Context, req MessageReactRemoveRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodDelete, "/chat/react", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) Presence(ctx context.Context, req MessagePresenceRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/presence", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) MarkRead(ctx context.Context, req MessageMarkReadRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/markread", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) DownloadImage(ctx context.Context, req MessageDownloadRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/downloadimage", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) DownloadVideo(ctx context.Context, req MessageDownloadRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/downloadvideo", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) DownloadAudio(ctx context.Context, req MessageDownloadRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/downloadaudio", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) DownloadDocument(ctx context.Context, req MessageDownloadRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPost, "/chat/downloaddocument", req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) Edit(ctx context.Context, req MessageEditRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodPatch, "/chat/messages/"+url.PathEscape(req.MessageID), req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) Delete(ctx context.Context, req MessageDeleteRequest) (map[string]any, error) {
	return r.send(ctx, http.MethodDelete, "/chat/messages/"+url.PathEscape(req.MessageID), req.InstanceID, req.InstanceToken, req)
}

func (r *MessagesResource) Send(ctx context.Context, instanceID, path string, body any, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/"+url.PathEscape(instanceID)+"/chat/"+strings.TrimLeft(path, "/"), body, options.InstanceToken)
}

func (r *MessagesResource) send(ctx context.Context, method, path, instanceID, instanceToken string, body any) (map[string]any, error) {
	return requestMap(ctx, r.http, method, "/instances/"+url.PathEscape(instanceID)+path, body, instanceToken)
}

type WebhooksResource struct{ http *transport.Client }

func NewWebhooks(httpClient *transport.Client) *WebhooksResource {
	return &WebhooksResource{http: httpClient}
}

func (r *WebhooksResource) Get(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(instanceID)+"/webhook", nil, options.InstanceToken)
}

func (r *WebhooksResource) Set(ctx context.Context, req WebhookConfigRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/"+url.PathEscape(req.InstanceID)+"/webhook", req, req.InstanceToken)
}

func (r *WebhooksResource) SearchLogs(ctx context.Context, req WebhookLogSearchRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/"+url.PathEscape(req.InstanceID)+"/webhook/logs/search", req, req.InstanceToken)
}

func (r *WebhooksResource) GetLog(ctx context.Context, req WebhookLogRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(req.InstanceID)+"/webhook/logs/"+url.PathEscape(req.LogID), nil, req.InstanceToken)
}

func (r *WebhooksResource) RetryLog(ctx context.Context, req WebhookLogRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/"+url.PathEscape(req.InstanceID)+"/webhook/logs/"+url.PathEscape(req.LogID)+"/retry", nil, req.InstanceToken)
}

type ContactsResource struct{ http *transport.Client }

func NewContacts(httpClient *transport.Client) *ContactsResource {
	return &ContactsResource{http: httpClient}
}

func (r *ContactsResource) List(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(instanceID)+"/contact/list", nil, options.InstanceToken)
}

func (r *ContactsResource) Add(ctx context.Context, req ContactAddRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/"+url.PathEscape(req.InstanceID)+"/contact/add", req, req.InstanceToken)
}

type GroupsResource struct{ http *transport.Client }

func NewGroups(httpClient *transport.Client) *GroupsResource {
	return &GroupsResource{http: httpClient}
}

func (r *GroupsResource) List(ctx context.Context, req GroupListRequest) (map[string]any, error) {
	return requestMapQuery(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(req.InstanceID)+"/group/list", nil, req.InstanceToken, map[string]any{"page": req.Page, "pageSize": req.PageSize})
}

func (r *GroupsResource) Get(ctx context.Context, req GroupInfoRequest) (map[string]any, error) {
	return requestMapQuery(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(req.InstanceID)+"/group/info", nil, req.InstanceToken, map[string]any{"groupId": req.GroupID})
}

func (r *GroupsResource) InviteLink(ctx context.Context, req GroupInviteLinkRequest) (map[string]any, error) {
	return requestMapQuery(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(req.InstanceID)+"/group/invitelink", nil, req.InstanceToken, map[string]any{"groupId": req.GroupID, "reset": req.Reset})
}

func (r *GroupsResource) SetPhoto(ctx context.Context, req GroupMutationRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/group/photo", req, req.InstanceToken)
}

func (r *GroupsResource) SetName(ctx context.Context, req GroupMutationRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/group/name", req, req.InstanceToken)
}

func (r *GroupsResource) SetDescription(ctx context.Context, req GroupMutationRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/group/description", req, req.InstanceToken)
}

func (r *GroupsResource) SetSettings(ctx context.Context, req GroupMutationRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/group/settings", req, req.InstanceToken)
}

func (r *GroupsResource) Create(ctx context.Context, req GroupMutationRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/group/create", req, req.InstanceToken)
}

func (r *GroupsResource) AddAdmin(ctx context.Context, req GroupMutationRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/group/add-admin", req, req.InstanceToken)
}

func (r *GroupsResource) RemoveAdmin(ctx context.Context, req GroupMutationRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/group/remove-admin", req, req.InstanceToken)
}

type SessionsResource struct{ http *transport.Client }

func NewSessions(httpClient *transport.Client) *SessionsResource {
	return &SessionsResource{http: httpClient}
}

func (r *SessionsResource) Status(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(instanceID)+"/session/status", nil, options.InstanceToken)
}

func (r *SessionsResource) QR(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(instanceID)+"/session/qr", nil, options.InstanceToken)
}

func (r *SessionsResource) PairCode(ctx context.Context, instanceID, phone string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(instanceID)+"/session/paircode/"+url.PathEscape(phone), nil, options.InstanceToken)
}

func (r *SessionsResource) Disconnect(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPost, "/instances/"+url.PathEscape(instanceID)+"/session/disconnect", nil, options.InstanceToken)
}

type UsersResource struct{ http *transport.Client }

func NewUsers(httpClient *transport.Client) *UsersResource {
	return &UsersResource{http: httpClient}
}

func (r *UsersResource) Info(ctx context.Context, req UserPhonesRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/user/info", req, req.InstanceToken)
}

func (r *UsersResource) Check(ctx context.Context, req UserPhonesRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/user/check", req, req.InstanceToken)
}

func (r *UsersResource) Avatar(ctx context.Context, req UserAvatarRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/user/avatar", req, req.InstanceToken)
}

func (r *UsersResource) Contacts(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(instanceID)+"/user/contacts", nil, options.InstanceToken)
}

type QueuesResource struct{ http *transport.Client }

func NewQueues(httpClient *transport.Client) *QueuesResource {
	return &QueuesResource{http: httpClient}
}

func (r *QueuesResource) ListMessages(ctx context.Context, req QueueRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(req.InstanceID)+"/queue/messages", nil, req.InstanceToken)
}

func (r *QueuesResource) PurgeMessages(ctx context.Context, req QueueRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodDelete, "/instances/"+url.PathEscape(req.InstanceID)+"/queue/messages", nil, req.InstanceToken)
}

func (r *QueuesResource) RemoveMessage(ctx context.Context, req QueueMessageRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodDelete, "/instances/"+url.PathEscape(req.InstanceID)+"/queue/messages/"+url.PathEscape(req.MessageID), nil, req.InstanceToken)
}

func (r *QueuesResource) ListOperations(ctx context.Context, req QueueRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(req.InstanceID)+"/queue/operations", nil, req.InstanceToken)
}

func (r *QueuesResource) PurgeOperations(ctx context.Context, req QueueRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodDelete, "/instances/"+url.PathEscape(req.InstanceID)+"/queue/operations", nil, req.InstanceToken)
}

func (r *QueuesResource) RemoveOperation(ctx context.Context, req QueueMessageRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodDelete, "/instances/"+url.PathEscape(req.InstanceID)+"/queue/operations/"+url.PathEscape(req.MessageID), nil, req.InstanceToken)
}

type TypebotsResource struct{ http *transport.Client }

func NewTypebots(httpClient *transport.Client) *TypebotsResource {
	return &TypebotsResource{http: httpClient}
}

func (r *TypebotsResource) List(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(instanceID)+"/typebots", nil, options.InstanceToken)
}

func (r *TypebotsResource) Create(ctx context.Context, req TypebotRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/typebots", req, req.InstanceToken)
}

func (r *TypebotsResource) Update(ctx context.Context, req TypebotMutationRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodPatch, "/instances/"+url.PathEscape(req.InstanceID)+"/typebots/"+url.PathEscape(req.TypebotID), req, req.InstanceToken)
}

func (r *TypebotsResource) Delete(ctx context.Context, req TypebotMutationRequest) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodDelete, "/instances/"+url.PathEscape(req.InstanceID)+"/typebots/"+url.PathEscape(req.TypebotID), nil, req.InstanceToken)
}

func (r *TypebotsResource) ListSessions(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(instanceID)+"/typebots/sessions", nil, options.InstanceToken)
}

func (r *TypebotsResource) StartSession(ctx context.Context, req TypebotStartSessionRequest) (map[string]any, error) {
	path := "/typebots/sessions/start"
	if req.TypebotID != "" {
		path = "/typebots/" + url.PathEscape(req.TypebotID) + "/sessions/start"
	}
	return instancePost(ctx, r.http, req.InstanceID, path, req, req.InstanceToken)
}

func (r *TypebotsResource) CloseSession(ctx context.Context, req TypebotSessionRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/typebots/sessions/"+url.PathEscape(req.Session)+"/close", req, req.InstanceToken)
}

func (r *TypebotsResource) PauseSession(ctx context.Context, req TypebotSessionRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/typebots/sessions/"+url.PathEscape(req.Session)+"/pause", req, req.InstanceToken)
}

type ChatwootResource struct{ http *transport.Client }

func NewChatwoot(httpClient *transport.Client) *ChatwootResource {
	return &ChatwootResource{http: httpClient}
}

func (r *ChatwootResource) Get(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(instanceID)+"/chatwoot", nil, options.InstanceToken)
}

func (r *ChatwootResource) Set(ctx context.Context, req ChatwootRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/chatwoot", req, req.InstanceToken)
}

func (r *ChatwootResource) Delete(ctx context.Context, instanceID string, options InstanceOptions) (map[string]any, error) {
	return requestMap(ctx, r.http, http.MethodDelete, "/instances/"+url.PathEscape(instanceID)+"/chatwoot", nil, options.InstanceToken)
}

func (r *ChatwootResource) TriggerImport(ctx context.Context, req ChatwootImportRequest) (map[string]any, error) {
	return instancePost(ctx, r.http, req.InstanceID, "/chatwoot/import/"+url.PathEscape(req.What), req, req.InstanceToken)
}

type ChatsResource struct{ http *transport.Client }

func NewChats(httpClient *transport.Client) *ChatsResource {
	return &ChatsResource{http: httpClient}
}

func (r *ChatsResource) List(ctx context.Context, req ChatListRequest) (map[string]any, error) {
	return requestMapQuery(ctx, r.http, http.MethodGet, "/instances/"+url.PathEscape(req.InstanceID)+"/chats", nil, req.InstanceToken, map[string]any{"page": req.Page, "pageSize": req.PageSize})
}

func (r *ChatsResource) Get(ctx context.Context, req ChatRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodGet, req.ChatRequest(), "", nil)
}

func (r *ChatsResource) Archive(ctx context.Context, req ChatRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodPost, req, "/archive", nil)
}

func (r *ChatsResource) Unarchive(ctx context.Context, req ChatRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodPost, req, "/unarchive", nil)
}

func (r *ChatsResource) Mute(ctx context.Context, req ChatMuteRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodPost, req.ChatRequest, "/mute", req)
}

func (r *ChatsResource) Unmute(ctx context.Context, req ChatRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodPost, req, "/unmute", nil)
}

func (r *ChatsResource) Pin(ctx context.Context, req ChatRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodPost, req, "/pin", nil)
}

func (r *ChatsResource) Unpin(ctx context.Context, req ChatRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodPost, req, "/unpin", nil)
}

func (r *ChatsResource) Read(ctx context.Context, req ChatRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodPost, req, "/read", nil)
}

func (r *ChatsResource) Unread(ctx context.Context, req ChatRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodPost, req, "/unread", nil)
}

func (r *ChatsResource) Clear(ctx context.Context, req ChatClearRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodPost, req.ChatRequest, "/clear", req)
}

func (r *ChatsResource) Delete(ctx context.Context, req ChatDeleteRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodDelete, req.ChatRequest, "", req)
}

func (r *ChatsResource) SetExpiration(ctx context.Context, req ChatExpirationRequest) (map[string]any, error) {
	return r.chatAction(ctx, http.MethodPut, req.ChatRequest, "/expiration", req)
}

func (r *ChatsResource) chatAction(ctx context.Context, method string, req ChatRequest, suffix string, body any) (map[string]any, error) {
	return requestMap(ctx, r.http, method, "/instances/"+url.PathEscape(req.InstanceID)+"/chats/"+url.PathEscape(req.Phone)+suffix, body, req.InstanceToken)
}

func (r ChatRequest) ChatRequest() ChatRequest { return r }

func instancePost(ctx context.Context, httpClient *transport.Client, instanceID, path string, body any, instanceToken string) (map[string]any, error) {
	return requestMap(ctx, httpClient, http.MethodPost, "/instances/"+url.PathEscape(instanceID)+path, body, instanceToken)
}

func requestMap(ctx context.Context, httpClient *transport.Client, method, path string, body any, instanceToken string) (map[string]any, error) {
	return requestMapQuery(ctx, httpClient, method, path, body, instanceToken, nil)
}

func requestMapQuery(ctx context.Context, httpClient *transport.Client, method, path string, body any, instanceToken string, query map[string]any) (map[string]any, error) {
	var out map[string]any
	err := httpClient.Request(ctx, method, path, transport.RequestOptions{
		Query:         query,
		Body:          body,
		InstanceToken: instanceToken,
	}, &out)
	return out, err
}
