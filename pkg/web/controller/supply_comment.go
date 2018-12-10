package controller

import (
	"context"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
)

func QuerySupplyComments(ctx context.Context, queryMap map[string][]string) ([]model.SupplyCommentModel, int64, error) {
	supplyComments, num, err := manager.QuerySupplyComments(ctx, queryMap)

	if err != nil {
		return nil, 0, err
	}

	for index, supplyComment := range supplyComments {
		user, err := manager.GetUserByUsername(ctx, supplyComment.Username, false)
		if err != nil {
			return nil, 0, err
		}
		supplyComments[index].User = user
		supplyModel, err := manager.GetSupplyById(ctx, supplyComment.CommentedId, false)
		if err != nil {
			return nil, 0, err
		}
		supplyModel.Content = ""
		supplyComments[index].Supply = supplyModel
	}

	return supplyComments, num, err
}

func GetSupplyCommentById(ctx context.Context, supplyCommentId uint) (*model.SupplyCommentModel, error) {
	supplyComment, err := manager.GetSupplyCommentById(ctx, supplyCommentId, true)

	if err != nil {
		return nil, err
	}

	user, err := manager.GetUserByUsername(ctx, supplyComment.Username, false)
	if err != nil {
		return nil, err
	}
	supplyComment.User = user

	supplyModel, err := manager.GetSupplyById(ctx, supplyComment.CommentedId, false)
	if err != nil {
		return nil, err
	}
	supplyModel.Content = ""
	supplyComment.Supply = supplyModel

	return supplyComment, err
}

func PostSupplyComment(ctx context.Context, supplyCommentModel *model.SupplyCommentModel) (*model.SupplyCommentModel, error) {
	supplyComment, err := manager.PostSupplyComment(ctx, supplyCommentModel)
	if err != nil {
		return nil, err
	}
	return supplyComment, err
}

func PutSupplyComment(ctx context.Context, supplyCommentId uint, updateMap map[string]interface{}) error {
	_, err := manager.PutSupplyComment(ctx, supplyCommentId, updateMap)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSupplyComment(ctx context.Context, supplyCommentId uint) (int64, error) {
	num, err := manager.DeleteSupplyComment(ctx, supplyCommentId)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
