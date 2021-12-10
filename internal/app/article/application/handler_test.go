package application_test

import (
	"errors"

	"github.com/google/uuid"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"

	"api-server/internal/app/article/application"
	"api-server/internal/app/article/domain"
)

func (ts HandlerTestSuite) TestArticleHandler_CreateArticle() {
	tests := []struct {
		scenario string

		title  string
		author string
		source string
		body   string
		status int

		mockSaveArticleResult error
		expectErr             bool
	}{
		{
			scenario: "happy path.",

			title:  faker.Sentence(),
			author: faker.Name(),
			source: faker.URL(),
			body:   faker.Paragraph(),
			status: 1,

			mockSaveArticleResult: nil,
			expectErr:             false,
		},
		{
			scenario: "If the repository returns an error while creating an article, handler return an error also.",

			title:  faker.Sentence(),
			author: faker.Name(),
			source: faker.URL(),
			body:   faker.Paragraph(),
			status: 1,

			mockSaveArticleResult: errors.New("test error"),
			expectErr:             true,
		},
	}

	for _, tt := range tests {
		ts.repo.On("SaveArticle", mock.AnythingOfType("*domain.Article")).
			Return(tt.mockSaveArticleResult).Once()

		actual, err := ts.handler.CreateArticle(tt.title, tt.author, tt.source, tt.body, tt.status)

		if tt.expectErr {
			ts.Error(err)
		} else {
			ts.NoError(err)

			_, err = uuid.Parse(actual.ID())
			ts.NoError(err)
			ts.Equal(tt.title, actual.Title())
			ts.Equal(tt.author, actual.Author())
			ts.Equal(tt.source, actual.Source())
			ts.Equal(tt.body, actual.Body())
			ts.Equal(tt.status, actual.Status())
		}
	}
}

func (ts *HandlerTestSuite) TestArticleHandler_GetArticleByID() {
	tests := []struct {
		scenario string

		id string

		mockGetArticleMethod func(string) (*domain.Article, error)
		expectErr            bool
	}{
		{
			scenario: "happy path.",
			id:       uuid.NewString(),

			mockGetArticleMethod: func(articleID string) (*domain.Article, error) {
				ret := domain.NewArticle(
					articleID,
					faker.Sentence(),
					faker.Sentence(),
					faker.Sentence(),
					faker.Sentence(),
					1,
				)
				return &ret, nil
			},
			expectErr: false,
		},
		{
			scenario: "If the repository returns an error while creating an article, handler return an error also.",
			id:       uuid.NewString(),

			mockGetArticleMethod: func(articleID string) (*domain.Article, error) {
				return nil, errors.New("test error")
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		ts.repo.On("GetArticleByID", tt.id).Return(tt.mockGetArticleMethod(tt.id)).Once()
		actual, err := ts.handler.GetArticleByID(tt.id)
		if tt.expectErr {
			ts.Error(err)
		} else {
			ts.Equal(tt.id, actual.ID())
		}
	}
}

type HandlerTestSuite struct {
	suite.Suite

	handler application.Handler
	repo    *mockRepository
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (ts *HandlerTestSuite) SetupTest() {
	ts.repo = new(mockRepository)
	handler := application.NewArticleHandler(ts.repo)
	ts.handler = handler
}

type mockRepository struct {
	mock.Mock
}

func (r *mockRepository) GetArticleByID(articleID string) (*domain.Article, error) {
	ret := r.Called(articleID)
	return ret.Get(0).(*domain.Article), ret.Error(1)
}

func (r *mockRepository) SaveArticle(article *domain.Article) error {
	ret := r.Called(article)
	return ret.Error(0)
}
