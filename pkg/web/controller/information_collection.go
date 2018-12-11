package controller

import (
	"context"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
)

func QueryInformationCollections(ctx context.Context, queryMap map[string][]string) ([]model.InformationCollectionModel, int64, error) {
	informationCollections, num, err := manager.QueryInformationCollections(ctx, queryMap)

	if err != nil {
		return nil, 0, err
	}

	return informationCollections, num, err
}

func GetInformationCollectionById(ctx context.Context, informationCollectionId uint) (*model.InformationCollectionModel, error) {
	informationCollection, err := manager.GetInformationCollectionById(ctx, informationCollectionId, true)

	if err != nil {
		return nil, err
	}

	user, err := manager.GetUserByUsername(ctx, informationCollection.Username, false)
	if err != nil {
		return nil, err
	}
	informationCollection.User = user

	informationModel, err := manager.GetInformationById(ctx, informationCollection.CollectedId, false)
	if err != nil {
		return nil, err
	}
	informationModel.Content = ""
	informationCollection.Information = informationModel

	return informationCollection, err
}

func PostInformationCollection(ctx context.Context, informationCollectionModel *model.InformationCollectionModel) (*model.InformationCollectionModel, error) {
	informationCollection, err := manager.PostInformationCollection(ctx, informationCollectionModel)
	if err != nil {
		return nil, err
	}
	return informationCollection, err
}

func PutInformationCollection(ctx context.Context, informationCollectionId uint, updateMap map[string]interface{}) error {
	_, err := manager.PutInformationCollection(ctx, informationCollectionId, updateMap)
	if err != nil {
		return err
	}
	return nil
}

func DeleteInformationCollection(ctx context.Context, informationCollectionId uint) (int64, error) {
	num, err := manager.DeleteInformationCollection(ctx, informationCollectionId)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
