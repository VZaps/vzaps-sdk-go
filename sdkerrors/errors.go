package sdkerrors

import "fmt"

// Error represents a VZaps API or SDK error.
type Error struct {
	Message string
	Status  int
	Code    string
	Details any
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.Status > 0 {
		return fmt.Sprintf("vzaps: %s (status %d)", e.Message, e.Status)
	}
	return "vzaps: " + e.Message
}

// AuthenticationError is returned when credentials or JWT authentication fail.
type AuthenticationError struct{ Base *Error }

func (e *AuthenticationError) Error() string {
	if e == nil || e.Base == nil {
		return ""
	}
	return e.Base.Error()
}

// TimeoutError is returned when an HTTP request exceeds the configured timeout.
type TimeoutError struct{ Base *Error }

func (e *TimeoutError) Error() string {
	if e == nil || e.Base == nil {
		return ""
	}
	return e.Base.Error()
}

func New(message string, status int, code string, details any) *Error {
	return &Error{Message: message, Status: status, Code: code, Details: details}
}

func NewAuthenticationError(message string, details any) *AuthenticationError {
	if message == "" {
		message = "invalid VZaps client credentials"
	}
	return &AuthenticationError{Base: &Error{Message: message, Status: 401, Code: "AUTHENTICATION_FAILED", Details: details}}
}

func NewTimeoutError() *TimeoutError {
	return &TimeoutError{Base: &Error{Message: "VZaps request timed out", Code: "REQUEST_TIMEOUT"}}
}
