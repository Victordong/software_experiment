package formatter

import (
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/model"
)

func SupplyCommentModelToDao(exhibitionComment *model.SupplyCommentModel) *dao.SupplyComment {
	return &dao.SupplyComment{
		Username:      exhibitionComment.Username,
		Content:       exhibitionComment.Content,
		CommentedId:   exhibitionComment.CommentedId,
		CommentedName: exhibitionComment.CommentedName,
	}
}

func SupplyCommentDaoToModel(exhibitionComment *dao.SupplyComment) *model.SupplyCommentModel {
	return &model.SupplyCommentModel{
		ID:            exhibitionComment.ID,
		CreatedAt:     exhibitionComment.CreatedAt.Format("2006-02-01"),
		Username:      exhibitionComment.Username,
		Content:       exhibitionComment.Content,
		CommentedId:   exhibitionComment.CommentedId,
		CommentedName: exhibitionComment.CommentedName,
	}
}
