package http

import (
	"net/http"
	"strconv"

	"github.com/kodinggo/gb-2-api-story-service/internal/model"
	"github.com/labstack/echo/v4"
)

type StoryHandler struct {
	storyUsecase model.IStoryUsecase
}

func NewStoryHandler(e *echo.Echo, us model.IStoryUsecase) {
	handlers := &StoryHandler{
		storyUsecase: us,
	}

	routeStories := e.Group("/v1/stories")
	routeStories.GET("", handlers.GetStories)
	routeStories.GET("/:id", handlers.GetStory)
	routeStories.POST("", handlers.CreateStory)
	routeStories.PUT("/:id", handlers.UpdateStory)
	routeStories.DELETE("/:id", handlers.DeleteStory)
}

func (s *StoryHandler) GetStories(c echo.Context) error {
	var param model.FindAllParam

	if limitParam := c.QueryParam("limit"); limitParam != "" {
		parsedLimit, err := strconv.Atoi(limitParam)
		if err != nil || parsedLimit <= 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid limit value")
		}
		param.Limit = int64(parsedLimit)
	}

	if pageParam := c.QueryParam("page"); pageParam != "" {
		parsedPage, err := strconv.Atoi(pageParam)
		if err != nil || parsedPage <= 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid page value")
		}
		param.Page = int64(parsedPage)
	}

	stories, err := s.storyUsecase.FindAll(c.Request().Context(), param)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching stories")
	}

	return c.JSON(http.StatusOK, response{
		Status: "success",
		Data:   stories,
	})
}

func (s *StoryHandler) GetStory(c echo.Context) error {
	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid story ID")
	}

	story, err := s.storyUsecase.FindById(c.Request().Context(), int64(parsedId))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Sorry, story not found!")
	}

	return c.JSON(http.StatusOK, response{
		Status: "success",
		Data:   story,
	})
}

func (s *StoryHandler) CreateStory(c echo.Context) error {
	var input model.CreateStoryInput

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.storyUsecase.Create(c.Request().Context(), input); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create story")
	}

	return c.JSON(http.StatusCreated, response{
		Status:  "success",
		Message: "Success Create Story",
	})
}

func (s *StoryHandler) UpdateStory(c echo.Context) error {
	storyIdParam := c.Param("id")

	storyId, err := strconv.ParseInt(storyIdParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid story ID")
	}

	var input model.UpdateStoryInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	if err := s.storyUsecase.Update(c.Request().Context(), storyId, input); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update story")
	}

	return c.JSON(http.StatusOK, response{
		Status:  "success",
		Message: "Success Update Story",
	})
}

func (s *StoryHandler) DeleteStory(c echo.Context) error {
	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := s.storyUsecase.Delete(c.Request().Context(), int64(parsedId)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, response{
		Status: http.StatusNoContent,
	})
}
