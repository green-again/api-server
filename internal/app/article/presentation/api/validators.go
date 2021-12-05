package api

import (
	"fmt"

	apipkg "api-server/internal/pkg/api"
	"github.com/google/uuid"
)

type GetArticleRequestValidator struct{}

func (v GetArticleRequestValidator) Validate(requestID interface{}) error {
	val, ok := requestID.(string)
	if !ok {
		return apipkg.NewInvalidRequestError(fmt.Sprintf("request id %s is not string.", requestID))
	}
	_, err := uuid.Parse(val)
	if err != nil {
		return apipkg.NewInvalidRequestError(fmt.Sprintf("request id %s is not uuid format.", requestID))
	}
	return nil
}

func NewGetArticleRequestValidator() apipkg.Validator {
	return GetArticleRequestValidator{}
}
