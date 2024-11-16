package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/kodinggo/gb-2-api-story-service/internal/model"
)

type StoryRepo struct {
	db *sql.DB
}

func NewStoryRepo(db *sql.DB) model.IStoryRepository {
	return &StoryRepo{
		db: db,
	}
}

func (s *StoryRepo) FindAll(ctx context.Context, filter model.StoryFilter) ([]*model.Story, error) {
	query := `SELECT s.id, s.title, s.content, s.thumbnail_url, c.id AS category_id, c.name AS category_name, s.created_at, s.updated_at FROM stories AS s LEFT JOIN stories AS sc ON s.id = sc.id LEFT JOIN categories AS c ON sc.category_id = c.id LIMIT ? OFFSET ?`

	// Execute query
	res, err := s.db.QueryContext(ctx, query, filter.Limit, filter.Offset)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	storyMap := make(map[int64]*model.Story)
	for res.Next() {
		var storyId int64
		var title, content, thumbnailUrl, categoryName string
		var categoryId sql.NullInt64
		var createdAt, updatedAt time.Time

		if err := res.Scan(&storyId, &title, &content, &thumbnailUrl, &categoryId, &categoryName, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		if story, exists := storyMap[storyId]; exists {
			if categoryId.Valid && story.Category.Id == 0 {
				story.Category = model.Category{
					Id:   categoryId.Int64,
					Name: categoryName,
				}
			}
		} else {
			newStory := &model.Story{
				Id:           storyId,
				Title:        title,
				Content:      content,
				ThumbnailUrl: thumbnailUrl,
				CreatedAt:    createdAt,
				UpdatedAt:    updatedAt,
			}

			if categoryId.Valid {
				newStory.Category = model.Category{
					Id:   categoryId.Int64,
					Name: categoryName,
				}
			}

			storyMap[storyId] = newStory
		}
	}

	var stories []*model.Story
	for _, story := range storyMap {
		stories = append(stories, story)
	}

	return stories, nil

}

func (s *StoryRepo) FindById(ctx context.Context, id int64) (*model.Story, error) {
	query := `SELECT s.id, s.title, s.content, s.thumbnail_url, c.id AS category_id, c.name AS category_name, s.created_at, s.updated_at FROM stories AS s LEFT JOIN stories AS sc ON s.id = sc.id LEFT JOIN categories AS c ON sc.category_id = c.id WHERE s.id = ? LIMIT 1`

	// Execute query to fetch one story by id
	res, err := s.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var story model.Story

	if res.Next() {
		var categoryId sql.NullInt64
		var categoryName sql.NullString
		var createdAt, updatedAt time.Time

		if err := res.Scan(&story.Id, &story.Title, &story.Content, &story.ThumbnailUrl, &categoryId, &categoryName, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		if categoryId.Valid && categoryName.Valid {
			story.Category = model.Category{
				Id:   categoryId.Int64,
				Name: categoryName.String,
			}
		}

		story.CreatedAt = createdAt
		story.UpdatedAt = updatedAt
	}

	return &story, nil

}

func (s *StoryRepo) Create(ctx context.Context, story model.Story) error {
	var exists bool
	// Checks if a category with the matching id exists
	err := s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM categories WHERE id = ?)`, story.Category.Id).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("category not found")
	}

	_, err = s.db.ExecContext(ctx, `INSERT INTO stories (title, content, thumbnail_url, category_id) VALUES (?, ?, ?, ?)`,
		story.Title, story.Content, story.ThumbnailUrl, story.Category.Id)
	if err != nil {
		return err
	}

	return nil

}

func (s *StoryRepo) Delete(ctx context.Context, id int64) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM stories WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
