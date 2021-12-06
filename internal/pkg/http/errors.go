package http

import (
	"errors"
)

type ErrorCode string

type ErrorResponse struct {
	ErrorCode ErrorCode `json:"error"`
	Message   string    `json:"message"`
	Detail    string    `json:"detail"`
}

func NewErrorResponse(code ErrorCode, message, detail string) ErrorResponse {
	return ErrorResponse{
		ErrorCode: code,
		Message:   message,
		Detail:    detail,
	}
}

var (
	InvalidRequestError = errors.New("invalid request error")
)
