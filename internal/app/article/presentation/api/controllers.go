package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"api-server/internal/app/article/application"
	httpapi "api-server/internal/pkg/http"
)

type ArticleController struct {
	httpapi.RequestBinder
	handler   application.Handler
	validator httpapi.Validator
}

// GetArticle godoc
// @Summary      Show an article details.
// @Description  GetArticle finds and returns one Article by request ID.
// @Tags         Articles
// @Param        id   path      string  true  "Article ID"
// @Produce      json
// @Success      200  {object}  Article
// @Failure      400  {object}  http.ErrorResponse
// @Failure      404  {object}  http.ErrorResponse
// @Failure      500  {object}  http.ErrorResponse
// @Router       /http/v1/articles/{id} [get]
func (con *ArticleController) GetArticle(c echo.Context) error {
	id := c.Param("id")
	if err := con.validator.Validate(id); err != nil {
		return con.handleErrorResponse(c, err)
	}

	article, err := con.handler.GetArticleByID(id)
	if err != nil {
		return con.handleErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, MapArticleResponse(article))
}

// PostArticle godoc
// @Summary      Create new article.
// @Description  PostArticle creates a new article and stores it in the data source.
// @Tags         Articles
// @Param        article   body Article true  "Article ingredient"
// @Produce      json
// @Success      200  {object}  Article
// @Failure      400  {object}  http.ErrorResponse
// @Failure      404  {object}  http.ErrorResponse
// @Failure      500  {object}  http.ErrorResponse
// @Router       /http/v1/articles [post]
func (con *ArticleController) PostArticle(c echo.Context) error {
	article := new(Article)
	if err := con.Bind(c, article); err != nil {
		return con.handleErrorResponse(c, err)
	}

	res, err := con.handler.CreateArticle(article.Title, article.Author, article.Source, article.Body, article.Status)
	if err != nil {
		return con.handleErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, MapArticleResponse(res))
}

func (con *ArticleController) handleErrorResponse(c echo.Context, err error) error {
	switch err.(type) {
	case httpapi.InvalidRequestError:
		return c.JSON(http.StatusBadRequest, httpapi.NewErrorResponse(InvalidRequest, "invalid request.", err.Error()))
	case application.NotFoundError:
		return c.JSON(http.StatusNotFound, httpapi.NewErrorResponse(NotFound, "resource not found.", err.Error()))
	case application.UnknownError:
		return c.JSON(http.StatusInternalServerError, httpapi.NewErrorResponse(Unknown, "internal server error.", err.Error()))
	}
	return c.JSON(http.StatusInternalServerError, httpapi.NewErrorResponse(Unknown, "internal server error.", err.Error()))
}

func NewController(handler application.Handler, validator httpapi.Validator) *ArticleController {
	return &ArticleController{handler: handler, validator: validator}
}
