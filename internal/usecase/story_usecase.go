package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kodinggo/gb-2-api-comment-service/pb/comment_service"
	"github.com/kodinggo/gb-2-api-story-service/internal/model"
	"github.com/sirupsen/logrus"
)

type StoryUsecase struct {
	storyRepo       model.IStoryRepository
	categoryUsecase model.ICategoryRepository
	grpcCommentClient comment_service.CommentServiceClient
}

var v = validator.New()

func NewStoryUsecase(
	storyRepo model.IStoryRepository,
	grpcCommentClient comment_service.CommentServiceClient,
	categoryUsecase model.ICategoryUsecase,
) model.IStoryUsecase {
	return &StoryUsecase{
		storyRepo:       storyRepo,
		categoryUsecase: categoryUsecase,
		grpcCommentClient:grpcCommentClient,
	}
}

func (s *StoryUsecase) FindAll(ctx context.Context, filter model.FindAllParam) ( []*model.Story,  error) {
	if filter.Limit <= 0 {
		filter.Limit = model.DefaultLimit
	}

	if filter.Page <= 0 {
		filter.Page = model.DefaultPage
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx":   ctx,
		"limit": filter.Limit,
		"page":  filter.Page,
	})

	storyFilter := model.FindAllParam{
		Limit: filter.Limit,
		Page:  filter.Page,
	}

	stories, err := s.storyRepo.FindAll(ctx, storyFilter)
	if err != nil {
		log.Error("Error fetching stories: ", err)
		return nil, err
	}

	for idx,result := range stories{
		// Calls gRPC	
	commentPb,err:= s.grpcCommentClient.FindAllByStoryID(ctx,&comment_service.FindAllByStoryIDRequest{
		StoryId: result.Id,
	})
	if err != nil || commentPb == nil{
	 log.Errorf("failed when resolve comments,storyID:%d,error:%v",result.Id,err)
	}
	//Convert protobuf to main comment entity
	var comments []model.Comment
	for _, pbComment := range commentPb.Comments{
		comments = append(comments, model.Comment{
			ID: pbComment.Id,
			Comment: pbComment.Comment,
		})
	}
	stories[idx].Comments = comments
	}
	fmt.Println(stories)
	return stories,nil
	// TODO: Resolve field comments by calling grpc from comment-service
	/*
	1. go get service comment "go get github.com/kodinggo/gb-2-api-comment-service"
 	2. setup koneksi ke server grpc comment service (lihat https://github.com/kodinggo/rest-api-service-golang-private-1/blob/main/internal/cmd/server.go#L138C6-L138C26)
  	3. grpc client dipanggil di story detail usecase
 	*/
}

func (s *StoryUsecase) FindById(ctx context.Context, id int64) (*model.Story, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	story, err := s.storyRepo.FindById(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}


	if story.DeletedAt.Valid {
		return nil, fmt.Errorf("story not found")
	}

	return story, nil
}

func (s *StoryUsecase) Create(ctx context.Context, in model.CreateStoryInput) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":           ctx,
		"title":         in.Title,
		"content":       in.Content,
		"thumbnail_url": in.ThumbnailUrl,
		"category_id":   in.CategoryId,
	})

	err := v.StructCtx(ctx, in)
	if err != nil {
		log.Error("Validation error:", err)
		return err
	}

	category, err := s.categoryUsecase.FindById(ctx, int64(in.CategoryId))
	if err != nil {
		log.Error("Error fetching category:", err)
		return err
	}

	if category == nil {
		log.Error("Category not found", err)
		return errors.New("category not found")
	}

	story := model.Story{
		Title:        in.Title,
		Content:      in.Content,
		ThumbnailUrl: in.ThumbnailUrl,
		Category: model.Category{
			Id: int64(in.CategoryId),
		},
	}

	err = s.storyRepo.Create(ctx, story)
	if err != nil {
		log.Error("Error creating story:", err)
		return err
	}

	return nil
}

func (s *StoryUsecase) Update(ctx context.Context, id int64, in model.UpdateStoryInput) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":           ctx,
		"id":            id,
		"title":         in.Title,
		"content":       in.Content,
		"thumbnail_url": in.ThumbnailUrl,
		"category_id":   in.CategoryId,
	})

	err := v.StructCtx(ctx, in)
	if err != nil {
		log.Error("Validation error:", err)
		return err
	}

	category, err := s.categoryUsecase.FindById(ctx, int64(in.CategoryId))
	if err != nil {
		log.Error("Error fetching category:", err)
		return err
	}

	if category == nil {
		log.Error("Category not found")
		return errors.New("category not found")
	}

	updatedStory := model.Story{
		Id:           id,
		Title:        in.Title,
		Content:      in.Content,
		ThumbnailUrl: in.ThumbnailUrl,
		Category: model.Category{
			Id: int64(in.CategoryId),
		},
	}

	err = s.storyRepo.Update(ctx, updatedStory)
	if err != nil {
		log.Error("Error updating story:", err)
		return err
	}

	return nil
}

func (s *StoryUsecase) Delete(ctx context.Context, id int64) error {
	err := s.storyRepo.Delete(ctx, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx": ctx,
			"id":  id,
		}).Error("Failed to delete story:", err)
	}

	return nil
}