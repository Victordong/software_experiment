package controller

import (
	"context"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
)

func QuerySupplys(ctx context.Context, queryMap map[string][]string) ([]model.SupplyModel, int64, error) {
	supplys, num, err := manager.QuerySupplies(ctx, queryMap)

	if err != nil {
		return nil, 0, err
	}

	for index, supply := range supplys {
		user, err := manager.GetUserByUsername(ctx, supply.Username, false)
		if err != nil {
			return nil, 0, err
		}
		supplys[index].User = user
		supplys[index].Content = ""
	}

	return supplys, num, err
}

func GetSupplyById(ctx context.Context, supplyId uint) (*model.SupplyModel, error) {
	supply, err := manager.GetSupplyById(ctx, supplyId, true)

	if err != nil {
		return nil, err
	}

	user, err := manager.GetUserByUsername(ctx, supply.Username, false)
	if err != nil {
		return nil, err
	}

	supply.User = user

	return supply, err
}

func PostSupply(ctx context.Context, supplyModel *model.SupplyModel) (*model.SupplyModel, error) {
	supply, err := manager.PostSupply(ctx, supplyModel)
	if err != nil {
		return nil, err
	}
	return supply, err
}

func PutSupply(ctx context.Context, supplyId uint, updateMap map[string]interface{}) error {
	_, err := manager.PutSupply(ctx, supplyId, updateMap)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSupply(ctx context.Context, supplyId uint) (int64, error) {
	num, err := manager.DeleteSupply(ctx, supplyId)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
