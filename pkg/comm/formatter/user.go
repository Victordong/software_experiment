package formatter

import (
	"software_experiment/pkg/comm/dao"
	"software_experiment/pkg/comm/model"
)

func UserModelToDao(userModel *model.UserModel) *dao.User {
	return &dao.User{
		Username:     userModel.Username,
		Name:         userModel.Name,
		Telephone:    userModel.Telephone,
		Email:        userModel.Email,
		PasswordHash: userModel.PasswordHash,
		IconUrl:      userModel.IconUrl,
		Role:         userModel.Role,
		Address:      userModel.Address,
		QqNumber:     userModel.QqNumber,
		Information:  userModel.Information,
	}
}

func UserDaoToModel(user *dao.User) *model.UserModel {
	return &model.UserModel{
		ID:          user.ID,
		Username:    user.Username,
		Name:        user.Name,
		Telephone:   user.Telephone,
		Email:       user.Email,
		IconUrl:     user.IconUrl,
		Role:        user.Role,
		Address:     user.Address,
		QqNumber:    user.QqNumber,
		Information: user.Information,
	}
}
