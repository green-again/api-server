package domain

import "time"

type Article struct {
	ID string
	Title string
	Author string
	Source string

	Body string

	PublishedDate time.Time
}
