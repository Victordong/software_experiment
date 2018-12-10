package formatter

import (
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/model"
)

var SupplyStatusMap = map[string]int{"supply": 1, "close": 2}
var ShopStatusMapRsv = map[int]string{1: "supply", 2: "close"}

func SupplyModelToDao(informationModel *model.SupplyModel) *dao.Supply {
	return &dao.Supply{
		Username:   informationModel.Username,
		Content:    informationModel.Content,
		Name:       informationModel.Name,
		VisitNum:   informationModel.VisitNum,
		Type:       SupplyStatusMap[informationModel.Type],
		ExpiryDate: informationModel.ExpiryDate,
	}
}

func SupplyListNodeDaoToModel(information *dao.Supply) *model.SupplyListNodeModel {
	return &model.SupplyListNodeModel{
		ID:        information.ID,
		CreatedAt: information.CreatedAt.Format("2006-02-01"),
		Username:  information.Username,
		Name:      information.Name,
		VisitNum:  information.VisitNum,
	}
}

func SupplyDaoToModel(information *dao.Supply) *model.SupplyModel {
	return &model.SupplyModel{
		ID:        information.ID,
		CreatedAt: information.CreatedAt.Format("2006-02-01"),
		Username:  information.Username,
		Name:      information.Name,
		VisitNum:  information.VisitNum,
		Content:   information.Content,
	}
}
