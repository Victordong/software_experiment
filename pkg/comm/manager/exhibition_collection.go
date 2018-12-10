package manager

import (
	"context"
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/formatter"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func QueryExhibitionCollections(ctx context.Context, queryMap map[string][]string) (exhibitionCollectionModelsRes []model.ExhibitionCollectionModel, num int64, err error) {
	plugin.CtxQueryMap(ctx, queryMap)
	exhibitionCollections, num, err := dao.QueryExhibitionCollection(ctx, queryMap)
	exhibitionCollectionModels := make([]model.ExhibitionCollectionModel, 0)
	if err != nil {
		return make([]model.ExhibitionCollectionModel, 0), 0, nil
	}
	for _, exhibitionCollection := range exhibitionCollections {
		exhibitionCollectionModel := formatter.ExhibitionCollectionDaoToModel(&exhibitionCollection)
		exhibitionCollectionModels = append(exhibitionCollectionModels, *exhibitionCollectionModel)
	}
	return exhibitionCollectionModels, num, nil

}

func GetExhibitionCollectionById(ctx context.Context, exhibitionCollectionId uint, unscoped bool) (*model.ExhibitionCollectionModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(exhibitionCollectionId))}
	plugin.CtxQueryMap(ctx, queryMap)
	exhibitionCollectionDao, err := dao.GetExhibitionCollectionById(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	exhibitionCollectionModel := formatter.ExhibitionCollectionDaoToModel(exhibitionCollectionDao)
	return exhibitionCollectionModel, nil
}

func PostExhibitionCollection(ctx context.Context, exhibitionCollection *model.ExhibitionCollectionModel) (*model.ExhibitionCollectionModel, error) {
	exhibitionCollectionDao := formatter.ExhibitionCollectionModelToDao(exhibitionCollection)
	exhibitionCollectionDao, err := dao.InsertExhibitionCollection(ctx, exhibitionCollectionDao)
	if err != nil {
		return nil, err
	}
	return formatter.ExhibitionCollectionDaoToModel(exhibitionCollectionDao), nil
}

func DeleteExhibitionCollection(ctx context.Context, exhibitionCollectionId uint) (int64, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(exhibitionCollectionId))}
	plugin.CtxQueryMap(ctx, queryMap)
	shop, err := dao.GetExhibitionCollectionById(ctx, queryMap, false)
	if err != nil {
		return 0, err
	}
	num, err := dao.DeleteExhibitionCollection(ctx, shop)
	if err != nil {
		return 0, err
	}
	return num, err
}

func PutExhibitionCollection(ctx context.Context, shopId uint, updateMap map[string]interface{}) (*model.ExhibitionCollectionModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(shopId))}
	exhibitionCollectionDao, err := dao.GetExhibitionCollectionById(ctx, queryMap, false)
	if err != nil {
		return nil, err
	}
	args := make(map[string]interface{})
	exhibitionCollectionDao, err = dao.UpdateExhibitionCollection(ctx, exhibitionCollectionDao, args)
	if err != nil {
		return nil, err
	}
	return formatter.ExhibitionCollectionDaoToModel(exhibitionCollectionDao), nil
}
