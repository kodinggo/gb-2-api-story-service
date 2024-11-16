package http

import (
	"fmt"
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

	e.GET("", handlers.GetStories)
	e.GET("/:id", handlers.GetStory)
	e.POST("", handlers.CreateStory)
}

func (s *StoryHandler) GetStories(c echo.Context) error {
	var limit, offset int32
	if c.QueryParam("limit") != "" {
		parsedLimit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response{
				Status:  "Success",
				Message: err.Error(),
			})
		}
		limit = int32(parsedLimit)
	} else {
		limit = 10 // Set default limit
	}

	if c.QueryParam("offset") != "" {
		parsedOffset, err := strconv.Atoi(c.QueryParam("offset"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response{
				Status:  "Success",
				Message: err.Error(),
			})
		}
		offset = int32(parsedOffset)
	} else {
		offset = 0 // Set default offset
	}

	stories, err := s.storyUsecase.FindAll(c.Request().Context(), model.StoryFilter{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		fmt.Println("Error fetching stories:", err)
		return c.JSON(http.StatusInternalServerError, response{
			Status:  "Failed",
			Message: err.Error(),
		})
	}

	if len(stories) == 0 {
		return c.JSON(http.StatusOK, response{
			Status: "Success",
			Data:   []interface{}{},
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: "Success",
		Data:   stories,
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
		Status: "Success",
		Data:   story,
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
		Message: "Success Create Story",
	})
}
