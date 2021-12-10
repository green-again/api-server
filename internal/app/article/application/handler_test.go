package application_test

import (
	"api-server/internal/app/article/infrastructure/persistence"
	"api-server/internal/pkg/db"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"

	"api-server/internal/app/article/application"
)

func (ts HandlerTestSuite) TestArticleHandler_CreateArticle() {
	tests := []struct {
		scenario   string

		title      string
		author     string
		source     string
		body       string
		status     int

		expectErr  bool
	}{
		{
			scenario: "happy path.",

			title: faker.Sentence(),
			author: faker.Name(),
			source: faker.URL(),
			body: faker.Paragraph(),
			status: 1,

			expectErr: false,
		},
		{
			scenario: "If the repository returns an error while creating an article, handler return an error also.",

			title: faker.Sentence(),
			author: faker.Name(),
			source: faker.URL(),
			body: faker.Paragraph(),
			status: 1,

			expectErr: true,
		},
	}

	for _, tt := range tests {
		ts.dbMock.ExpectBegin()

		if tt.expectErr {
			ts.dbMock.ExpectExec("UPDATE `articles`").WillReturnError(errors.New("raise error"))
		} else {
			ts.dbMock.ExpectExec("UPDATE `articles`").WillReturnResult(sqlmock.NewResult(1, 1))
		}
		ts.dbMock.ExpectCommit()

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

type HandlerTestSuite struct {
	suite.Suite
	dbMock sqlmock.Sqlmock

	handler	application.Handler
	repo *persistence.ArticleRepository
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (ts *HandlerTestSuite) SetupTest() {
	dbSource, sqlmock := db.NewMockSQL()
	ts.dbMock = sqlmock
	ts.repo = persistence.NewArticleRepository(dbSource)
	handler := application.NewArticleHandler(ts.repo)
	ts.handler = handler
}
