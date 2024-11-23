package http

import (
	"net/http"
	"strconv"

	"github.com/kodinggo/gb-2-api-story-service/internal/model"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryUsecase model.ICategoryUsecase
}

func NewCategoryHandler(e *echo.Echo, us model.ICategoryUsecase) {
	handlers := &CategoryHandler{
		categoryUsecase: us,
	}

	routeCategories := e.Group("/v1/categories")
	routeCategories.GET("", handlers.GetCategories)
	routeCategories.GET("/:id", handlers.GetCategory)
	routeCategories.POST("", handlers.CreateCategory)
	routeCategories.PUT("/:id", handlers.UpdateCategory)
	routeCategories.DELETE("/:id", handlers.DeleteCategory)

}

func (ch *CategoryHandler) GetCategories(c echo.Context) error {
	categories, err := ch.categoryUsecase.FindAll(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: "success",
		Data:   categories,
	})
}

func (s *CategoryHandler) GetCategory(c echo.Context) error {
	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	category, err := s.categoryUsecase.FindById(c.Request().Context(), int64(parsedId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "category not found")
	}

	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    category,
	})
}

func (s *CategoryHandler) CreateCategory(c echo.Context) error {
	var input model.CreateCategoryInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	if err := s.categoryUsecase.Create(c.Request().Context(), model.Categories{Name: input.Name}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create category")
	}

	return c.JSON(http.StatusCreated, response{
		Status: "success",
	})
}

func (s *CategoryHandler) UpdateCategory(c echo.Context) error {
	categoryId := c.Param("id")
	parsedId, err := strconv.Atoi(categoryId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}

	var in model.UpdateCategoryInput
	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	if parsedId > 0 {
		in.Id = int64(parsedId)
	}

	if err := s.categoryUsecase.Update(c.Request().Context(), model.Categories{Id: in.Id, Name: in.Name}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update category")
	}

	return c.JSON(http.StatusOK, response{
		Status: "success",
	})
}

func (s *CategoryHandler) DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := s.categoryUsecase.Delete(c.Request().Context(), int64(parsedId)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, response{
		Status: http.StatusNoContent,
	})
}
