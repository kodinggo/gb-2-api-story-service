package usecase

import (
	"context"

	"github.com/kodinggo/gb-2-api-story-service/internal/model"
	"github.com/sirupsen/logrus"
)

type CategoryUsecase struct {
	CategoryRepo model.ICategoryRepository
}

func NewCategoryUsecase(categoryRepo model.ICategoryRepository) model.ICategoryUsecase {
	return &CategoryUsecase{
		CategoryRepo: categoryRepo,
	}
}

func (c *CategoryUsecase) FindAll(ctx context.Context) ([]*model.Categories, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	categories, err := c.CategoryRepo.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return categories, nil
}

func (c *CategoryUsecase) FindById(ctx context.Context, id int64) (*model.Categories, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	category, err := c.CategoryRepo.FindById(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return category, nil
}

func (c *CategoryUsecase) Create(ctx context.Context, category model.Categories) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"name": category.Name,
	})

	err := v.StructCtx(ctx, category)
	if err != nil {
		log.Error(err)
		return err
	}

	newCategory := model.Categories{
		Id:   category.Id,
		Name: category.Name,
	}

	err = c.CategoryRepo.Create(ctx, newCategory)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (c *CategoryUsecase) Update(ctx context.Context, category model.Categories) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"name": category.Name,
	})

	err := v.StructCtx(ctx, category)
	if err != nil {
		log.Error(err)
		return err
	}

	newCategory := model.Categories{
		Id:   category.Id,
		Name: category.Name,
	}

	err = c.CategoryRepo.Update(ctx, newCategory)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (c *CategoryUsecase) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := c.CategoryRepo.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
