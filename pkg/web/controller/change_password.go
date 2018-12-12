package controller

import (
	"context"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/service"
)

func ForgetPassword() {

}

func ChangePassword(ctx context.Context, changePasswordModel model.ChangePasswordModel) (bool, error) {
	_, err := service.ChangePassword(ctx, changePasswordModel)
	if err != nil {
		return false, err
	}
	return true, nil
}
