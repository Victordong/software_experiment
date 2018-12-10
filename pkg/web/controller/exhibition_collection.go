package controller

import (
	"context"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
)

func QueryExhibitionCollections(ctx context.Context, queryMap map[string][]string) ([]model.ExhibitionCollectionModel, int64, error) {
	exhibitionCollections, num, err := manager.QueryExhibitionCollections(ctx, queryMap)

	if err != nil {
		return nil, 0, err
	}

	for index, exhibitionCollection := range exhibitionCollections {
		user, err := manager.GetUserByUsername(ctx, exhibitionCollection.Username, false)
		if err != nil {
			return nil, 0, err
		}
		exhibitionCollections[index].User = user
		exhibitionModel, err := manager.GetExhibitionById(ctx, exhibitionCollection.CollectedId, false)
		if err != nil {
			return nil, 0, err
		}
		exhibitionModel.Content = ""
		exhibitionCollections[index].Exhibition = exhibitionModel
	}

	return exhibitionCollections, num, err
}

func GetExhibitionCollectionById(ctx context.Context, exhibitionCollectionId uint) (*model.ExhibitionCollectionModel, error) {
	exhibitionCollection, err := manager.GetExhibitionCollectionById(ctx, exhibitionCollectionId, true)

	if err != nil {
		return nil, err
	}

	user, err := manager.GetUserByUsername(ctx, exhibitionCollection.Username, false)
	if err != nil {
		return nil, err
	}
	exhibitionCollection.User = user

	exhibitionModel, err := manager.GetExhibitionById(ctx, exhibitionCollection.CollectedId, false)
	if err != nil {
		return nil, err
	}
	exhibitionModel.Content = ""
	exhibitionCollection.Exhibition = exhibitionModel

	return exhibitionCollection, err
}

func PostExhibitionCollection(ctx context.Context, exhibitionCollectionModel *model.ExhibitionCollectionModel) (*model.ExhibitionCollectionModel, error) {
	exhibitionCollection, err := manager.PostExhibitionCollection(ctx, exhibitionCollectionModel)
	if err != nil {
		return nil, err
	}
	return exhibitionCollection, err
}

func PutExhibitionCollection(ctx context.Context, exhibitionCollectionId uint, updateMap map[string]interface{}) error {
	_, err := manager.PutExhibitionCollection(ctx, exhibitionCollectionId, updateMap)
	if err != nil {
		return err
	}
	return nil
}

func DeleteExhibitionCollection(ctx context.Context, exhibitionCollectionId uint) (int64, error) {
	num, err := manager.DeleteExhibitionCollection(ctx, exhibitionCollectionId)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
