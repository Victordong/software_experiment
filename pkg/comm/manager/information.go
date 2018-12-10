package manager

import (
	"context"
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/formatter"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func QueryInformations(ctx context.Context, queryMap map[string][]string) (informationModelsRes []model.InformationModel, num int64, err error) {
	plugin.CtxQueryMap(ctx, queryMap)
	informations, num, err := dao.QueryInformation(ctx, queryMap)
	informationModels := make([]model.InformationModel, 0)
	if err != nil {
		return make([]model.InformationModel, 0), 0, nil
	}
	for _, information := range informations {
		informationModel := formatter.InformationDaoToModel(&information)
		informationModel.Content = ""
		informationModels = append(informationModels, *informationModel)
	}
	return informationModels, num, nil

}

func GetInformationById(ctx context.Context, informationId uint, unscoped bool) (*model.InformationModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(informationId))}
	plugin.CtxQueryMap(ctx, queryMap)
	informationDao, err := dao.GetInformationById(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	informationModel := formatter.InformationDaoToModel(informationDao)
	return informationModel, nil
}

func PostInformation(ctx context.Context, information *model.InformationModel) (*model.InformationModel, error) {
	informationDao := formatter.InformationModelToDao(information)
	informationDao, err := dao.InsertInformation(ctx, informationDao)
	if err != nil {
		return nil, err
	}
	return formatter.InformationDaoToModel(informationDao), nil
}

func DeleteInformation(ctx context.Context, informationId uint) (int64, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(informationId))}
	plugin.CtxQueryMap(ctx, queryMap)
	shop, err := dao.GetInformationById(ctx, queryMap, false)
	if err != nil {
		return 0, err
	}
	num, err := dao.DeleteInformation(ctx, shop)
	if err != nil {
		return 0, err
	}
	return num, err
}

func PutInformation(ctx context.Context, informationId uint, updateMap map[string]interface{}) (*model.InformationModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(informationId))}
	informationDao, err := dao.GetInformationById(ctx, queryMap, false)
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
	informationDao, err = dao.UpdateInformation(ctx, informationDao, args)
	if err != nil {
		return nil, err
	}
	return formatter.InformationDaoToModel(informationDao), nil
}
