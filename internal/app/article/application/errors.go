package application

import "fmt"

type NotFoundError struct {
	requestID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("article id %s does not found.", e.requestID)
}

func NewNotFoundError(requestID string) NotFoundError {
	return NotFoundError{requestID: requestID}
}

type UnknownError struct {
	message string
}

func (e UnknownError) Error() string {
	return fmt.Sprintf("unknown error. %s", e.message)
}

func NewUnknownError(message string) UnknownError {
	return UnknownError{message: message}
}
