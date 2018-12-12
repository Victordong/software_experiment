package service

import (
	"context"
	"github.com/dchest/captcha"
	"github.com/disintegration/imaging"
	"software_experiment/pkg/comm/database"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"time"
)

func B2S(bs []byte) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = v + 48
	}
	return string(b)
}

func SetIdentifyRedis(captchaId string, digits string) (bool, error) {
	err := database.RedisClient.Set(captchaId, digits, time.Minute*3).Err()
	if err != nil {
		return false, plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: err.Error(),
		}
	}
	return true, nil
}

func GetIdentifyRedis(captchaId string) (string, error) {
	result, err := database.RedisClient.Get(captchaId).Result()
	if err != nil {
		return "", plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: err.Error(),
		}
	}
	return result, nil
}
func GetIdentifyID(ctx context.Context) (*model.CaptchaModel, error) {
	captchaId := captcha.New()
	digits := captcha.RandomDigits(captcha.DefaultLen)
	digitStr := B2S(digits)
	img := captcha.NewImage(captchaId, digits, captcha.StdWidth, captcha.StdHeight)
	err := imaging.Save(img, "./assets/identify_code/"+captchaId+".png")
	if err != nil {
		return nil, plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: err.Error(),
		}
	}
	_, err = SetIdentifyRedis(captchaId, digitStr)
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
	result, err := GetIdentifyRedis(captchaModel.CaptchaId)
	if err != nil {
		return false, err
	}
	if result != captchaModel.Result {
		return false, plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: "验证失败",
		}
	}
	return true, nil
}
