package resources

type InstanceOptions struct {
	InstanceToken string
}

type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type InstanceCreateRequest struct {
	Name            string   `json:"name"`
	Webhook         string   `json:"webhook,omitempty"`
	EventsSubscribe []string `json:"events_subscribe,omitempty"`
}

type InstanceListRequest struct {
	Page     int            `json:"page,omitempty"`
	Size     int            `json:"size,omitempty"`
	PageSize int            `json:"-"`
	Filter   map[string]any `json:"filter,omitempty"`
	Search   string         `json:"-"`
	Sort     string         `json:"sort,omitempty"`
	SortDesc bool           `json:"sort_desc,omitempty"`
}

type InstanceScopedRequest struct {
	InstanceID    string `json:"-"`
	InstanceToken string `json:"-"`
}

type MessageSendBaseRequest struct {
	InstanceScopedRequest
	Phone string `json:"phone"`
}

type MessageSendTextRequest struct {
	MessageSendBaseRequest
	Message string `json:"message"`
}

type MessageSendImageRequest struct {
	MessageSendBaseRequest
	Image   string `json:"image"`
	Caption string `json:"caption,omitempty"`
}

type MessageSendAudioRequest struct {
	MessageSendBaseRequest
	Audio string `json:"audio"`
	PTT   bool   `json:"ptt,omitempty"`
}

type MessageSendDocumentRequest struct {
	MessageSendBaseRequest
	Document string `json:"document"`
	FileName string `json:"file_name,omitempty"`
	Caption  string `json:"caption,omitempty"`
}

type MessageSendVideoRequest struct {
	MessageSendBaseRequest
	Video   string `json:"video"`
	Caption string `json:"caption,omitempty"`
}

type MessageSendStickerRequest struct {
	MessageSendBaseRequest
	Sticker string `json:"sticker"`
}

type MessageSendGifRequest struct {
	MessageSendBaseRequest
	GIF     string `json:"gif"`
	Caption string `json:"caption,omitempty"`
}

type MessageSendLocationRequest struct {
	MessageSendBaseRequest
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name,omitempty"`
	Address   string  `json:"address,omitempty"`
}

type MessageSendContactRequest struct {
	MessageSendBaseRequest
	ContactName  string `json:"contact_name,omitempty"`
	ContactPhone string `json:"contact_phone,omitempty"`
}

type MessageButton struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type MessageSendButtonsRequest struct {
	MessageSendBaseRequest
	Message string          `json:"message"`
	Buttons []MessageButton `json:"buttons"`
	Footer  string          `json:"footer,omitempty"`
}

type MessageListRow struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type MessageListSection struct {
	Title string           `json:"title"`
	Rows  []MessageListRow `json:"rows"`
}

type MessageSendListRequest struct {
	MessageSendBaseRequest
	Title       string               `json:"title"`
	Description string               `json:"description"`
	ButtonText  string               `json:"button_text"`
	Sections    []MessageListSection `json:"sections"`
	Footer      string               `json:"footer,omitempty"`
}

type MessageSendLinkRequest struct {
	MessageSendBaseRequest
	Message         string `json:"message"`
	LinkURL         string `json:"link_url"`
	Title           string `json:"title"`
	LinkDescription string `json:"link_description"`
	JPEGThumbnail   string `json:"jpeg_thumbnail,omitempty"`
}

type MessageSendPollRequest struct {
	MessageSendBaseRequest
	Name                   string   `json:"name"`
	Options                []string `json:"options"`
	SelectableOptionsCount int      `json:"selectable_options_count,omitempty"`
	HideParticipantNames   bool     `json:"hide_participant_names,omitempty"`
	EndTime                string   `json:"end_time,omitempty"`
	AllowAddOption         bool     `json:"allow_add_option,omitempty"`
}

type MessagePollVoteRequest struct {
	MessageSendBaseRequest
	MessageID       string   `json:"message_id"`
	Vote            any      `json:"vote,omitempty"`
	SelectedOptions []string `json:"selected_options,omitempty"`
	PollSender      string   `json:"poll_sender,omitempty"`
	FromMe          bool     `json:"from_me,omitempty"`
}

type MessageReactRequest struct {
	MessageSendBaseRequest
	MessageID string `json:"message_id"`
	Reaction  string `json:"reaction"`
}

type MessageReactRemoveRequest struct {
	MessageSendBaseRequest
	MessageID string `json:"message_id"`
}

type MessagePresenceRequest struct {
	MessageSendBaseRequest
	State string `json:"state"`
	Media string `json:"media,omitempty"`
}

type MessageMarkReadRequest struct {
	InstanceScopedRequest
	ID     []string `json:"id"`
	Chat   string   `json:"chat"`
	Sender string   `json:"sender,omitempty"`
}

type MessageDownloadRequest struct {
	InstanceScopedRequest
}

type MessageEditRequest struct {
	InstanceScopedRequest
	MessageID string `json:"-"`
	Message   string `json:"message"`
}

type MessageDeleteRequest struct {
	InstanceScopedRequest
	MessageID string `json:"-"`
}

type WebhookConfigRequest struct {
	InstanceScopedRequest
	WebhookURL string   `json:"webhook_url"`
	Events     []string `json:"events,omitempty"`
}

type WebhookLogSearchRequest struct {
	InstanceScopedRequest
}

type WebhookLogRequest struct {
	InstanceScopedRequest
	LogID string `json:"-"`
}

