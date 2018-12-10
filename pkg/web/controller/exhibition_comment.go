package controller

import (
	"context"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
)

func QueryExhibitionComments(ctx context.Context, queryMap map[string][]string) ([]model.ExhibitionCommentModel, int64, error) {
	exhibitionComments, num, err := manager.QueryExhibitionComments(ctx, queryMap)

	if err != nil {
		return nil, 0, err
	}

	for index, exhibitionComment := range exhibitionComments {
		user, err := manager.GetUserByUsername(ctx, exhibitionComment.Username, false)
		if err != nil {
			return nil, 0, err
		}
		exhibitionComments[index].User = user
		exhibitionModel, err := manager.GetExhibitionById(ctx, exhibitionComment.CommentedId, false)
		if err != nil {
			return nil, 0, err
		}
		exhibitionModel.Content = ""
		exhibitionComments[index].Exhibition = exhibitionModel
	}

	return exhibitionComments, num, err
}

func GetExhibitionCommentById(ctx context.Context, exhibitionCommentId uint) (*model.ExhibitionCommentModel, error) {
	exhibitionComment, err := manager.GetExhibitionCommentById(ctx, exhibitionCommentId, true)

	if err != nil {
		return nil, err
	}

	user, err := manager.GetUserByUsername(ctx, exhibitionComment.Username, false)
	if err != nil {
		return nil, err
	}
	exhibitionComment.User = user

	exhibitionModel, err := manager.GetExhibitionById(ctx, exhibitionComment.CommentedId, false)
	if err != nil {
		return nil, err
	}
	exhibitionModel.Content = ""
	exhibitionComment.Exhibition = exhibitionModel

	return exhibitionComment, err
}

func PostExhibitionComment(ctx context.Context, exhibitionCommentModel *model.ExhibitionCommentModel) (*model.ExhibitionCommentModel, error) {
	exhibitionComment, err := manager.PostExhibitionComment(ctx, exhibitionCommentModel)
	if err != nil {
		return nil, err
	}
	return exhibitionComment, err
}

func PutExhibitionComment(ctx context.Context, exhibitionCommentId uint, updateMap map[string]interface{}) error {
	_, err := manager.PutExhibitionComment(ctx, exhibitionCommentId, updateMap)
	if err != nil {
		return err
	}
	return nil
}

func DeleteExhibitionComment(ctx context.Context, exhibitionCommentId uint) (int64, error) {
	num, err := manager.DeleteExhibitionComment(ctx, exhibitionCommentId)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
