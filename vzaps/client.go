package vzaps

import (
	"context"
	"time"

	"github.com/vzaps/vzaps-sdk-go/events"
	"github.com/vzaps/vzaps-sdk-go/internal/transport"
	"github.com/vzaps/vzaps-sdk-go/resources"
)

type ClientOptions struct {
	ClientToken  string
	ClientSecret string
	BaseURL      string
	RealtimeURL  string
	Timeout      time.Duration
	TokenSkew    time.Duration
	UserAgent    string
	HTTPClient   transport.HTTPDoer
	Dialer       events.WebSocketDialer
}

type RequestOptions = transport.RequestOptions

type Client struct {
	Auth      *resources.AuthResource
	Instances *resources.InstancesResource
	Sessions  *resources.SessionsResource
	Messages  *resources.MessagesResource
	Webhooks  *resources.WebhooksResource
	Contacts  *resources.ContactsResource
	Groups    *resources.GroupsResource
	Users     *resources.UsersResource
	Queues    *resources.QueuesResource
	Typebots  *resources.TypebotsResource
	Chatwoot  *resources.ChatwootResource
	Chats     *resources.ChatsResource
	Events    *events.Resource

	http *transport.Client
}

func NewClient(options ClientOptions) (*Client, error) {
	httpClient, err := transport.New(transport.Options{
		ClientToken: options.ClientToken, ClientSecret: options.ClientSecret, BaseURL: options.BaseURL,
		RealtimeURL: options.RealtimeURL, Timeout: options.Timeout, TokenSkew: options.TokenSkew,
		UserAgent: options.UserAgent, HTTPClient: options.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		http: httpClient,
		Auth: resources.NewAuth(httpClient), Instances: resources.NewInstances(httpClient), Sessions: resources.NewSessions(httpClient),
		Messages: resources.NewMessages(httpClient), Webhooks: resources.NewWebhooks(httpClient), Contacts: resources.NewContacts(httpClient),
		Groups: resources.NewGroups(httpClient), Users: resources.NewUsers(httpClient), Queues: resources.NewQueues(httpClient),
		Typebots: resources.NewTypebots(httpClient), Chatwoot: resources.NewChatwoot(httpClient), Chats: resources.NewChats(httpClient),
		Events: events.NewResource(httpClient, options.Dialer),
	}, nil
}

func MustNewClient(options ClientOptions) *Client {
	client, err := NewClient(options)
	if err != nil {
		panic(err)
	}
	return client
}

func (c *Client) Request(ctx context.Context, method, path string, options RequestOptions, out any) error {
	return c.http.Request(ctx, method, path, options, out)
}
