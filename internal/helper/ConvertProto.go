package helper

import (
	"github.com/kodinggo/gb-2-api-comment-service/pb/comment_service"
	"github.com/kodinggo/gb-2-api-story-service/internal/model"
)


func ConvertPbCommentToModelComments(pbcomment []*comment_service.Comment)[]*model.Comment{

	var comments []*model.Comment
	for _, pbComment := range pbcomment {
		comments = append(comments, &model.Comment{
			ID: pbComment.Id,
			Comment: pbComment.Comment,
			StoryID: pbComment.StoryId,
		})
	}
	return comments
}