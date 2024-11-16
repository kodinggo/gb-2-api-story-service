package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrInvalidInput  = errors.New("invalid input")
	ErrUpdateInput   = errors.New("invalid update input")
	ErrStoryNotFound = errors.New("story not found")
)

type IStoryRepository interface {
	FindAll(ctx context.Context, filter StoryFilter) ([]*Story, error)
	FindById(ctx context.Context, id int64) (*Story, error)
	Create(ctx context.Context, story Story) error
	Update(ctx context.Context, story Story) error
	Delete(ctx context.Context, id int64) error
}

type IStoryUsecase interface {
	FindAll(ctx context.Context, filter StoryFilter) ([]*Story, error)
	FindById(ctx context.Context, id int64) (*Story, error)
	Create(ctx context.Context, in CreateStoryInput) error
	Update(ctx context.Context, id int64, in UpdateStoryInput) error
	Delete(ctx context.Context, id int64) error
}

type Story struct {
	Id           int64        `json:"id"`
	Title        string       `json:"title"`
	Content      string       `json:"content"`
	ThumbnailUrl string       `json:"thumbnail_url"`
	Category     Category     `json:"category"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at,omitempty"`
}

type StoryResponse struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	Category     Category  `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type StoryFilter struct {
	Offset int32
	Limit  int32
}

type CreateStoryInput struct {
	Title        string `json:"title" validate:"required,min=3,max=255"`
	Content      string `json:"content" validate:"required"`
	ThumbnailUrl string `json:"thumbnail_url" validate:"required"`
	CategoryId   int    `json:"category_id" validate:"required"`
}

type UpdateStoryInput struct {
	Title        string `json:"title" validate:"required,min=3,max=255"`
	Content      string `json:"content" validate:"required"`
	ThumbnailUrl string `json:"thumbnail_url" validate:"required"`
	CategoryId   int    `json:"category_id" validate:"required"`
}
