package api

import (
	"fmt"

	"api-server/internal/pkg/api/errors"
)

const (
	InvalidRequest errors.ErrorCode = "article-01"
	NotFound       errors.ErrorCode = "article-02"
	Unknown        errors.ErrorCode = "article-03"
)

type InvalidRequestError struct {
	message string
}

func (e InvalidRequestError) Error() string {
	return fmt.Sprintf("invalid request. %s", e.message)
}

func NewInvalidRequestError(message string) InvalidRequestError {
	return InvalidRequestError{message: message}
}
