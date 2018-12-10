package manager

import (
	"context"
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/formatter"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func QueryExhibitions(ctx context.Context, queryMap map[string][]string) ([]model.ExhibitionModel, int64, error) {
	plugin.CtxQueryMap(ctx, queryMap)
	exhibitions, num, err := dao.QueryExhibition(ctx, queryMap)
	exhibitionModels := make([]model.ExhibitionModel, 0)
	if err != nil {
		return make([]model.ExhibitionModel, 0), 0, nil
	}
	for _, exhibition := range exhibitions {
		exhibitionModel := formatter.ExhibitionDaoToModel(&exhibition)
		exhibitionModel.Content = ""
		exhibitionModels = append(exhibitionModels, *exhibitionModel)
	}
	return exhibitionModels, num, nil

}

func GetExhibitionById(ctx context.Context, exhibitionId uint, unscoped bool) (*model.ExhibitionModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(exhibitionId))}
	plugin.CtxQueryMap(ctx, queryMap)
	exhibitionDao, err := dao.GetExhibitionById(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	exhibitionModel := formatter.ExhibitionDaoToModel(exhibitionDao)
	return exhibitionModel, nil
}

func PostExhibition(ctx context.Context, exhibition *model.ExhibitionModel) (*model.ExhibitionModel, error) {
	exhibitionDao := formatter.ExhibitionModelToDao(exhibition)
	exhibitionDao, err := dao.InsertExhibition(ctx, exhibitionDao)
	if err != nil {
		return nil, err
	}
	return formatter.ExhibitionDaoToModel(exhibitionDao), nil
}

func DeleteExhibition(ctx context.Context, exhibitionId uint) (int64, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(exhibitionId))}
	plugin.CtxQueryMap(ctx, queryMap)
	shop, err := dao.GetExhibitionById(ctx, queryMap, false)
	if err != nil {
		return 0, err
	}
	num, err := dao.DeleteExhibition(ctx, shop)
	if err != nil {
		return 0, err
	}
	return num, err
}

func PutExhibition(ctx context.Context, exhibitionId uint, updateMap map[string]interface{}) (*model.ExhibitionModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(exhibitionId))}
	exhibitionDao, err := dao.GetExhibitionById(ctx, queryMap, false)
	if err != nil {
		return nil, err
	}
	args := make(map[string]interface{})
	if v, ok := updateMap["content"]; ok {
		args["content"] = v
	}
	if v, ok := updateMap["name"]; ok {
		args["name"] = v
	}
	if v, ok := updateMap["theme"]; ok {
		args["theme"] = v
	}
	exhibitionDao, err = dao.UpdateExhibition(ctx, exhibitionDao, args)
	if err != nil {
		return nil, err
	}
	return formatter.ExhibitionDaoToModel(exhibitionDao), nil
}
