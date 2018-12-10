package controller

import (
	"context"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
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
	_, err := manager.PutInformation(ctx, informationId, updateMap)
	if err != nil {
		return err
	}
	return nil
}

func DeleteInformation(ctx context.Context, informationId uint) (int64, error) {
	num, err := manager.DeleteInformation(ctx, informationId)
	if err != nil {
		return 0, nil
	}
	return num, nil
}
