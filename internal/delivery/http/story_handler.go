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
	e.PUT("/:id", handlers.UpdateStory)
	e.DELETE("/:id", handlers.DeleteStory)
}

func (s *StoryHandler) GetStories(c echo.Context) error {
	var limit, offset int64

	if c.QueryParam("limit") != "" {
		parsedLimit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid limit value")
		}
		limit = int64(parsedLimit)
	} else {
		limit = 10 // Default limit
	}

	if c.QueryParam("offset") != "" {
		parsedOffset, err := strconv.Atoi(c.QueryParam("offset"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid offset value")
		}
		offset = int64(parsedOffset)
	} else {
		offset = 0 // Default offset
	}

	stories, err := s.storyUsecase.FindAll(c.Request().Context(), model.StoryFilter{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		fmt.Println("Error fetching stories:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching stories")
	}

	var storyResponses []model.StoryResponse
	for _, story := range stories {
		if story.DeletedAt.Valid {
			continue
		}

		storyResponses = append(storyResponses, model.StoryResponse{
			Id:           story.Id,
			Title:        story.Title,
			Content:      story.Content,
			ThumbnailUrl: story.ThumbnailUrl,
			Category:     story.Category,
			CreatedAt:    story.CreatedAt,
			UpdatedAt:    story.UpdatedAt,
		})
	}

	// Handle empty stories
	if len(storyResponses) == 0 {
		return c.JSON(http.StatusOK, response{
			Status: "Success",
			Data:   []interface{}{},
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: "success",
		Data:   storyResponses,
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
		if err.Error() == model.ErrStoryNotFound.Error() {
			return c.JSON(http.StatusNotFound, response{
				Status: "Sorry, story not found",
			})
		}

		return c.JSON(http.StatusNoContent, response{
			Status: "Story Not Found",
		})
	}

	storyResponse := model.StoryResponse{
		Id:           story.Id,
		Title:        story.Title,
		Content:      story.Content,
		ThumbnailUrl: story.ThumbnailUrl,
		Category:     story.Category,
		CreatedAt:    story.CreatedAt,
		UpdatedAt:    story.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response{
		Status: "Success",
		Data:   storyResponse,
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

func (s *StoryHandler) UpdateStory(c echo.Context) error {
	storyIdParam := c.Param("id")

	storyId, err := strconv.ParseInt(storyIdParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: "Invalid story ID",
		})
	}

	var input model.UpdateStoryInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := s.storyUsecase.Update(c.Request().Context(), storyId, input); err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
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
		return echo.NewHTTPError(500, err.Error())
	}

	if err := s.storyUsecase.Delete(c.Request().Context(), int64(parsedId)); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(http.StatusNoContent, response{
		Status: http.StatusNoContent,
	})
}
