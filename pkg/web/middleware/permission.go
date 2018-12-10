package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"software_experiment/pkg/comm/model"
)

func RolesIn(allow []string, has string) bool {
	for _, a := range allow {
		if a == has {
			return true
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
		currentUserModel := currentUser.(*model.UserModel)
		if RolesIn(roles, currentUserModel.Role) {
			handle(c)
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "权限不足",
			})
		}
	}
}
