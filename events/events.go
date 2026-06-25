package events

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/vzaps/vzaps-sdk-go/internal/transport"
)

type WebSocketConn interface {
	ReadMessage() (int, []byte, error)
	WriteJSON(v any) error
	WriteMessage(messageType int, data []byte) error
	Close() error
}

type WebSocketDialer interface {
	Dial(urlStr string, requestHeader http.Header) (WebSocketConn, *http.Response, error)
}

type gorillaDialer struct {
	dialer *websocket.Dialer
}

func (d gorillaDialer) Dial(urlStr string, requestHeader http.Header) (WebSocketConn, *http.Response, error) {
	return d.dialer.Dial(urlStr, requestHeader)
}

type Resource struct {
	http   *transport.Client
	dialer WebSocketDialer
}

func NewResource(httpClient *transport.Client, dialer WebSocketDialer) *Resource {
	return &Resource{http: httpClient, dialer: dialer}
}

func (r *Resource) Subscribe(ctx context.Context, req SubscribeRequest) (*Subscription, error) {
	sub := &Subscription{
		http:              r.http,
		request:           req,
		dialer:            r.dialer,
		handlers:          map[EventType][]EventHandler{},
		lifecycleHandlers: map[string][]LifecycleHandler{},
		errors:            []ErrorHandler{},
	}
	if err := sub.open(ctx); err != nil {
		return nil, err
	}
	return sub, nil
}

type EventHandler func(Event)
type ErrorHandler func(error)
type LifecycleHandler func()

type Subscription struct {
	http    *transport.Client
	request SubscribeRequest
	dialer  WebSocketDialer

	mu                sync.RWMutex
	conn              WebSocketConn
	closed            bool
	retryCount        int
	handlers          map[EventType][]EventHandler
	errors            []ErrorHandler
	lifecycleHandlers map[string][]LifecycleHandler
}

func (s *Subscription) On(event EventType, handler EventHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[event] = append(s.handlers[event], handler)
}

func (s *Subscription) OnError(handler ErrorHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.errors = append(s.errors, handler)
}

func (s *Subscription) OnOpen(handler LifecycleHandler) {
	s.addLifecycle("open", handler)
}

func (s *Subscription) OnClose(handler LifecycleHandler) {
	s.addLifecycle("close", handler)
}

func (s *Subscription) Close() error {
	s.mu.Lock()
	s.closed = true
	conn := s.conn
	s.mu.Unlock()
	if conn == nil {
		return nil
	}
	_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closed subscription"))
	return conn.Close()
}

func (s *Subscription) open(ctx context.Context) error {
	token, err := s.http.GetAccessToken(ctx)
	if err != nil {
		return err
	}

	events := make([]string, 0, len(s.request.Events))
	for _, event := range s.request.Events {
		events = append(events, string(event))
	}

	url := s.http.BuildRealtimeURL("/events/ws", map[string]any{
		"instance_id":    s.request.InstanceID,
		"events":         strings.Join(events, ","),
		"access_token":   token,
		"client_token":   s.http.ClientToken(),
		"instance_token": s.request.InstanceToken,
		"last_event_id":  s.request.LastEventID,
	})
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+token)
	headers.Set("X-Client-Token", s.http.ClientToken())
	headers.Set("X-Instance-Token", s.request.InstanceToken)

	dialer := s.dialer
	if dialer == nil {
		dialer = gorillaDialer{dialer: websocket.DefaultDialer}
	}

	conn, _, err := dialer.Dial(url, headers)
	if err != nil {
		return err
	}

	s.mu.Lock()
	s.conn = conn
	s.retryCount = 0
	s.mu.Unlock()
	s.emitLifecycle("open")

	go s.readLoop(ctx, conn)
	return nil
}

func (s *Subscription) readLoop(ctx context.Context, conn WebSocketConn) {
	defer func() {
		s.emitLifecycle("close")
		go s.reconnectIfNeeded(ctx)
	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			if !s.isClosed() {
				s.emitError(err)
			}
			return
		}
		var event Event
		if err := json.Unmarshal(data, &event); err != nil {
			s.emitError(err)
			continue
		}
		s.dispatch(event)
		_ = conn.WriteJSON(map[string]string{"type": "ack", "event_id": event.ID})
	}
}

func (s *Subscription) dispatch(event Event) {
	s.mu.RLock()
	handlers := append([]EventHandler{}, s.handlers[event.Type]...)
	handlers = append(handlers, s.handlers[EventAll]...)
	s.mu.RUnlock()

	for _, handler := range handlers {
		handler(event)
	}
}

func (s *Subscription) reconnectIfNeeded(ctx context.Context) {
	if s.isClosed() || !s.request.Reconnect {
		return
	}

	s.mu.Lock()
	maxRetries := s.request.MaxRetries
	if maxRetries > 0 && s.retryCount >= maxRetries {
		s.mu.Unlock()
		return
	}
	s.retryCount++
	retryCount := s.retryCount
	s.mu.Unlock()

	delay := s.request.RetryDelay
	if delay == 0 {
		delay = time.Duration(retryCount) * time.Second
		if delay > 30*time.Second {
			delay = 30 * time.Second
		}
	}

	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return
	case <-timer.C:
	}
	if err := s.open(ctx); err != nil {
		s.emitError(err)
		go s.reconnectIfNeeded(ctx)
	}
}

func (s *Subscription) addLifecycle(event string, handler LifecycleHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lifecycleHandlers[event] = append(s.lifecycleHandlers[event], handler)
}

func (s *Subscription) emitLifecycle(event string) {
	s.mu.RLock()
	handlers := append([]LifecycleHandler{}, s.lifecycleHandlers[event]...)
	s.mu.RUnlock()
	for _, handler := range handlers {
		handler()
	}
}

func (s *Subscription) emitError(err error) {
	s.mu.RLock()
	handlers := append([]ErrorHandler{}, s.errors...)
	s.mu.RUnlock()
	for _, handler := range handlers {
		handler(err)
	}
}

func (s *Subscription) isClosed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.closed
}
