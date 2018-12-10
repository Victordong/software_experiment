package manager

import (
	"context"
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/formatter"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func QuerySupplyComments(ctx context.Context, queryMap map[string][]string) (supplyCommentModelsRes []model.SupplyCommentModel, num int64, err error) {
	plugin.CtxQueryMap(ctx, queryMap)
	supplyComments, num, err := dao.QuerySupplyComment(ctx, queryMap)
	supplyCommentModels := make([]model.SupplyCommentModel, 0)
	if err != nil {
		return make([]model.SupplyCommentModel, 0), 0, nil
	}
	for _, supplyComment := range supplyComments {
		supplyCommentModel := formatter.SupplyCommentDaoToModel(&supplyComment)
		supplyCommentModels = append(supplyCommentModels, *supplyCommentModel)
	}
	return supplyCommentModels, num, nil

}

func GetSupplyCommentById(ctx context.Context, supplyCommentId uint, unscoped bool) (*model.SupplyCommentModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(supplyCommentId))}
	plugin.CtxQueryMap(ctx, queryMap)
	supplyCommentDao, err := dao.GetSupplyCommentById(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	supplyCommentModel := formatter.SupplyCommentDaoToModel(supplyCommentDao)
	return supplyCommentModel, nil
}

func PostSupplyComment(ctx context.Context, supplyComment *model.SupplyCommentModel) (*model.SupplyCommentModel, error) {
	supplyCommentDao := formatter.SupplyCommentModelToDao(supplyComment)
	supplyCommentDao, err := dao.InsertSupplyComment(ctx, supplyCommentDao)
	if err != nil {
		return nil, err
	}
	return formatter.SupplyCommentDaoToModel(supplyCommentDao), nil
}

func DeleteSupplyComment(ctx context.Context, supplyCommentId uint) (int64, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(supplyCommentId))}
	plugin.CtxQueryMap(ctx, queryMap)
	shop, err := dao.GetSupplyCommentById(ctx, queryMap, false)
	if err != nil {
		return 0, err
	}
	num, err := dao.DeleteSupplyComment(ctx, shop)
	if err != nil {
		return 0, err
	}
	return num, err
}

func PutSupplyComment(ctx context.Context, shopId uint, updateMap map[string]interface{}) (*model.SupplyCommentModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(shopId))}
	supplyCommentDao, err := dao.GetSupplyCommentById(ctx, queryMap, false)
	if err != nil {
		return nil, err
	}
	args := make(map[string]interface{})
	supplyCommentDao, err = dao.UpdateSupplyComment(ctx, supplyCommentDao, args)
	if err != nil {
		return nil, err
	}
	return formatter.SupplyCommentDaoToModel(supplyCommentDao), nil
}
