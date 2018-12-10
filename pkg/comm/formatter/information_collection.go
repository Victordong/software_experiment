package formatter

import (
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/model"
)

func InformationCollectionModelToDao(exhibitionCollection *model.InformationCollectionModel) *dao.InformationCollection {
	return &dao.InformationCollection{
		Username:      exhibitionCollection.Username,
		CollectedId:   exhibitionCollection.CollectedId,
		CollectedName: exhibitionCollection.CollectedName,
	}
}

func InformationCollectionDaoToModel(exhibitionCollection *dao.InformationCollection) *model.InformationCollectionModel {
	return &model.InformationCollectionModel{
		ID:            exhibitionCollection.ID,
		CreatedAt:     exhibitionCollection.CreatedAt.Format("2006-02-01"),
		Username:      exhibitionCollection.Username,
		CollectedName: exhibitionCollection.CollectedName,
		CollectedId:   exhibitionCollection.CollectedId,
	}
}
