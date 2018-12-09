package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	parts := strings.Split(file.Filename, ".")
	tag := parts[len(parts)-1]
	file.Filename = time.Now().Format("20060102150405") + "." + tag
	err = c.SaveUploadedFile(file, "./assets/crop/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}
