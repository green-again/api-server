package http

import "github.com/labstack/echo/v4"

type Binder interface {
	Bind(c echo.Context, model interface{}) error
}

type RequestBinder struct{}

func (b RequestBinder) Bind(c echo.Context, model interface{}) error {
	if err := c.Bind(model); err != nil {
		return InvalidRequestError
	}

	if err := c.Validate(model); err != nil {
		return InvalidRequestError
	}
	return nil
}
