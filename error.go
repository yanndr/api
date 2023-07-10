package api

import (
	"fmt"
	"net/http"
)

const (
	UnexpectedErrorCode         = "ERR-UNEXPECTED"
	DuplicateErrorCode          = "ERR-DUPLICATE"
	BadFormatErrorCode          = "ERR-BAD-FORMAT"
	NotFoundErrorCode           = "ERR-NOT-FOUND"
	PreconditionFailedErrorCode = "ERR-PRE-FAILED"
	UnexpectedErrorMessage      = "An unexpected error occurred."
)

// StatusErr is the API error that contains all the information for the model.ResponseAPI
type StatusErr struct {
	StatusCode   int
	ErrorCode    string
	ErrorMessage string
	ErrorDetails string
}

func (s StatusErr) Error() string {
	return fmt.Sprintf("error %v - %v, %v", s.ErrorCode, s.StatusCode, s.ErrorMessage)
}

func NewBadFormatErr(details string) StatusErr {
	return StatusErr{
		StatusCode:   http.StatusBadRequest,
		ErrorCode:    BadFormatErrorCode,
		ErrorMessage: "The format of data sent to the server was not expected",
		ErrorDetails: details,
	}
}

func NewDuplicateErr(details string) StatusErr {
	return StatusErr{
		StatusCode:   http.StatusConflict,
		ErrorCode:    DuplicateErrorCode,
		ErrorMessage: "A duplicate of the entity has been found.",
		ErrorDetails: details,
	}
}

func NewNotFoundErr(details string) StatusErr {
	return StatusErr{
		StatusCode:   http.StatusNotFound,
		ErrorCode:    NotFoundErrorCode,
		ErrorMessage: "The requested resource was not found",
		ErrorDetails: details,
	}
}

func NewPreconditionFailError() StatusErr {
	return StatusErr{
		StatusCode:   http.StatusPreconditionFailed,
		ErrorCode:    PreconditionFailedErrorCode,
		ErrorMessage: "Precondition failed",
	}
}
