package controller

import (
	"auto_fertilizer_back/pkg/comm/model"
	"auto_fertilizer_back/pkg/web/plugin"
	"context"
	"crypto/sha256"
	"fmt"
)

var salt = "strong#"

func GenPasswordHash(ctx context.Context, password string) string {
	sum := sha256.Sum256([]byte(password + salt))
	return fmt.Sprintf("%x", sum)
}

func HasRoot(ctx context.Context, shop_id uint) (bool, error) {
	currentUserValue := ctx.Value("currentUser")
	var currentUser *model.OperatorModel
	switch currentUserValue.(type) {
	case *model.OperatorModel:
		currentUser = currentUserValue.(*model.OperatorModel)
	default:
		return false, plugin.CustomErr{
			Code:        404,
			StatusCode:  404,
			Information: "当前用户未找到",
		}
	}
	if currentUser.ShopId == 1 || currentUser.ShopId == shop_id {
		return true, nil
	}
	return false, plugin.CustomErr{
		Code:        500,
		StatusCode:  200,
		Information: "当前用户没有修改的权限",
	}
}
