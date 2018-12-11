package manager

import (
	"context"
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/formatter"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func QueryInformationComments(ctx context.Context, queryMap map[string][]string) (informationCommentModelsRes []model.InformationCommentModel, num int64, err error) {
	plugin.CtxQueryMap(ctx, queryMap)
	informationComments, num, err := dao.QueryInformationComment(ctx, queryMap)
	informationCommentModels := make([]model.InformationCommentModel, 0)
	if err != nil {
		return make([]model.InformationCommentModel, 0), 0, nil
	}
	for _, informationComment := range informationComments {
		informationCommentModel := formatter.InformationCommentDaoToModel(&informationComment)
		informationCommentModels = append(informationCommentModels, *informationCommentModel)
	}
	return informationCommentModels, num, nil

}

func GetInformationCommentById(ctx context.Context, informationCommentId uint, unscoped bool) (*model.InformationCommentModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(informationCommentId))}
	plugin.CtxQueryMap(ctx, queryMap)
	informationCommentDao, err := dao.GetInformationCommentById(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	informationCommentModel := formatter.InformationCommentDaoToModel(informationCommentDao)
	return informationCommentModel, nil
}

func PostInformationComment(ctx context.Context, informationComment *model.InformationCommentModel) (*model.InformationCommentModel, error) {
	informationCommentDao := formatter.InformationCommentModelToDao(informationComment)
	informationCommentDao, err := dao.InsertInformationComment(ctx, informationCommentDao)
	if err != nil {
		return nil, err
	}
	return formatter.InformationCommentDaoToModel(informationCommentDao), nil
}

func DeleteInformationComment(ctx context.Context, informationCommentId uint) (int64, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(informationCommentId))}
	plugin.CtxQueryMap(ctx, queryMap)
	shop, err := dao.GetInformationCommentById(ctx, queryMap, false)
	if err != nil {
		return 0, err
	}
	num, err := dao.DeleteInformationComment(ctx, shop)
	if err != nil {
		return 0, err
	}
	return num, err
}

func PutInformationComment(ctx context.Context, informationCommentId uint, updateMap map[string]interface{}) (*model.InformationCommentModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(informationCommentId))}
	informationCommentDao, err := dao.GetInformationCommentById(ctx, queryMap, false)
	if err != nil {
		return nil, err
	}
	args := make(map[string]interface{})
	if v, ok := updateMap["content"]; ok {
		args["content"] = v
	}
	informationCommentDao, err = dao.UpdateInformationComment(ctx, informationCommentDao, args)
	if err != nil {
		return nil, err
	}
	return formatter.InformationCommentDaoToModel(informationCommentDao), nil
}