type ContactAddRequest struct {
	InstanceScopedRequest
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	FullName  string `json:"full_name"`
}

type UserPhonesRequest struct {
	InstanceScopedRequest
	Phone string `json:"phone,omitempty"`
}

type UserAvatarRequest struct {
	InstanceScopedRequest
	Phone string `json:"phone,omitempty"`
}

type GroupListRequest struct {
	InstanceScopedRequest
	Page     int `json:"-"`
	PageSize int `json:"-"`
}

type GroupInfoRequest struct {
	InstanceScopedRequest
	GroupID string `json:"group_id"`
}

type GroupInviteLinkRequest struct {
	GroupInfoRequest
	Reset bool `json:"reset,omitempty"`
}

type GroupMutationRequest struct {
	InstanceScopedRequest
	GroupID           string   `json:"group_id,omitempty"`
	Image             string   `json:"image,omitempty"`
	Name              string   `json:"name,omitempty"`
	Description       string   `json:"description,omitempty"`
	AdminOnlyMessage  bool     `json:"admin_only_message,omitempty"`
	AdminOnlySettings bool     `json:"admin_only_settings,omitempty"`
	DelayMessage      int      `json:"delay_message,omitempty"`
	GroupName         string   `json:"group_name,omitempty"`
	GroupDescription  string   `json:"group_description,omitempty"`
	GroupImage        string   `json:"group_image,omitempty"`
	Participants      []string `json:"participants,omitempty"`
}

type QueueRequest struct {
	InstanceScopedRequest
}

type QueueMessageRequest struct {
	InstanceScopedRequest
	MessageID string `json:"-"`
}

type TypebotRequest struct {
	InstanceScopedRequest
	Enabled         bool   `json:"enabled,omitempty"`
	Description     string `json:"description,omitempty"`
	TypebotURL      string `json:"typebot_url,omitempty"`
	PublicID        string `json:"public_id,omitempty"`
	TriggerType     string `json:"trigger_type,omitempty"`
	TriggerOperator string `json:"trigger_operator,omitempty"`
	TriggerValue    string `json:"trigger_value,omitempty"`
	Priority        int    `json:"priority,omitempty"`
	ExpireInMinutes int    `json:"expire_in_minutes,omitempty"`
	KeywordFinish   string `json:"keyword_finish,omitempty"`
	DefaultDelayMs  int    `json:"default_delay_ms,omitempty"`
	UnknownMessage  string `json:"unknown_message,omitempty"`
	ListenFromMe    bool   `json:"listen_from_me,omitempty"`
	StopBotFromMe   bool   `json:"stop_bot_from_me,omitempty"`
	KeepOpen        bool   `json:"keep_open,omitempty"`
	DebounceMs      int    `json:"debounce_ms,omitempty"`
	IgnoreGroups    bool   `json:"ignore_groups,omitempty"`
	TranscribeAudio bool   `json:"transcribe_audio,omitempty"`
}

type TypebotMutationRequest struct {
	TypebotRequest
	TypebotID string `json:"-"`
}

type TypebotSessionRequest struct {
	InstanceScopedRequest
	Session string `json:"-"`
}

type TypebotStartSessionRequest struct {
	InstanceScopedRequest
	TypebotID string `json:"-"`
	PublicID  string `json:"public_id,omitempty"`
	Phone     string `json:"phone"`
	PushName  string `json:"push_name,omitempty"`
	Message   string `json:"message"`
}

type ChatwootRequest struct {
	InstanceScopedRequest
	Enabled                 bool   `json:"enabled,omitempty"`
	URL                     string `json:"url,omitempty"`
	AccountID               string `json:"account_id,omitempty"`
	Token                   string `json:"token,omitempty"`
	NameInbox               string `json:"name_inbox,omitempty"`
	SignMsg                 bool   `json:"sign_msg,omitempty"`
	SignDelimiter           string `json:"sign_delimiter,omitempty"`
	Number                  string `json:"number,omitempty"`
	ReopenConversation      bool   `json:"reopen_conversation,omitempty"`
	ConversationPending     bool   `json:"conversation_pending,omitempty"`
	ImportContacts          bool   `json:"import_contacts,omitempty"`
	ImportMessages          bool   `json:"import_messages,omitempty"`
	DaysLimitImportMessages int    `json:"days_limit_import_messages,omitempty"`
	AutoCreate              bool   `json:"auto_create,omitempty"`
	Organization            string `json:"organization,omitempty"`
	Logo                    string `json:"logo,omitempty"`
	IgnoreJids              any    `json:"ignore_jids,omitempty"`
	IgnoreGroups            bool   `json:"ignore_groups,omitempty"`
}

type ChatwootImportRequest struct {
	InstanceScopedRequest
	What string `json:"-"`
}

type ChatRequest struct {
	InstanceScopedRequest
	Phone string `json:"-"`
}

type ChatListRequest struct {
	InstanceScopedRequest
	Page     int `json:"-"`
	PageSize int `json:"-"`
}

type ChatDeleteRequest struct {
	ChatRequest
	DeleteMedia bool `json:"delete_media,omitempty"`
}

type ChatMuteRequest struct {
	ChatRequest
	DurationSeconds int `json:"duration_seconds,omitempty"`
}

type ChatClearRequest struct {
	ChatRequest
	DeleteMedia bool `json:"delete_media,omitempty"`
}

type ChatExpirationRequest struct {
	ChatRequest
	Expiration string `json:"expiration"`
}
