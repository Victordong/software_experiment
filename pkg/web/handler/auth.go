package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentUserHandler(c *gin.Context) {
	currentUser, ok := c.Get("currentUser")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "无法获得当前用户",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
		"user": currentUser,
	})
}
