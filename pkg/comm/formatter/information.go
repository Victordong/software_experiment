package formatter

import (
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/model"
)

func InformationModelToDao(informationModel *model.InformationModel) *dao.Information {
	return &dao.Information{
		Username: informationModel.Username,
		Content:  informationModel.Content,
		Name:     informationModel.Name,
		VisitNum: informationModel.VisitNum,
	}
}

func InformationListNodeDaoToModel(information *dao.Information) *model.InformationListNodeModel {
	return &model.InformationListNodeModel{
		ID:        information.ID,
		CreatedAt: information.CreatedAt.Format("2006-02-01"),
		Username:  information.Username,
		Name:      information.Name,
		VisitNum:  information.VisitNum,
	}
}

func InformationDaoToModel(information *dao.Information) *model.InformationModel {
	return &model.InformationModel{
		ID:        information.ID,
		CreatedAt: information.CreatedAt.Format("2006-02-01"),
		Username:  information.Username,
		Name:      information.Name,
		VisitNum:  information.VisitNum,
		Content:   information.Content,
	}
}
