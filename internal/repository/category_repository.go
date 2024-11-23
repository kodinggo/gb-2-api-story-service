package repository

import (
	"context"
	"database/sql"

	"github.com/kodinggo/gb-2-api-story-service/internal/model"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) model.ICategoryRepository {
	return &CategoryRepo{
		db: db,
	}
}

func (c *CategoryRepo) FindAll(ctx context.Context) ([]*model.Categories, error) {
	res, err := c.db.QueryContext(ctx, `SELECT id, name, created_at, updated_at FROM categories`)
	if err != nil {
		return nil, err
	}

	var categories []*model.Categories
	for res.Next() {
		var category model.Categories
		if err := res.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (c *CategoryRepo) FindById(ctx context.Context, id int64) (*model.Categories, error) {
	res, err := c.db.QueryContext(ctx, `SELECT id, name, created_at, updated_at FROM categories WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	var category model.Categories
	for res.Next() {
		if err := res.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}
	}

	return &category, nil
}

func (c *CategoryRepo) Create(ctx context.Context, category model.Categories) error {
	_, err := c.db.ExecContext(ctx, `INSERT INTO categories (name) VALUES (?)`, category.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepo) Update(ctx context.Context, category model.Categories) error {
	_, err := c.db.ExecContext(ctx, `UPDATE categories SET name = ? WHERE id = ?`, category.Name, category.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepo) Delete(ctx context.Context, id int64) error {
	_, err := c.db.ExecContext(ctx, `DELETE FROM categories WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
