package api_test

import (
	"api-server/internal/app/article/domain"
	"api-server/internal/app/article/presentation/api"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
)

func (ts *ControllerTestSuite) TestGetArticles() {
	tests := []struct {
		scenario string
		description string
		pathParams string
		mockArticle domain.Article
		expectedResponse string
		expectedCode int
	}{
		{
			scenario: "happy path",
			pathParams: MockArticles()[0].ID,
			mockArticle: MockArticles()[0],
			expectedResponse: MockArticlesResponse()[0],
			expectedCode: http.StatusOK,
		},
		{
			scenario: "If the request has the wrong parameter, you should get a 400 error.",
			pathParams: MockArticles()[0].ID,
			mockArticle: MockArticles()[0],
			expectedResponse: MockArticlesResponse()[0],
			expectedCode: http.StatusBadRequest,
		},
		{
			scenario: "If the request is an article ID that does not exist, a 404 error should be returned.",
			pathParams: MockArticles()[0].ID,
			mockArticle: MockArticles()[0],
			expectedResponse: MockArticlesResponse()[0],
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/articles/:id")
		c.SetParamNames("id")
		c.SetParamValues(tt.pathParams)

		err := ts.controller.GetArticles(c)

		ts.NoError(err)
		ts.Equal(tt.expectedCode, rec.Code)
		ts.Equal(tt.expectedResponse, rec.Body.String())
	}
}

type ControllerTestSuite struct {
	suite.Suite
	controller api.ArticleController
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}


var mockArticles []domain.Article
func MockArticles() []domain.Article {
	if mockArticles != nil {
		return mockArticles
	}

	for i := 0; i < 3; i++ {
		mockArticles = append(mockArticles, domain.Article{
			ID:            uuid.NewString(),
			Title:         faker.Sentence(),
			Author:        faker.Name(),
			Source:        faker.URL(),
			Body:          faker.Paragraph(),
			PublishedDate: time.Now(),
		})
	}
	return mockArticles
}

var mockArticlesResponse []string
func MockArticlesResponse() []string {
	if mockArticlesResponse != nil {
		return mockArticlesResponse
	}

	for _, article := range MockArticles() {
		res := fmt.Sprintf(`{"article": {"id": "%s","title": "%s","author": "%s","source": "%s","publishedDate": "%s","body": "%s"}}`,
			article.ID, article.Title, article.Author, article.Source, article.PublishedDate, article.Body)

		mockArticlesResponse = append(mockArticlesResponse, res)
	}
	return mockArticlesResponse
}
