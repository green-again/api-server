package persistence

import (
	"errors"

	"gorm.io/gorm"
)

var (
	NotFoundError        = gorm.ErrRecordNotFound
	ParseUUIDError       = errors.New("parse uuid error")
	UnmarshalBinaryError = errors.New("the article id unmarshal error")
)
