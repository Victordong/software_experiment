package formatter

import (
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/model"
)

func ExhibitionCollectionModelToDao(exhibitionCollection *model.ExhibitionCollectionModel) *dao.ExhibitionCollection {
	return &dao.ExhibitionCollection{
		Username:      exhibitionCollection.Username,
		CollectedId:   exhibitionCollection.CollectedId,
		CollectedName: exhibitionCollection.CollectedName,
	}
}

func ExhibitionCollectionDaoToModel(exhibitionCollection *dao.ExhibitionCollection) *model.ExhibitionCollectionModel {
	return &model.ExhibitionCollectionModel{
		ID:            exhibitionCollection.ID,
		CreatedAt:     exhibitionCollection.CreatedAt.Format("2006-02-01"),
		Username:      exhibitionCollection.Username,
		CollectedName: exhibitionCollection.CollectedName,
		CollectedId:   exhibitionCollection.CollectedId,
	}
}
