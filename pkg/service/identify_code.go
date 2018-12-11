package service

import (
	"context"
	"github.com/dchest/captcha"
	"github.com/disintegration/imaging"
	"math/rand"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
)

func GetIdentifyID(ctx context.Context) (*model.CaptchaModel, error) {
	captchaId := captcha.New()
	var digits []byte
	for i := 0; i < 6; i++ {
		digits = append(digits, uint8(rand.Intn(10)))
	}
	img := captcha.NewImage(captchaId, digits, 200, 100)
	err := imaging.Save(img, "./assets/identify_code/"+captchaId+".png")
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: err.Error(),
		}
	}
	return &model.CaptchaModel{
		CaptchaId: captchaId,
	}, nil
}

func VerifyCaptcha(ctx context.Context, captchaModel model.CaptchaModel) (bool, error) {
	ifSuccess := captcha.VerifyString(captchaModel.CaptchaId, captchaModel.Result)
	if !ifSuccess {
		return false, plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: "验证失败",
		}
	}
	return true, nil
}
