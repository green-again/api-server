package api

import (
	"fmt"

	"github.com/google/uuid"

	"api-server/internal/pkg/api/validators"
)

type GetArticleRequestValidator struct{}

func (v GetArticleRequestValidator) Validate(requestID interface{}) error {
	val, ok := requestID.(string)
	if !ok {
		return NewInvalidRequestError(fmt.Sprintf("request id %s is not string.", requestID))
	}
	_, err := uuid.Parse(val)
	if err != nil {
		return NewInvalidRequestError(fmt.Sprintf("request id %s is not uuid format.", requestID))
	}
	return nil
}

func NewGetArticleRequestValidator() validators.Validator {
	return GetArticleRequestValidator{}
}
