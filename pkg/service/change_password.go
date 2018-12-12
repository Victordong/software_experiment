package service

import (
	"context"
	"crypto/rand"
	"github.com/jinzhu/gorm"
	"io"
	db "software_experiment/pkg/comm/database"
	"software_experiment/pkg/comm/manager"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/plugin"
	"time"
)

const idLen = 20

// idChars are characters allowed in captcha id.
var idChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func randomBytes(length int) (b []byte) {
	b = make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("captcha: error reading random source: " + err.Error())
	}
	return
}

func randomBytesMod(length int, mod byte) (b []byte) {
	if length == 0 {
		return nil
	}
	if mod == 0 {
		panic("captcha: bad mod argument for randomBytesMod")
	}
	maxrb := 255 - byte(256%int(mod))
	b = make([]byte, length)
	i := 0
	for {
		r := randomBytes(length + (length / 4))
		for _, c := range r {
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = c % mod
			i++
			if i == length {
				return
			}
		}
	}

}

func randomId() string {
	b := randomBytesMod(idLen, byte(len(idChars)))
	for i, c := range b {
		b[i] = idChars[c]
	}
	return string(b)
}

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
	err := db.RedisClient.Set(session, username, 3*time.Minute).Err()
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

func SessionEmail(ctx context.Context, forgetPasswordModel model.ForgetPasswordModel) (bool, error) {
	session := randomId()
	err := db.RedisClient.Set(session, forgetPasswordModel.Username, time.Minute*3).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func ForgetPassword(ctx context.Context, forgetPasswordModel model.ForgetPasswordModel) (bool, error) {
	username, err := db.RedisClient.Get(forgetPasswordModel.Session).Result()
	if err != nil {
		return false, plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: err.Error(),
		}
	}
	_, err = manager.ChangePassword(ctx, username, forgetPasswordModel.Password)
	if err != nil {
		return false, plugin.CustomErr{
			Code:        500,
			StatusCode:  200,
			Information: err.Error(),
		}
	}
	return true, nil
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
		_, err = manager.ChangePassword(ctx, changePasswordModel.Username, changePasswordModel.Password)
		if err != nil {
			return false, plugin.CustomErr{
				Code:        500,
				StatusCode:  200,
				Information: err.Error(),
			}
		}

	} else {
		ctxValue := ctx.Value("currentUser")
		var currentUser *model.UserModel
		switch ctxValue.(type) {
		case *gorm.DB:
			currentUser = ctxValue.(*model.UserModel)
		default:
			return false, plugin.CustomErr{
				Code:        500,
				StatusCode:  200,
				Information: "未找到用户",
			}
		}
		_, err := manager.ChangePassword(ctx, currentUser.Username, changePasswordModel.Password)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
