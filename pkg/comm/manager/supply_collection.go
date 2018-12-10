package manager

import (
	"context"
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/formatter"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func QuerySupplyCollections(ctx context.Context, queryMap map[string][]string) (supplyCollectionModelsRes []model.SupplyCollectionModel, num int64, err error) {
	plugin.CtxQueryMap(ctx, queryMap)
	supplyCollections, num, err := dao.QuerySupplyCollection(ctx, queryMap)
	supplyCollectionModels := make([]model.SupplyCollectionModel, 0)
	if err != nil {
		return make([]model.SupplyCollectionModel, 0), 0, nil
	}
	for _, supplyCollection := range supplyCollections {
		supplyCollectionModel := formatter.SupplyCollectionDaoToModel(&supplyCollection)
		supplyCollectionModels = append(supplyCollectionModels, *supplyCollectionModel)
	}
	return supplyCollectionModels, num, nil

}

func GetSupplyCollectionById(ctx context.Context, supplyCollectionId uint, unscoped bool) (*model.SupplyCollectionModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(supplyCollectionId))}
	plugin.CtxQueryMap(ctx, queryMap)
	supplyCollectionDao, err := dao.GetSupplyCollectionById(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	supplyCollectionModel := formatter.SupplyCollectionDaoToModel(supplyCollectionDao)
	return supplyCollectionModel, nil
}

func PostSupplyCollection(ctx context.Context, supplyCollection *model.SupplyCollectionModel) (*model.SupplyCollectionModel, error) {
	supplyCollectionDao := formatter.SupplyCollectionModelToDao(supplyCollection)
	supplyCollectionDao, err := dao.InsertSupplyCollection(ctx, supplyCollectionDao)
	if err != nil {
		return nil, err
	}
	return formatter.SupplyCollectionDaoToModel(supplyCollectionDao), nil
}

func DeleteSupplyCollection(ctx context.Context, supplyCollectionId uint) (int64, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(supplyCollectionId))}
	plugin.CtxQueryMap(ctx, queryMap)
	shop, err := dao.GetSupplyCollectionById(ctx, queryMap, false)
	if err != nil {
		return 0, err
	}
	num, err := dao.DeleteSupplyCollection(ctx, shop)
	if err != nil {
		return 0, err
	}
	return num, err
}

func PutSupplyCollection(ctx context.Context, supplyCollectionId uint, updateMap map[string]interface{}) (*model.SupplyCollectionModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(supplyCollectionId))}
	supplyCollectionDao, err := dao.GetSupplyCollectionById(ctx, queryMap, false)
	if err != nil {
		return nil, err
	}
	args := make(map[string]interface{})
	supplyCollectionDao, err = dao.UpdateSupplyCollection(ctx, supplyCollectionDao, args)
	if err != nil {
		return nil, err
	}
	return formatter.SupplyCollectionDaoToModel(supplyCollectionDao), nil
}
