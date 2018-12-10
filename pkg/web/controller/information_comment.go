package controller

import (
	"context"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
)

func QueryInformationComments(ctx context.Context, queryMap map[string][]string) ([]model.InformationCommentModel, int64, error) {
	informationComments, num, err := manager.QueryInformationComments(ctx, queryMap)

	if err != nil {
		return nil, 0, err
	}

	for index, informationComment := range informationComments {
		user, err := manager.GetUserByUsername(ctx, informationComment.Username, false)
		if err != nil {
			return nil, 0, err
		}
		informationComments[index].User = user
		informationModel, err := manager.GetInformationById(ctx, informationComment.CommentedId, false)
		if err != nil {
			return nil, 0, err
		}
		informationModel.Content = ""
		informationComments[index].Information = informationModel
	}

	return informationComments, num, err
}

func GetInformationCommentById(ctx context.Context, informationCommentId uint) (*model.InformationCommentModel, error) {
	informationComment, err := manager.GetInformationCommentById(ctx, informationCommentId, true)

	if err != nil {
		return nil, err
	}

	user, err := manager.GetUserByUsername(ctx, informationComment.Username, false)
	if err != nil {
		return nil, err
	}
	informationComment.User = user

	informationModel, err := manager.GetInformationById(ctx, informationComment.CommentedId, false)
	if err != nil {
		return nil, err
	}
	informationModel.Content = ""
	informationComment.Information = informationModel

	return informationComment, err
}

func PostInformationComment(ctx context.Context, informationCommentModel *model.InformationCommentModel) (*model.InformationCommentModel, error) {
	informationComment, err := manager.PostInformationComment(ctx, informationCommentModel)
	if err != nil {
		return nil, err
	}
	return informationComment, err
}

func PutInformationComment(ctx context.Context, informationCommentId uint, updateMap map[string]interface{}) error {
	_, err := manager.PutInformationComment(ctx, informationCommentId, updateMap)
	if err != nil {
		return err
	}
	return nil
}

func DeleteInformationComment(ctx context.Context, informationCommentId uint) (int64, error) {
	num, err := manager.DeleteInformationComment(ctx, informationCommentId)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
