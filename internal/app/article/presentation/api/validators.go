package api

import (
	"fmt"

	"api-server/internal/pkg/http"
	"github.com/google/uuid"
)

type GetArticleRequestValidator struct{}

func (v GetArticleRequestValidator) Validate(requestID interface{}) error {
	val, ok := requestID.(string)
	if !ok {
		return http.NewInvalidRequestError(fmt.Sprintf("request id %s is not string.", requestID))
	}
	_, err := uuid.Parse(val)
	if err != nil {
		return http.NewInvalidRequestError(fmt.Sprintf("request id %s is not uuid format.", requestID))
	}
	return nil
}

func NewGetArticleRequestValidator() http.Validator {
	return GetArticleRequestValidator{}
}
