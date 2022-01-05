package domain

import "errors"

var (
	InvalidStatusError = errors.New("invalid status error")
	AlreadyPublishedError = errors.New("published article cannot be changed back to a draft")
)
