package manager

import (
	"context"
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/formatter"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func QueryExhibitionComments(ctx context.Context, queryMap map[string][]string) (exhibitionCommentModelsRes []model.ExhibitionCommentModel, num int64, err error) {
	plugin.CtxQueryMap(ctx, queryMap)
	exhibitionComments, num, err := dao.QueryExhibitionComment(ctx, queryMap)
	exhibitionCommentModels := make([]model.ExhibitionCommentModel, 0)
	if err != nil {
		return make([]model.ExhibitionCommentModel, 0), 0, nil
	}
	for _, exhibitionComment := range exhibitionComments {
		exhibitionCommentModel := formatter.ExhibitionCommentDaoToModel(&exhibitionComment)
		exhibitionCommentModels = append(exhibitionCommentModels, *exhibitionCommentModel)
	}
	return exhibitionCommentModels, num, nil

}

func GetExhibitionCommentById(ctx context.Context, exhibitionCommentId uint, unscoped bool) (*model.ExhibitionCommentModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(exhibitionCommentId))}
	plugin.CtxQueryMap(ctx, queryMap)
	exhibitionCommentDao, err := dao.GetExhibitionCommentById(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	exhibitionCommentModel := formatter.ExhibitionCommentDaoToModel(exhibitionCommentDao)
	return exhibitionCommentModel, nil
}

func PostExhibitionComment(ctx context.Context, exhibitionComment *model.ExhibitionCommentModel) (*model.ExhibitionCommentModel, error) {
	exhibitionCommentDao := formatter.ExhibitionCommentModelToDao(exhibitionComment)
	exhibitionCommentDao, err := dao.InsertExhibitionComment(ctx, exhibitionCommentDao)
	if err != nil {
		return nil, err
	}
	return formatter.ExhibitionCommentDaoToModel(exhibitionCommentDao), nil
}

func DeleteExhibitionComment(ctx context.Context, exhibitionCommentId uint) (int64, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(exhibitionCommentId))}
	plugin.CtxQueryMap(ctx, queryMap)
	shop, err := dao.GetExhibitionCommentById(ctx, queryMap, false)
	if err != nil {
		return 0, err
	}
	num, err := dao.DeleteExhibitionComment(ctx, shop)
	if err != nil {
		return 0, err
	}
	return num, err
}

func PutExhibitionComment(ctx context.Context, shopId uint, updateMap map[string]interface{}) (*model.ExhibitionCommentModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(shopId))}
	exhibitionCommentDao, err := dao.GetExhibitionCommentById(ctx, queryMap, false)
	if err != nil {
		return nil, err
	}
	args := make(map[string]interface{})
	exhibitionCommentDao, err = dao.UpdateExhibitionComment(ctx, exhibitionCommentDao, args)
	if err != nil {
		return nil, err
	}
	return formatter.ExhibitionCommentDaoToModel(exhibitionCommentDao), nil
}
