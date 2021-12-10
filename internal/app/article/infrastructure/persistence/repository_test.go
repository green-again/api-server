package persistence_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/google/uuid"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"

	"api-server/internal/app/article/domain"
	"api-server/internal/app/article/infrastructure/persistence"
	"api-server/internal/pkg/db"
)


func (ts *RepositoryTestSuite) TestGetArticleByID() {
	requestArticleID, _ := uuid.NewUUID()
	articleModelID, _ := requestArticleID.MarshalBinary()
	invalidArticleModelID := []byte{'1'}

	tests := []struct{
		scenario string
		requestID string
		expectErr  bool
		expectRecordNotFound bool
		mockQueryResult []byte
	}{
		{
			scenario:        "happy path",
			requestID:       requestArticleID.String(),
			mockQueryResult: articleModelID,
		},
		{
			scenario: "marshaling failed",
			requestID: "not uuid format",
			expectErr: true,
		},
		{
			scenario: "not found record",
			expectErr: true,
			expectRecordNotFound: true,
		},
		{
			scenario: "domain mapping error",
			mockQueryResult: invalidArticleModelID,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		if tt.mockQueryResult != nil {
			ts.dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `articles`")).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).
					AddRow(tt.mockQueryResult))
		}
		if tt.expectRecordNotFound {
			ts.dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `articles`")).
				WillReturnError(errors.New("not found error"))
		}
		actual, err := ts.Repository.GetArticleByID(tt.requestID)

		if tt.expectErr {
			ts.Error(err)
		} else {
			ts.NoError(err)
			ts.Equal(tt.requestID, actual.ID())
		}
	}
}

func (ts *RepositoryTestSuite) TestSaveArticle() {
	tests := []struct{
		scenario string
		inputArticle domain.Article
		expectErr  bool
		raiseDatabaseErr bool
	}{
		{
			scenario: "happy path",
			inputArticle: domain.NewArticle(
				"",
				faker.Sentence(),
				faker.Name(),
				faker.URL(),
				faker.Paragraph(),
				1,
			),
		},
		{
			scenario: "invalid uuid error",
			inputArticle: domain.NewArticle(
				"123",
				faker.Sentence(),
				faker.Name(),
				faker.URL(),
				faker.Paragraph(),
				1,
			),
			expectErr: true,
		},
		{
			scenario: "error occurred in the database",
			inputArticle: domain.NewArticle(
				"",
				faker.Sentence(),
				faker.Name(),
				faker.URL(),
				faker.Paragraph(),
				1,
			),
			raiseDatabaseErr: true,
		},
	}

	for _, tt := range tests {
		ts.dbMock.ExpectBegin()
		if tt.raiseDatabaseErr {
			ts.dbMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
				WillReturnError(errors.New("test error"))
		} else {
			ts.dbMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}
		ts.dbMock.ExpectCommit()

		err := ts.Repository.SaveArticle(&tt.inputArticle)
		if tt.expectErr {
			ts.Error(err)
		} else {
			ts.NoError(err)
		}
	}
}

type RepositoryTestSuite struct {
	suite.Suite
	dbMock sqlmock.Sqlmock

	Repository *persistence.ArticleRepository
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (ts *RepositoryTestSuite) SetupTest() {
	orm, sqlMock := db.NewMockSQL()
	ts.dbMock = sqlMock
	ts.Repository = persistence.NewArticleRepository(orm)
}
