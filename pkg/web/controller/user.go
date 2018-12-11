package controller

import (
	"context"
	"software_experiment/pkg/comm/database"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
)

func QueryUser(ctx context.Context, queryMap map[string][]string) ([]model.UserModel, int64, error) {
	users, num, err := manager.QueryUsers(ctx, queryMap)
	if err != nil {
		return nil, 0, err
	}
	return users, num, err
}

func GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error) {
	user, err := manager.GetUserByUsername(ctx, username, true)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func PostUser(ctx context.Context, userModel *model.UserModel) (*model.UserModel, error) {
	queryMap := make(map[string][]string)
	queryMap["username"] = []string{userModel.Username}
	newCtx := context.Background()
	_, num, err := manager.QueryUsers(newCtx, queryMap)
	if num != 0 {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: "用户名已经被使用",
		}
	}
	userModel, err = manager.PostUser(ctx, userModel)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func DeleteUser(ctx context.Context, username string) error {
	tx := database.SqlDB.Begin()
	newCtx := context.WithValue(ctx, "tx", tx)
	_, err := manager.DeleteUser(ctx, username)
	if err != nil {
		tx.Rollback()
		return err
	}
	exhibitionMap := make(map[string][]string)
	exhibitionMap["username"] = []string{username}
	exhibitions, _, err := manager.QueryExhibitions(newCtx, exhibitionMap)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, exhibition := range exhibitions {
		_, err = manager.DeleteExhibition(newCtx, exhibition.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	informationMap := make(map[string][]string)
	informationMap["username"] = []string{username}
	informations, _, err := manager.QueryInformations(newCtx, informationMap)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, information := range informations {
		_, err = manager.DeleteInformation(newCtx, information.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	supplyMap := make(map[string][]string)
	supplyMap["username"] = []string{username}
	supplies, _, err := manager.QuerySupplies(newCtx, supplyMap)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, supply := range supplies {
		_, err = manager.DeleteSupply(newCtx, supply.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func PutUser(ctx context.Context, username string, updateMap map[string]interface{}) (*model.UserModel, error) {
	userModel, err := manager.PutUser(ctx, username, updateMap)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}
