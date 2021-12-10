package domain_test

import (
	"github.com/google/uuid"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"

	"api-server/internal/app/article/domain"
)

const expectNewUUID = ""

var existedUUID = uuid.NewString()

func TestArticle(t *testing.T) {
	tests := []struct{
		scenario string

		inputID  string
		expected string
	} {
		{
			scenario: "When a new article is created, a new UUID must be created.",
			expected: expectNewUUID,
		},
		{
			scenario: "If the article already exists, it returns with the existing UUID.",
			inputID: existedUUID,
			expected: existedUUID,
		},
	}

	for _, tt := range tests {
		actual := domain.NewArticle(tt.inputID, faker.Sentence(), faker.Name(), faker.URL(), faker.Paragraph(), 1)
		if tt.expected == expectNewUUID {
			_, err := uuid.Parse(actual.ID())
			assert.NoError(t, err)
		} else {
			assert.Equal(t, tt.expected, actual.ID())
		}
	}
}
