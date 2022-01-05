package domain_test

import (
	"github.com/google/uuid"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"

	"api-server/internal/app/article/domain"
)

const expectNewUUID = ""

const (
	statusDraft = iota
	statusPublished
)

var existedUUID = uuid.NewString()

func TestCreateArticle(t *testing.T) {
	tests := []struct {
		scenario string

		inputID  string
		expected string
	}{
		{
			scenario: "When a new article is created, a new UUID must be created.",
			expected: expectNewUUID,
		},
		{
			scenario: "If the article already exists, it returns with the existing UUID.",
			inputID:  existedUUID,
			expected: existedUUID,
		},
	}

	for _, tt := range tests {
		actual := domain.NewArticle(tt.inputID, faker.Sentence(), faker.Name(), faker.URL(), faker.Paragraph(), statusDraft)
		if tt.expected == expectNewUUID {
			_, err := uuid.Parse(actual.ID())
			assert.NoError(t, err)
		} else {
			assert.Equal(t, tt.expected, actual.ID())
		}
	}
}

func TestUpdateArticleStatus(t *testing.T) {
	tests := []struct {
		scenario string
		existedArticle domain.Article
		newStatus int
		expectedStatus int
		expectedErr error
	}{
		{
			scenario: "happy path",
			existedArticle: domain.NewArticle(existedUUID, faker.Sentence(), faker.Name(), faker.URL(), faker.Paragraph(), statusDraft),
			newStatus: statusPublished,
			expectedStatus: statusPublished,
			expectedErr: nil,
		},
		{
			scenario: "if the article status is published. the status cannot be changed back to a draft",
			existedArticle: domain.NewArticle(existedUUID, faker.Sentence(), faker.Name(), faker.URL(), faker.Paragraph(), statusPublished),
			newStatus: statusDraft,
			expectedStatus: statusPublished,
			expectedErr: domain.AlreadyPublishedError,
		},
		{
			scenario: "if new article status is invalid, the status cannot be changed",
			existedArticle: domain.NewArticle(existedUUID, faker.Sentence(), faker.Name(), faker.URL(), faker.Paragraph(), statusDraft),
			newStatus: 4,
			expectedStatus: statusDraft,
			expectedErr: domain.InvalidStatusError,
		},
	}

	for _, tt := range tests {
		err := tt.existedArticle.Update(faker.Sentence(), faker.Name(), faker.URL(), faker.Paragraph(), tt.newStatus)

		assert.ErrorIs(t, err, tt.expectedErr)
		assert.Equal(t, tt.existedArticle.Status(), tt.expectedStatus)
	}
}
