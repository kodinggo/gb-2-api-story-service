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

func NewStoryHandler(e *echo.Group, us model.IStoryUsecase) {
	handlers := &StoryHandler{
		storyUsecase: us,
	}

	stories := e.Group("/stories")

	stories.GET("", handlers.GetStories)
	stories.GET("/:id", handlers.GetStory)
	stories.POST("", handlers.CreateStory)
}

func (s *StoryHandler) GetStories(c echo.Context) error {

	reqLimit := c.QueryParam("limit")
	reqOffset := c.QueryParam("offset")

	var limit, offset int32
	if reqLimit == "" {
		limit = 10 //default limit
	}
	if reqOffset == "" {
		offset = 0 //default offset
	}

	stories, err := s.storyUsecase.FindAll(c.Request().Context(), model.StoryFilter{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    stories,
	})
}

func (s *StoryHandler) GetStory(c echo.Context) error {

	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	story, err := s.storyUsecase.FindById(c.Request().Context(), int64(parsedId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    story,
	})
}

func (s *StoryHandler) CreateStory(c echo.Context) error {

	var input model.CreateStoryInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := s.storyUsecase.Create(c.Request().Context(), input); err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response{
		Status:  http.StatusCreated,
		Message: "Success",
	})
}
