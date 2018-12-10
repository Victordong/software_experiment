package manager

import (
	"context"
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/database"
	"software_experiment/pkg/comm/formatter"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func QueryUsers(ctx context.Context, queryMap map[string][]string) (userModelsRes []model.UserModel, num int64, err error) {
	plugin.CtxQueryMap(ctx, queryMap)
	users, num, err := dao.QueryUser(ctx, queryMap)
	userModels := make([]model.UserModel, 0)
	if err != nil {
		return make([]model.UserModel, 0), 0, nil
	}
	for _, user := range users {
		userModel := formatter.UserDaoToModel(&user)
		userModels = append(userModels, *userModel)
	}
	return userModels, num, nil

}

func GetUserById(ctx context.Context, userId uint, unscoped bool) (*model.UserModel, error) {
	queryMap := make(map[string][]string)
	queryMap["id"] = []string{strconv.Itoa(int(userId))}
	plugin.CtxQueryMap(ctx, queryMap)
	userDao, err := dao.GetUserById(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	userModel := formatter.UserDaoToModel(userDao)
	return userModel, nil
}

func GetUserByUsername(ctx context.Context, userUsername string, unscoped bool) (*model.UserModel, error) {
	queryMap := make(map[string][]string)
	queryMap["username"] = []string{userUsername}
	plugin.CtxQueryMap(ctx, queryMap)
	userDao, err := dao.GetUserByUsername(ctx, queryMap, unscoped)
	if err != nil {
		return nil, err
	}
	userModel := formatter.UserDaoToModel(userDao)
	return userModel, nil
}

func PostUser(ctx context.Context, user *model.UserModel) (*model.UserModel, error) {
	tx := database.SqlDB.Begin()
	newCtx := context.WithValue(ctx, "tx", tx)
	userMap := map[string][]string{
		"username": []string{user.Username},
	}
	userOld, err := dao.GetUserByUsername(newCtx, userMap, false)
	if userOld != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: "username has been used",
		}
	}
	user.PasswordHash = GenPasswordHash(newCtx, user.Password)
	userDao := formatter.UserModelToDao(user)
	userDao, err = dao.InsertUser(newCtx, userDao)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return formatter.UserDaoToModel(userDao), nil
}

func DeleteUser(ctx context.Context, username string) (int64, error) {
	queryMap := make(map[string][]string)
	queryMap["username"] = []string{username}
	plugin.CtxQueryMap(ctx, queryMap)
	tx := database.SqlDB.Begin()
	newCtx := context.WithValue(ctx, "tx", tx)
	userDao, err := dao.GetUserByUsername(newCtx, queryMap, true)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	num, err := dao.DeleteUser(newCtx, userDao)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return num, nil
}

func PutUser(ctx context.Context, username string, updateMap map[string]interface{}) (*model.UserModel, error) {
	queryMap := make(map[string][]string)
	queryMap["username"] = []string{username}
	plugin.CtxQueryMap(ctx, queryMap)
	tx := database.SqlDB.Begin()
	newCtx := context.WithValue(ctx, "tx", tx)
	userDao, err := dao.GetUserByUsername(newCtx, queryMap, true)
	if err != nil {
		return nil, err
	}
	args := make(map[string]interface{})
	if v, ok := updateMap["name"]; ok {
		args["name"] = v
	}
	if v, ok := updateMap["telephone"]; ok {
		args["telephone"] = v
	}
	if v, ok := updateMap["email"]; ok {
		args["email"] = v
	}
	if v, ok := updateMap["icon_url"]; ok {
		args["icon_url"] = v
	}
	if v, ok := updateMap["role"]; ok {
		args["role"] = v
	}
	if v, ok := updateMap["address"]; ok {
		args["address"] = v
	}
	if v, ok := updateMap["qq_number"]; ok {
		args["qq_number"] = v
	}
	if v, ok := updateMap["information"]; ok {
		args["information"] = v
	}
	userDao, err = dao.UpdateUser(newCtx, userDao, args)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return formatter.UserDaoToModel(userDao), nil
}

func ChangePassword(ctx context.Context, username string, newPassword string) (bool, error) {
	queryMap := make(map[string][]string)
	queryMap["username"] = []string{username}
	user, err := dao.GetUserByUsername(ctx, queryMap, false)
	if err != nil {
		return false, err
	}
	updateMap := make(map[string]interface{})
	passwordHash := GenPasswordHash(ctx, newPassword)
	updateMap["password_hash"] = passwordHash
	_, err = dao.UpdateUser(ctx, user, updateMap)
	if err != nil {
		return false, err
	}
	return true, nil
}
