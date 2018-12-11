package controller

import (
	"context"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/service"
)

func GetIdentifyID(ctx context.Context) (*model.CaptchaModel, error) {
	captchaId, err := service.GetIdentifyID(ctx)
	if err != nil {
		return nil, nil
	}
	return captchaId, nil
}

func VerifyCaptcha(ctx context.Context, captchaModel model.CaptchaModel) (bool, error) {
	ifSuccess, err := service.VerifyCaptcha(ctx, captchaModel)
	if err != nil {
		return false, nil
	}
	return ifSuccess, nil
}
