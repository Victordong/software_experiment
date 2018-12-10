package formatter

import (
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/model"
)

func SupplyCollectionModelToDao(exhibitionCollection *model.SupplyCollectionModel) *dao.SupplyCollection {
	return &dao.SupplyCollection{
		Username:      exhibitionCollection.Username,
		CollectedId:   exhibitionCollection.CollectedId,
		CollectedName: exhibitionCollection.CollectedName,
	}
}

func SupplyCollectionDaoToModel(exhibitionCollection *dao.SupplyCollection) *model.SupplyCollectionModel {
	return &model.SupplyCollectionModel{
		ID:            exhibitionCollection.ID,
		CreatedAt:     exhibitionCollection.CreatedAt.Format("2006-02-01"),
		Username:      exhibitionCollection.Username,
		CollectedName: exhibitionCollection.CollectedName,
		CollectedId:   exhibitionCollection.CollectedId,
	}
}
