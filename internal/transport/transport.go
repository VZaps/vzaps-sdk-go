package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/vzaps/vzaps-sdk-go/sdkerrors"
)

const (
	defaultBaseURL     = "https://api.vzaps.com"
	defaultRealtimeURL = "wss://realtime.vzaps.com"
	defaultTimeout     = 30 * time.Second
	defaultTokenSkew   = time.Minute
)

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Options struct {
	ClientToken  string
	ClientSecret string
	BaseURL      string
	RealtimeURL  string
	Timeout      time.Duration
	TokenSkew    time.Duration
	UserAgent    string
	HTTPClient   HTTPDoer
}

type RequestOptions struct {
	Query         map[string]any
	Body          any
	Headers       map[string]string
	NoAuth        bool
	InstanceToken string
}

type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type cachedToken struct {
	accessToken string
	expiresAt   time.Time
}

type Client struct {
	baseURL      string
	realtimeURL  string
	clientToken  string
	clientSecret string
	timeout      time.Duration
	tokenSkew    time.Duration
	userAgent    string
	http         HTTPDoer
	mu           sync.Mutex
	token        *cachedToken
}

func New(options Options) (*Client, error) {
	if strings.TrimSpace(options.ClientToken) == "" {
		return nil, sdkerrors.New("VZaps clientToken is required", 0, "", nil)
	}
	if strings.TrimSpace(options.ClientSecret) == "" {
		return nil, sdkerrors.New("VZaps clientSecret is required", 0, "", nil)
	}
	baseURL := strings.TrimRight(options.BaseURL, "/")
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	realtimeURL := strings.TrimRight(options.RealtimeURL, "/")
	if realtimeURL == "" {
		realtimeURL = defaultRealtimeURL
	}
	timeout := options.Timeout
	if timeout == 0 {
		timeout = defaultTimeout
	}
	tokenSkew := options.TokenSkew
	if tokenSkew == 0 {
		tokenSkew = defaultTokenSkew
	}
	doer := options.HTTPClient
	if doer == nil {
		doer = &http.Client{}
	}
	return &Client{baseURL: baseURL, realtimeURL: realtimeURL, clientToken: options.ClientToken, clientSecret: options.ClientSecret, timeout: timeout, tokenSkew: tokenSkew, userAgent: options.UserAgent, http: doer}, nil
}

func (c *Client) ClientToken() string { return c.clientToken }

func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
	c.mu.Lock()
	if c.token != nil && time.Now().Before(c.token.expiresAt) {
		token := c.token.accessToken
		c.mu.Unlock()
		return token, nil
	}
	c.mu.Unlock()
	var response AuthTokenResponse
	err := c.Request(ctx, http.MethodPost, "/token", RequestOptions{NoAuth: true, Body: map[string]string{"client_token": c.clientToken, "client_secret": c.clientSecret}}, &response)
	if err != nil {
		return "", err
	}
	if response.AccessToken == "" || response.ExpiresIn == 0 {
		return "", sdkerrors.NewAuthenticationError("VZaps token response is missing access_token or expires_in", response)
	}
	c.mu.Lock()
	c.token = &cachedToken{accessToken: response.AccessToken, expiresAt: time.Now().Add(time.Duration(response.ExpiresIn)*time.Second - c.tokenSkew)}
	c.mu.Unlock()
	return response.AccessToken, nil
}

func (c *Client) Request(ctx context.Context, method, path string, options RequestOptions, out any) error {
	if ctx == nil {
		ctx = context.Background()
	}
	reqCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	var body io.Reader
	if options.Body != nil {
		payload, err := json.Marshal(options.Body)
		if err != nil {
			return err
		}
		body = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(reqCtx, method, c.buildURL(path, options.Query), body)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	if options.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	for key, value := range options.Headers {
		if value != "" {
			req.Header.Set(key, value)
		}
	}
	if !options.NoAuth {
		token, err := c.GetAccessToken(ctx)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("X-Client-Token", c.clientToken)
	}
	if options.InstanceToken != "" {
		req.Header.Set("X-Instance-Token", options.InstanceToken)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		if errors.Is(reqCtx.Err(), context.DeadlineExceeded) {
			return sdkerrors.NewTimeoutError()
		}
		return err
	}
	defer resp.Body.Close()
	return parseResponse(resp, out)
}

func (c *Client) buildURL(path string, query map[string]any) string {
	u, _ := url.Parse(c.baseURL + "/" + strings.TrimLeft(path, "/"))
	values := u.Query()
	for key, value := range query {
		if value != nil {
			values.Set(toSnakeCase(key), fmt.Sprint(value))
		}
	}
	u.RawQuery = values.Encode()
	return u.String()
}

func (c *Client) BuildRealtimeURL(path string, query map[string]any) string {
	u, _ := url.Parse(c.realtimeURL + "/" + strings.TrimLeft(path, "/"))
	values := u.Query()
	for key, value := range query {
		if value != nil && fmt.Sprint(value) != "" {
			values.Set(key, fmt.Sprint(value))
		}
	}
	u.RawQuery = values.Encode()
	return u.String()
}

func parseResponse(resp *http.Response, out any) error {
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var data any
	if len(text) > 0 {
		if err := json.Unmarshal(text, &data); err != nil {
			data = string(text)
		}
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		message := readErrorMessage(data, resp.Status)
		if resp.StatusCode == http.StatusUnauthorized {
			return sdkerrors.NewAuthenticationError(message, data)
		}
		return sdkerrors.New(message, resp.StatusCode, "", data)
	}
	if out == nil || len(text) == 0 {
		return nil
	}
	return json.Unmarshal(text, out)
}

func readErrorMessage(data any, fallback string) string {
	if m, ok := data.(map[string]any); ok {
		if value, ok := m["error"].(string); ok && value != "" {
			return value
		}
		if value, ok := m["message"].(string); ok && value != "" {
			return value
		}
	}
	if fallback != "" {
		return fallback
	}
	return "VZaps request failed"
}

func toSnakeCase(key string) string {
	var out strings.Builder
	for i, r := range key {
		if i > 0 && r >= 'A' && r <= 'Z' {
			out.WriteByte('_')
		}
		out.WriteRune(r)
	}
	return strings.ToLower(out.String())
}
