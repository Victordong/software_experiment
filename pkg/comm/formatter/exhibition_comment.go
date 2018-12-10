package formatter

import (
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/model"
)

func ExhibitionCommentModelToDao(exhibitionComment *model.ExhibitionCommentModel) *dao.ExhibitionComment {
	return &dao.ExhibitionComment{
		Username:      exhibitionComment.Username,
		Content:       exhibitionComment.Content,
		CommentedId:   exhibitionComment.CommentedId,
		CommentedName: exhibitionComment.CommentedName,
	}
}

func ExhibitionCommentDaoToModel(exhibitionComment *dao.ExhibitionComment) *model.ExhibitionCommentModel {
	return &model.ExhibitionCommentModel{
		ID:            exhibitionComment.ID,
		CreatedAt:     exhibitionComment.CreatedAt.Format("2006-02-01"),
		Username:      exhibitionComment.Username,
		Content:       exhibitionComment.Content,
		CommentedId:   exhibitionComment.CommentedId,
		CommentedName: exhibitionComment.CommentedName,
	}
}
