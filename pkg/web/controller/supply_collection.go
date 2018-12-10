package controller

import (
	"context"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
)

func QuerySupplyCollections(ctx context.Context, queryMap map[string][]string) ([]model.SupplyCollectionModel, int64, error) {
	supplyCollections, num, err := manager.QuerySupplyCollections(ctx, queryMap)

	if err != nil {
		return nil, 0, err
	}

	for index, supplyCollection := range supplyCollections {
		user, err := manager.GetUserByUsername(ctx, supplyCollection.Username, false)
		if err != nil {
			return nil, 0, err
		}
		supplyCollections[index].User = user
		supplyModel, err := manager.GetSupplyById(ctx, supplyCollection.CollectedId, false)
		if err != nil {
			return nil, 0, err
		}
		supplyModel.Content = ""
		supplyCollections[index].Supply = supplyModel
	}

	return supplyCollections, num, err
}

func GetSupplyCollectionById(ctx context.Context, supplyCollectionId uint) (*model.SupplyCollectionModel, error) {
	supplyCollection, err := manager.GetSupplyCollectionById(ctx, supplyCollectionId, true)

	if err != nil {
		return nil, err
	}

	user, err := manager.GetUserByUsername(ctx, supplyCollection.Username, false)
	if err != nil {
		return nil, err
	}
	supplyCollection.User = user

	supplyModel, err := manager.GetSupplyById(ctx, supplyCollection.CollectedId, false)
	if err != nil {
		return nil, err
	}
	supplyModel.Content = ""
	supplyCollection.Supply = supplyModel

	return supplyCollection, err
}

func PostSupplyCollection(ctx context.Context, supplyCollectionModel *model.SupplyCollectionModel) (*model.SupplyCollectionModel, error) {
	supplyCollection, err := manager.PostSupplyCollection(ctx, supplyCollectionModel)
	if err != nil {
		return nil, err
	}
	return supplyCollection, err
}

func PutSupplyCollection(ctx context.Context, supplyCollectionId uint, updateMap map[string]interface{}) error {
	_, err := manager.PutSupplyCollection(ctx, supplyCollectionId, updateMap)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSupplyCollection(ctx context.Context, supplyCollectionId uint) (int64, error) {
	num, err := manager.DeleteSupplyCollection(ctx, supplyCollectionId)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
