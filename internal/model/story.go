package model

import (
	"context"
	"database/sql"
	"time"
)

const (
	DefaultLimit = 20
	DefaultPage  = 0
)

type IStoryRepository interface {
	FindAll(ctx context.Context, filter FindAllParam) ([]*Story, error)
	FindById(ctx context.Context, id int64) (*Story, error)
	Create(ctx context.Context, story Story) error
	Update(ctx context.Context, story Story) error
	Delete(ctx context.Context, id int64) error
}

type IStoryUsecase interface {
	FindAll(ctx context.Context, filter FindAllParam) ([]*Story, error)
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
	DeletedAt    sql.NullTime `json:"-"`
}

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type FindAllParam struct {
	Limit int64
	Page  int64
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
