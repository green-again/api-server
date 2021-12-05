package api

import "github.com/labstack/echo/v4"

type Binder interface {
	Bind(c echo.Context, model interface{}) error
}

type RequestBinder struct {}

func (b RequestBinder) Bind(c echo.Context, model interface{}) error {
	if err := c.Bind(model); err != nil {
		return NewInvalidRequestError(err.Error())
	}

	if err := c.Validate(model); err != nil {
		return NewInvalidRequestError(err.Error())
	}
	return nil
}
