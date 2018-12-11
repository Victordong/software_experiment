package controller

import (
	"context"
	"software_experiment/pkg/comm/database"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
	"strconv"
)

func QueryInformations(ctx context.Context, queryMap map[string][]string) ([]model.InformationModel, int64, error) {
	informations, num, err := manager.QueryInformations(ctx, queryMap)

	if err != nil {
		return nil, 0, err
	}

	for index, information := range informations {
		user, err := manager.GetUserByUsername(ctx, information.Username, false)
		if err != nil {
			return nil, 0, err
		}
		informations[index].User = user
		informations[index].Content = ""
	}

	return informations, num, err
}

func GetInformationById(ctx context.Context, informationId uint) (*model.InformationModel, error) {
	information, err := manager.GetInformationById(ctx, informationId, true)

	if err != nil {
		return nil, err
	}

	user, err := manager.GetUserByUsername(ctx, information.Username, false)
	if err != nil {
		return nil, err
	}

	information.User = user

	return information, err
}

func PostInformation(ctx context.Context, informationModel *model.InformationModel) (*model.InformationModel, error) {
	information, err := manager.PostInformation(ctx, informationModel)
	if err != nil {
		return nil, err
	}
	return information, err
}

func PutInformation(ctx context.Context, informationId uint, updateMap map[string]interface{}) error {
	tx := database.SqlDB.Begin()
	newCtx := context.WithValue(ctx, "tx", tx)
	_, err := manager.PutInformation(newCtx, informationId, updateMap)
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
		collectionQueryMap["collected_id"] = []string{strconv.Itoa(int(informationId))}
		commentQueryMap := make(map[string][]string)
		commentQueryMap["commented_id"] = []string{strconv.Itoa(int(informationId))}
		informationCollections, _, err := manager.QueryInformationCollections(newCtx, collectionQueryMap)
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, informationCollection := range informationCollections {
			_, err := manager.PutInformationCollection(newCtx, informationCollection.ID, collectionChangeMap)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		informationComments, _, err := manager.QueryInformationComments(newCtx, commentQueryMap)
		for _, informationComment := range informationComments {
			_, err := manager.PutInformationComment(newCtx, informationComment.ID, commentChangeMap)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

func DeleteInformation(ctx context.Context, informationId uint) (int64, error) {
	tx := database.SqlDB.Begin()
	newCtx := context.WithValue(ctx, "tx", tx)
	num, err := manager.DeleteInformation(newCtx, informationId)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}
	collectionQueryMap := make(map[string][]string)
	collectionQueryMap["collected_id"] = []string{strconv.Itoa(int(informationId))}
	commentQueryMap := make(map[string][]string)
	commentQueryMap["commented_id"] = []string{strconv.Itoa(int(informationId))}
	informationCollections, _, err := manager.QueryInformationCollections(newCtx, collectionQueryMap)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	for _, informationCollection := range informationCollections {
		_, err := manager.DeleteInformationCollection(newCtx, informationCollection.ID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	informationComments, _, err := manager.QueryInformationComments(newCtx, commentQueryMap)
	for _, informationComment := range informationComments {
		_, err := manager.DeleteInformationComment(newCtx, informationComment.ID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return num, nil
}
