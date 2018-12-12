package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"software_experiment/pkg/comm/model"
)

func DummyMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		currentUser, ok := c.Get("currentUser")
		currentUserModel := currentUser.(*model.UserModel)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "无法获得当前用户",
			})
			c.Abort()
		}
		if currentUserModel.Role != "admin" {
			if v := ctx.Value("filterMap"); v != nil {
				filterMapModel := v
				filterMap := filterMapModel.(map[string]interface{})
				filterMap["username"] = currentUserModel.Username
				c.Set("filterMap", filterMap)
			} else {
				filterMap := make(map[string]interface{})
				filterMap["username"] = currentUserModel.Username
				c.Set("filterMap", filterMap)
			}
		}
		c.Next()

	}

}
