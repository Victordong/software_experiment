package middleware

import (
	"auto_fertilizer_back/pkg/comm/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RolesIn(allow []string, has []string) bool {
	for _, a := range allow {
		for _, h := range has {
			if a == h {
				return true
			}
		}

	}
	return false
}
func RolesFilterMidlle(handle gin.HandlerFunc, roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, ok := c.Get("currentUser")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "未登录",
			})
			return
		}
		currentUserModel := currentUser.(*model.OperatorModel)
		if RolesIn(roles, currentUserModel.RoleNames) {
			handle(c)
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "权限不足",
			})
		}
	}
}
