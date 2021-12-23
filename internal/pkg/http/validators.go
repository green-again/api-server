package http

import (
	"github.com/go-playground/validator"
)

type Validator interface {
	Validate(interface{}) error
}

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(i interface{}) error {
	return rv.validator.Struct(i)
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{validator: validator.New()}
}
