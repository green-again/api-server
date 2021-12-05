package api

import "fmt"

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

type InvalidRequestError struct {
	message string
}

func (e InvalidRequestError) Error() string {
	return fmt.Sprintf("invalid request. %s", e.message)
}

func NewInvalidRequestError(message string) InvalidRequestError {
	return InvalidRequestError{message: message}
}
