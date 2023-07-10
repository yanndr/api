package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var NotImplementedErr = fmt.Errorf("not implemented")

// EndpointHandler represents a request method that returns a model.Response.
type EndpointHandler func(r *http.Request) (Response, error)

// Route represents a specific route of a server. It contains all the EndpointHandler mapped
// to their http.methods and all the sub routes
type Route struct {
	Pattern   string
	Actions   map[string]EndpointHandler
	SubRoutes []*Route
}

// Handle responds to an HTTP request,  executes the action EndpointHandler
// and writes the result to the http.ResponseWriter
func Handle(action EndpointHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response, err := action(r)
		if err != nil {
			response = getErrorResponse(err)
		}
		switch t := response.Content.(type) {
		case Serializable:
			eTag, err := t.Serialize()
			if err != nil {
				log.Printf("cannot generate the ETag")
			} else {
				w.Header().Set("ETag", eTag)
			}
		}
		w.WriteHeader(response.StatusCode)
		_ = json.NewEncoder(w).Encode(response)
	}
}

func NotImplementedHandler(_ *http.Request) (Response, error) {
	return Response{}, NotImplementedErr
}

func getErrorResponse(err error) Response {
	var statusCode int
	var errorCode string
	var errMessage string
	var details string
	var statusError StatusErr
	if errors.As(err, &statusError) {
		statusCode = statusError.StatusCode
		errorCode = statusError.ErrorCode
		errMessage = statusError.ErrorMessage
		details = statusError.ErrorDetails
	} else {
		statusCode = http.StatusInternalServerError
		errorCode = UnexpectedErrorCode
		errMessage = UnexpectedErrorMessage
		details = err.Error()
	}

	return NewErrorResponse(statusCode, errMessage, errorCode, details)
}
