package api

import (
	"net/http"
	"strings"
)

const (
	BadRequest     = "BAD_REQUEST"
	InternalServer = "INTERNAL_SERVER_ERROR"
)

// HTTPError - represnts a standard error structure for the api.
type HTTPError struct {
	Err Error `json:"error"`
}

// Error - satisfies the error interface.
func (e HTTPError) Error() string {
	return e.Err.Message
}

// Error - represents base http error structure.
type Error struct {
	Code       string         `json:"code"`
	Message    string         `json:"message"`
	StatusCode int            `json:"status_code"`
	Details    map[string]any `json:"details,omitempty"`
}

// New - returns a new instance of Error with the provided details.
func New(statusCode int, code, message string, details map[string]any) HTTPError {
	if details == nil {
		details = map[string]any{}
	}
	return HTTPError{
		Err: Error{
			Code:       code,
			Message:    message,
			StatusCode: statusCode,
			Details:    details,
		},
	}
}

// ValidationError - represents a wrapper for holding input validation error details.
type ValidationError struct {
	details map[string][]string
}

// NewValidationError - returns a new ValidationError.
func NewValidationError() *ValidationError {
	return &ValidationError{
		details: map[string][]string{},
	}
}

// Add - will add key value details to the ValidationError.
func (v *ValidationError) Add(key string, values ...string) {
	if _, ok := v.details[key]; !ok {
		v.details[key] = values
		return
	}
	v.details[key] = append(v.details[key], values...)
}

// Details - will return all the details.
func (v *ValidationError) Details() map[string]interface{} {
	m := make(map[string]interface{}, len(v.details))
	for k, val := range v.details {
		m[k] = val
	}
	return m
}

// IsClean - will return wether or not there are any validation errors.
func (v *ValidationError) IsClean() bool {
	return len(v.details) == 0
}

// Error - satisfies the error interface.
func (v *ValidationError) Error() string {
	var b strings.Builder
	for k, val := range v.details {
		b.WriteString(k)
		b.WriteString(": ")
		for _, vv := range val {
			b.WriteString(vv)
		}
		b.WriteString("\n")
	}
	return b.String()
}

// BadRequestError - returns a new instance of a bad request error.
func BadRequestError(msg string, err error, details map[string]any) HTTPError {
	return buildError(http.StatusBadRequest, BadRequest, msg, err, details)
}

// InternalServerError - returns a new instance of the error with an internal server error message and status codes.
func InternalServerError(msg string, err error, details map[string]any) HTTPError {
	return buildError(http.StatusInternalServerError, InternalServer, msg, err, details)
}

func buildError(statusCode int, code, msg string, err error, details map[string]any) HTTPError {
	if details == nil {
		details = map[string]any{}
	}
	if err != nil {
		details["error"] = err.Error()
	}
	return New(statusCode, code, msg, details)
}
