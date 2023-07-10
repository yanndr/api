package api

import (
	"encoding/json"
	"net/http"
)

const (
	SuccessStatus = "Success"
	ErrorStatus   = "Error"
)

// Response represents the default response to an API request.
type Response struct {
	StatusCode   int    `json:"status_code"`
	Status       string `json:"status"`
	Content      any    `json:"content,omitempty"`
	Location     string `json:"location,omitempty"`
	Error        string `json:"error,omitempty"`
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorDetails string `json:"error_details,omitempty"`
}

// NewContentResponse creates a new Response with StatusCode "Status Ok" and Status "Success".
// It also sets the content from the parameter.
func NewContentResponse(content any) Response {
	return Response{
		StatusCode: http.StatusOK,
		Status:     SuccessStatus,
		Content:    content,
	}
}

// NewErrorResponse creates a new Response with Status "ErrorStatus"
// and sets the StatusCode, Error and ErrorDetails from the parameters.
func NewErrorResponse(statusCode int, err, errorCode, details string) Response {
	return Response{
		StatusCode:   statusCode,
		Status:       ErrorStatus,
		Error:        err,
		ErrorDetails: details,
		ErrorCode:    errorCode,
	}
}

// NewCreatedResponse creates a new Response with StatusCode "StatusCreated" and Status "Success".
// Its also sets the location from the parameter
func NewCreatedResponse(location string) Response {
	return Response{
		StatusCode: http.StatusCreated,
		Status:     SuccessStatus,
		Location:   location,
	}
}

// NewEmptyResponse creates a new Response with StatusCode "StatusOK" and Status "Success".
func NewEmptyResponse() Response {
	return Response{
		StatusCode: http.StatusOK,
		Status:     SuccessStatus,
	}
}

// RawResponse is similar to Response but is meant to be used to decode a response.
type RawResponse struct {
	StatusCode   int             `json:"status_code"`
	Status       string          `json:"status"`
	Content      json.RawMessage `json:"content,omitempty"`
	Location     string          `json:"location,omitempty"`
	Error        string          `json:"error,omitempty"`
	ErrorCode    string          `json:"error_code,omitempty"`
	ErrorDetails string          `json:"error_details,omitempty"`
}
