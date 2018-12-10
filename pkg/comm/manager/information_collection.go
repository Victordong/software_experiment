package manager

import (
	"context"
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/formatter"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func QueryInformationCollections(ctx context.Context, queryMap map[string][]string) (informationCollectionModelsRes []model.InformationCollectionModel, num int64, err error) {
	plugin.CtxQueryMap(ctx, queryMap)
	informationCollections, num, err := dao.QueryInformationCollection(ctx, queryMap)
	informationCollectionModels := make([]model.InformationCollectionModel, 0)
	if err != nil {
		return make([]model.InformationCollectionModel, 0), 0, nil
	}
	for _, informationCollection := range informationCollections {
		informationCollectionModel := formatter.InformationCollectionDaoToModel(&informationCollection)
		informationCollectionModels = append(informationCollectionModels, *informationCollectionModel)
	}
	return informationCollectionModels, num, nil

}

func GetInformationCollectionById(ctx context.Context, informationCollectionId uint, unscoped bool) (*model.InformationCollectionModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(informationCollectionId))}
	plugin.CtxQueryMap(ctx, queryMap)
	informationCollectionDao, err := dao.GetInformationCollectionById(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	informationCollectionModel := formatter.InformationCollectionDaoToModel(informationCollectionDao)
	return informationCollectionModel, nil
}

func PostInformationCollection(ctx context.Context, informationCollection *model.InformationCollectionModel) (*model.InformationCollectionModel, error) {
	informationCollectionDao := formatter.InformationCollectionModelToDao(informationCollection)
	informationCollectionDao, err := dao.InsertInformationCollection(ctx, informationCollectionDao)
	if err != nil {
		return nil, err
	}
	return formatter.InformationCollectionDaoToModel(informationCollectionDao), nil
}

func DeleteInformationCollection(ctx context.Context, informationCollectionId uint) (int64, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(informationCollectionId))}
	plugin.CtxQueryMap(ctx, queryMap)
	shop, err := dao.GetInformationCollectionById(ctx, queryMap, false)
	if err != nil {
		return 0, err
	}
	num, err := dao.DeleteInformationCollection(ctx, shop)
	if err != nil {
		return 0, err
	}
	return num, err
}

func PutInformationCollection(ctx context.Context, shopId uint, updateMap map[string]interface{}) (*model.InformationCollectionModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(shopId))}
	informationCollectionDao, err := dao.GetInformationCollectionById(ctx, queryMap, false)
	if err != nil {
		return nil, err
	}
	args := make(map[string]interface{})
	informationCollectionDao, err = dao.UpdateInformationCollection(ctx, informationCollectionDao, args)
	if err != nil {
		return nil, err
	}
	return formatter.InformationCollectionDaoToModel(informationCollectionDao), nil
}
