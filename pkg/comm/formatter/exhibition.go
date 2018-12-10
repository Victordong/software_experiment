package formatter

import (
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/model"
)

func ExhibitionModelToDao(exhibition *model.ExhibitionModel) *dao.Exhibition {
	return &dao.Exhibition{
		Username: exhibition.Username,
		Content:  exhibition.Content,
		Theme:    exhibition.Theme,
		Name:     exhibition.Name,
		VisitNum: exhibition.VisitNum,
	}
}

func ExhibitionDaoToModel(exhibition *dao.Exhibition) *model.ExhibitionModel {
	return &model.ExhibitionModel{
		ID:        exhibition.ID,
		CreatedAt: exhibition.CreatedAt.Format("2006-02-01"),
		Username:  exhibition.Username,
		Content:   exhibition.Content,
		Theme:     exhibition.Theme,
		Name:      exhibition.Name,
		VisitNum:  exhibition.VisitNum,
	}
}

func ExhibitionListNodeDaoToModel(exhibition *dao.Exhibition) *model.ExhibitionListNodeModel {
	return &model.ExhibitionListNodeModel{
		ID:        exhibition.ID,
		CreatedAt: exhibition.CreatedAt.Format("2006-02-01"),
		Username:  exhibition.Username,
		Theme:     exhibition.Theme,
		VisitNum:  exhibition.VisitNum,
	}
}
