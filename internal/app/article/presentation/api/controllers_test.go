package api_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"

	"api-server/internal/app/article/domain"
	"api-server/internal/app/article/infrastructure/persistence"
	"api-server/internal/app/article/presentation/api"
	httpapi "api-server/internal/pkg/http"
)

func (ts *ControllerTestSuite) TestPostArticles() {
	tests := []struct {
		scenario     string
		requestBody  string
		mockFunc     func() (*domain.Article, error)
		expectedCode int
	}{
		{
			scenario: "Happy path",
			requestBody: fmt.Sprintf(
				`{"title":"%s","body":"%s","author":"%s","source":"%s","status":%d}`,
				faker.Sentence(),
				faker.Paragraph(),
				faker.Name(),
				faker.URL(),
				1,
			),
			mockFunc: func() (*domain.Article, error) {
				return &MockArticles()[0], nil
			},
			expectedCode: http.StatusCreated,
		},
		{
			scenario: "If parameter is missing, it should returns 400 error.",
			requestBody: fmt.Sprintf(
				`{"body":"%s","author":"%s","source":"%s","status":%d}`,
				faker.Paragraph(),
				faker.Name(),
				faker.URL(),
				1,
			),
			expectedCode: http.StatusBadRequest,
		},
		{
			scenario: "If an unknown error occurs during processing, an error 500 should be returned.",
			requestBody: fmt.Sprintf(
				`{"title":"%s","body":"%s","author":"%s","source":"%s","status":%d}`,
				faker.Sentence(),
				faker.Paragraph(),
				faker.Name(),
				faker.URL(),
				1,
			),
			mockFunc: func() (*domain.Article, error) {
				return nil, errors.New("unknown error")
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		e := echo.New()
		e.Validator = httpapi.NewRequestValidator()
		req := httptest.NewRequest(http.MethodPost, "/http/v1/articles", strings.NewReader(tt.requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if tt.mockFunc != nil {
			ts.handler.On(
				"CreateArticle",
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("int"),
			).Return(tt.mockFunc()).Once()
		}
		err := ts.controller.PostArticle(c)
		ts.NoError(err)
		ts.Equal(tt.expectedCode, rec.Code)
	}
}

func (ts *ControllerTestSuite) TestUpdateArticles() {
	tests := []struct {
		scenario     string
		pathParam string
		requestBody  string
		mockFunc     func() (*domain.Article, error)
		expectedCode int
	}{
		{
			scenario: "Happy path",
			pathParam: MockArticles()[0].ID(),
			requestBody: fmt.Sprintf(
				`{"title":"%s","body":"%s","author":"%s","source":"%s","status":%d}`,
				faker.Sentence(),
				faker.Paragraph(),
				faker.Name(),
				faker.URL(),
				2,
			),
			mockFunc: func() (*domain.Article, error) {
				return &MockArticles()[0], nil
			},
			expectedCode: http.StatusOK,
		},
		{
			scenario: "If parameter is missing, it should returns 400 error.",
			pathParam: MockArticles()[1].ID(),
			requestBody: fmt.Sprintf(
				`{"body":"%s","author":"%s","source":"%s","status":%d}`,
				faker.Paragraph(),
				faker.Name(),
				faker.URL(),
				1,
			),
			expectedCode: http.StatusBadRequest,
		},
		{
			scenario: "If an unknown error occurs during processing, an error 500 should be returned.",
			pathParam: MockArticles()[2].ID(),
			requestBody: fmt.Sprintf(
				`{"title":"%s","body":"%s","author":"%s","source":"%s","status":%d}`,
				faker.Sentence(),
				faker.Paragraph(),
				faker.Name(),
				faker.URL(),
				1,
			),
			mockFunc: func() (*domain.Article, error) {
				return nil, errors.New("unknown error")
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		e := echo.New()
		e.Validator = httpapi.NewRequestValidator()
		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/http/v1/articles/%s", tt.pathParam), strings.NewReader(tt.requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if tt.mockFunc != nil {
			ts.handler.On(
				"UpdateArticle",
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
				mock.AnythingOfType("int"),
			).Return(tt.mockFunc()).Once()
		}
		err := ts.controller.UpdateArticle(c)
		ts.NoError(err)
		ts.Equal(tt.expectedCode, rec.Code)

		ts.handler.AssertExpectations(ts.T())
	}
}

func (ts *ControllerTestSuite) TestGetArticles() {
	tests := []struct {
		scenario         string
		pathParams       string
		mockFunc         func() (*domain.Article, error)
		expectedResponse api.Article
		expectedCode     int
	}{
		{
			scenario:   "happy path",
			pathParams: MockArticles()[0].ID(),
			mockFunc: func() (*domain.Article, error) {
				return &MockArticles()[0], nil
			},
			expectedResponse: MockArticlesResponse()[0],
			expectedCode:     http.StatusOK,
		},
		{
			scenario:     "If the request has the wrong parameter, you should get a 400 error.",
			pathParams:   "1",
			expectedCode: http.StatusBadRequest,
		},
		{
			scenario:   "If the request is an article ID that does not exist, a 404 error should be returned.",
			pathParams: MockArticles()[1].ID(),
			mockFunc: func() (*domain.Article, error) {
				return nil, persistence.NotFoundError
			},
			expectedCode: http.StatusNotFound,
		},
		{
			scenario:   "If an unknown error occurs during processing, an error 500 should be returned.",
			pathParams: MockArticles()[2].ID(),
			mockFunc: func() (*domain.Article, error) {
				return nil, errors.New("unknown error")
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		e := echo.New()
		e.Validator = httpapi.NewRequestValidator()
		req := httptest.NewRequest(http.MethodGet, "/http/v1/", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/articles/:id")
		c.SetParamNames("id")
		c.SetParamValues(tt.pathParams)

		if tt.mockFunc != nil {
			ts.handler.On("GetArticleByID", tt.pathParams).Return(tt.mockFunc()).Once()
		}
		err := ts.controller.GetArticle(c)

		ts.NoError(err)
		ts.Equal(tt.expectedCode, rec.Code)

		if tt.expectedCode == http.StatusOK {
			var actualResp api.Article
			json.Unmarshal(rec.Body.Bytes(), &actualResp)

			ts.Equal(tt.expectedResponse, actualResp)
		}
	}
}

type ControllerTestSuite struct {
	suite.Suite
	handler    *mockHandler
	controller *api.ArticleController
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}

func (ts *ControllerTestSuite) SetupTest() {
	ts.handler = new(mockHandler)
	ts.controller = api.NewController(ts.handler, api.NewGetArticleRequestValidator())
}

var mockArticles []domain.Article

func MockArticles() []domain.Article {
	if mockArticles != nil {
		return mockArticles
	}

	for i := 0; i < 3; i++ {
		mockArticles = append(mockArticles, domain.NewArticle(uuid.NewString(), faker.Sentence(), faker.Name(), faker.URL(), faker.Paragraph(), 1))
	}
	return mockArticles
}

var mockArticlesResponse []api.Article

func MockArticlesResponse() []api.Article {
	if mockArticlesResponse != nil {
		return mockArticlesResponse
	}

	for _, article := range MockArticles() {
		mockArticlesResponse = append(mockArticlesResponse, api.MapArticleResponse(&article))
	}
	return mockArticlesResponse
}

type mockHandler struct {
	mock.Mock
}

func (h *mockHandler) GetArticleByID(id string) (*domain.Article, error) {
	ret := h.Called(id)
	return ret.Get(0).(*domain.Article), ret.Error(1)
}

func (h *mockHandler) CreateArticle(title, author, source, body string, status int) (*domain.Article, error) {
	ret := h.Called(title, author, source, body, status)
	return ret.Get(0).(*domain.Article), ret.Error(1)
}

func (h *mockHandler) UpdateArticle(id, title, author, source, body string, status int) (*domain.Article, error) {
	ret := h.Called(id, title, author, source, body, status)
	return ret.Get(0).(*domain.Article), ret.Error(1)
}
