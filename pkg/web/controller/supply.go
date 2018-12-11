package controller

import (
	"context"
	"software_experiment/pkg/comm/database"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
	"strconv"
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
	tx := database.SqlDB.Begin()
	newCtx := context.WithValue(ctx, "tx", tx)
	_, err := manager.PutSupply(newCtx, supplyId, updateMap)
	if err != nil {
		tx.Rollback()
		return err
	}
	if v, ok := updateMap["name"]; ok {
		collectionChangeMap := make(map[string]interface{})
		collectionChangeMap["collected_name"] = v
		commentChangeMap := make(map[string]interface{})
		commentChangeMap["commented_name"] = v
		collectionQueryMap := make(map[string][]string)
		collectionQueryMap["collected_id"] = []string{strconv.Itoa(int(supplyId))}
		commentQueryMap := make(map[string][]string)
		commentQueryMap["commented_id"] = []string{strconv.Itoa(int(supplyId))}
		supplyCollections, _, err := manager.QuerySupplyCollections(newCtx, collectionQueryMap)
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, supplyCollection := range supplyCollections {
			_, err := manager.PutSupplyCollection(newCtx, supplyCollection.ID, collectionChangeMap)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		supplyComments, _, err := manager.QuerySupplyComments(newCtx, commentQueryMap)
		for _, supplyComment := range supplyComments {
			_, err := manager.PutSupplyComment(newCtx, supplyComment.ID, commentChangeMap)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

func DeleteSupply(ctx context.Context, supplyId uint) (int64, error) {
	num, err := manager.DeleteSupply(ctx, supplyId)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
