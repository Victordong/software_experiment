package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/service"
	"software_experiment/pkg/web/controller"
	"software_experiment/pkg/web/plugin"
)

func GetIdentifyID(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	captchaModel, err := controller.GetIdentifyID(ctx)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":       http.StatusOK,
			"message":    err,
			"captcha_id": captchaModel.CaptchaId,
		})
		return
	}
}

func VerifyCaptcha(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	var captchaModel model.CaptchaModel
	err := c.ShouldBindJSON(&captchaModel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	_, err = service.VerifyCaptcha(ctx, captchaModel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":       http.StatusOK,
			"message":    err,
			"captcha_id": captchaModel.CaptchaId,
		})
		return
	}
}
