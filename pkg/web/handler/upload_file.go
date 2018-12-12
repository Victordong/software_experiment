package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"software_experiment/pkg/web/controller"
)

func UploadIconFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "读取异常",
		})
		return
	}
	filename, err := controller.Upload_image(file, "icon/")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "读取异常",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": err,
			"images":  filename,
		})
		return
	}
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "读取异常",
		})
		return
	}
	filename, err := controller.Upload_image(file, "file/")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "读取异常",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": err,
			"images":  filename,
		})
		return
	}
}
