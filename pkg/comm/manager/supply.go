package manager

import (
	"context"
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/formatter"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func QuerySupplies(ctx context.Context, queryMap map[string][]string) (supplyModelsRes []model.SupplyListNodeModel, num int64, err error) {
	plugin.CtxQueryMap(ctx, queryMap)
	supplies, num, err := dao.QuerySupply(ctx, queryMap)
	supplyModels := make([]model.SupplyListNodeModel, 0)
	if err != nil {
		return make([]model.SupplyListNodeModel, 0), 0, nil
	}
	for _, supply := range supplies {
		supplyModel := formatter.SupplyListNodeDaoToModel(&supply)
		supplyModels = append(supplyModels, *supplyModel)
	}
	return supplyModels, num, nil

}

func GetSupplyById(ctx context.Context, supplyId uint, unscoped bool) (*model.SupplyModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(supplyId))}
	plugin.CtxQueryMap(ctx, queryMap)
	supplyDao, err := dao.GetSupplyById(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	supplyModel := formatter.SupplyDaoToModel(supplyDao)
	return supplyModel, nil
}

func PostSupply(ctx context.Context, supply *model.SupplyModel) (*model.SupplyModel, error) {
	supplyDao := formatter.SupplyModelToDao(supply)
	supplyDao, err := dao.InsertSupply(ctx, supplyDao)
	if err != nil {
		return nil, err
	}
	return formatter.SupplyDaoToModel(supplyDao), nil
}

func DeleteSupply(ctx context.Context, supplyId uint) (int64, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(supplyId))}
	plugin.CtxQueryMap(ctx, queryMap)
	shop, err := dao.GetSupplyById(ctx, queryMap, false)
	if err != nil {
		return 0, err
	}
	num, err := dao.DeleteSupply(ctx, shop)
	if err != nil {
		return 0, err
	}
	return num, err
}

func PutSupply(ctx context.Context, shopId uint, updateMap map[string]interface{}) (*model.SupplyModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(shopId))}
	supplyDao, err := dao.GetSupplyById(ctx, queryMap, false)
	if err != nil {
		return nil, err
	}
	args := make(map[string]interface{})
	if v, ok := updateMap["content"]; ok {
		args["content"] = v
	}
	if v, ok := updateMap["type"]; ok {
		args["type"] = v
	}
	if v, ok := updateMap["expiry_date"]; ok {
		args["expiry_date"] = v
	}
	if v, ok := updateMap["name"]; ok {
		args["name"] = v
	}
	supplyDao, err = dao.UpdateSupply(ctx, supplyDao, args)
	if err != nil {
		return nil, err
	}
	return formatter.SupplyDaoToModel(supplyDao), nil
}
