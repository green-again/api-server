package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"api-server/internal/app/article/application"
	"api-server/internal/pkg/api/errors"
	"api-server/internal/pkg/api/validators"
)

type ArticleController struct {
	handler   application.Handler
	validator validators.Validator
}

// GetArticle godoc
// @Summary      Show an article details.
// @Description  GetArticle finds and returns one Article by request ID.
// @Tags         Articles
// @Param        id   path      string  true  "Article ID"
// @Produce      json
// @Success      200  {object}  Article
// @Failure      400  {object}  errors.ErrorResponse
// @Failure      404  {object}  errors.ErrorResponse
// @Failure      500  {object}  errors.ErrorResponse
// @Router       /api/v1/articles/{id} [get]
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

func (con *ArticleController) handleErrorResponse(c echo.Context, err error) error {
	switch err.(type) {
	case InvalidRequestError:
		return c.JSON(http.StatusBadRequest, errors.NewErrorResponse(InvalidRequest, "invalid request.", err.Error()))
	case application.NotFoundError:
		return c.JSON(http.StatusNotFound, errors.NewErrorResponse(NotFound, "resource not found.", err.Error()))
	case application.UnknownError:
		return c.JSON(http.StatusInternalServerError, errors.NewErrorResponse(Unknown, "internal server error.", err.Error()))
	}
	return c.JSON(http.StatusInternalServerError, err.Error())
}

func NewController(handler application.Handler, validator validators.Validator) *ArticleController {
	return &ArticleController{handler: handler, validator: validator}
}
