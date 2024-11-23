package model

import (
	"context"
	"time"
)

type ICategoryUsecase interface {
	FindAll(ctx context.Context) ([]*Categories, error)
	FindById(ctx context.Context, id int64) (*Categories, error)
	Create(ctx context.Context, category Categories) error
	Update(ctx context.Context, category Categories) error
	Delete(ctx context.Context, id int64) error
}

type ICategoryRepository interface {
	FindAll(ctx context.Context) ([]*Categories, error)
	FindById(ctx context.Context, id int64) (*Categories, error)
	Create(ctx context.Context, category Categories) error
	Update(ctx context.Context, category Categories) error
	Delete(ctx context.Context, id int64) error
}

type Categories struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCategoryInput struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryInput struct {
	Id       int64      `json:"id"`
	Name     string     `json:"name" validate:"required"`
	UpdateAt *time.Time `json:"updated_at"`
}
