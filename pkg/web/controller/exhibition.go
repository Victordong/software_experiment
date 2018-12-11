package controller

import (
	"context"
	"software_experiment/pkg/comm/database"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
	"strconv"
)

func QueryExhibitions(ctx context.Context, queryMap map[string][]string) ([]model.ExhibitionModel, int64, error) {
	exhibitions, num, err := manager.QueryExhibitions(ctx, queryMap)

	if err != nil {
		return nil, 0, err
	}

	for index, exhibition := range exhibitions {
		user, err := manager.GetUserByUsername(ctx, exhibition.Username, false)
		if err != nil {
			return nil, 0, err
		}
		exhibitions[index].User = user
		exhibitions[index].Content = ""
	}

	return exhibitions, num, err
}

func GetExhibitionById(ctx context.Context, exhibitionId uint) (*model.ExhibitionModel, error) {
	exhibition, err := manager.GetExhibitionById(ctx, exhibitionId, true)

	if err != nil {
		return nil, err
	}

	user, err := manager.GetUserByUsername(ctx, exhibition.Username, false)
	if err != nil {
		return nil, err
	}

	exhibition.User = user

	return exhibition, err
}

func PostExhibition(ctx context.Context, exhibitionModel *model.ExhibitionModel) (*model.ExhibitionModel, error) {
	exhibition, err := manager.PostExhibition(ctx, exhibitionModel)
	if err != nil {
		return nil, err
	}
	return exhibition, err
}

func PutExhibition(ctx context.Context, exhibitionId uint, updateMap map[string]interface{}) error {
	tx := database.SqlDB.Begin()
	newCtx := context.WithValue(ctx, "tx", tx)
	_, err := manager.PutExhibition(newCtx, exhibitionId, updateMap)
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
		collectionQueryMap["collected_id"] = []string{strconv.Itoa(int(exhibitionId))}
		commentQueryMap := make(map[string][]string)
		commentQueryMap["commented_id"] = []string{strconv.Itoa(int(exhibitionId))}
		exhibitionCollections, _, err := manager.QueryExhibitionCollections(newCtx, collectionQueryMap)
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, exhibitionCollection := range exhibitionCollections {
			_, err := manager.PutExhibitionCollection(newCtx, exhibitionCollection.ID, collectionChangeMap)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		exhibitionComments, _, err := manager.QueryExhibitionComments(newCtx, commentQueryMap)
		for _, exhibitionComment := range exhibitionComments {
			_, err := manager.PutExhibitionComment(newCtx, exhibitionComment.ID, commentChangeMap)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

func DeleteExhibition(ctx context.Context, exhibitionId uint) (int64, error) {
	num, err := manager.DeleteExhibition(ctx, exhibitionId)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
