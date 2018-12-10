package formatter

import (
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/model"
)

func InformationCommentModelToDao(exhibitionComment *model.InformationCommentModel) *dao.InformationComment {
	return &dao.InformationComment{
		Username:      exhibitionComment.Username,
		Content:       exhibitionComment.Content,
		CommentedId:   exhibitionComment.CommentedId,
		CommentedName: exhibitionComment.CommentedName,
	}
}

func InformationCommentDaoToModel(exhibitionComment *dao.InformationComment) *model.InformationCommentModel {
	return &model.InformationCommentModel{
		ID:            exhibitionComment.ID,
		CreatedAt:     exhibitionComment.CreatedAt.Format("2006-02-01"),
		Username:      exhibitionComment.Username,
		Content:       exhibitionComment.Content,
		CommentedId:   exhibitionComment.CommentedId,
		CommentedName: exhibitionComment.CommentedName,
	}
}
