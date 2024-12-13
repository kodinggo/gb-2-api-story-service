package model

import (
	"context"
	"database/sql"
	"time"
)

const (
	DefaultLimit = 20
	DefaultPage  = 1
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
	Comments     []*Comment   `json:"comments"`
	Category     Category     `json:"category"`
	Author       Account      `json:"author"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"-"`
}
type AccountUsecase interface {
	FindByID(id int64, db string) (*Account, error)
}

type Account struct {
	Id         int64  `json:"id"`
	Fullname   string `json:"fullname"`
	SortBio    string `json:"sort_bio"`
	Gender     string `json:"gender"`
	PictureUrl string `json:"picture_url"`
	Username   string `json:"username"`
	Email      string `json:"email"`
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

type Comment struct {
	ID        int64      `json:"id"`
	Comment   string     `json:"comment" validate:"required"`
	StoryID   int64      `json:"story_id,omitempty" validate:"required" `
	UserID    int64      `json:"user_id,omitempty" validate:"required" `
	Author    Author     `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
type Author struct {
	ID         int64  `json:"id"`
	Fullname   string `json:"fullname"`
	SortBio    string `json:"sort_bio"`
	Gender     string `json:"gender"`
	PictureURL string `json:"picture_url"`
	Username   string `json:"username"`
	Email      string `json:"email"`
}
