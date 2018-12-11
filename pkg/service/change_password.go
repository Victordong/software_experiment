package service

import (
	"context"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
)

func ConfirmOldPassword(ctx context.Context, username string, oldPassword string) (bool, error) {
	operator, err := manager.GetUserByUsername(ctx, username, false)
	if err != nil {
		return false, err
	}
	passwordHash := manager.GenPasswordHash(ctx, oldPassword)
	if passwordHash != operator.PasswordHash {
		return false, plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: "密码错误",
		}
	}
	return true, nil
}

func SetSessionRedis(ctx context.Context, username string) (string, error) {
	var session string
	err := db.RedisClient.Set(session, username, 0).Err()
	if err != nil {
		return "", plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: err.Error(),
		}
	}
	return session, nil
}

func GetUserFromRedis(ctx context.Context, session string) (string, error) {
	username, err := db.RedisClient.Get(session).Result()
	if err != nil {
		return "", plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: err.Error(),
		}
	}
	return username, nil
}

func ChangePassword(ctx context.Context, changePasswordModel model.ChangePasswordModel) (bool, error) {
	if changePasswordModel.Username != "" {
		_, err := manager.GetUserByUsername(ctx, changePasswordModel.Username, false)
		if err != nil {
			return false, plugin.CustomErr{
				Code:        500,
				StatusCode:  200,
				Information: err.Error(),
			}
		}
		if changePasswordModel.OldPassword == "" || changePasswordModel.Password == "" {
			return false, plugin.CustomErr{
				Code:        500,
				StatusCode:  200,
				Information: "old password and new password must be given",
			}
		}
		_, err = ConfirmOldPassword(ctx, changePasswordModel.Username, changePasswordModel.OldPassword)
		if err != nil {
			return false, plugin.CustomErr{
				Code:        500,
				StatusCode:  200,
				Information: err.Error(),
			}
		}
		_, err = manager.ChangePassword(ctx, changePasswordModel.Username, changePasswordModel.OldPassword)
		if err != nil {
			return false, plugin.CustomErr{
				Code:        500,
				StatusCode:  200,
				Information: err.Error(),
			}
		}

	} else {

	}
}
