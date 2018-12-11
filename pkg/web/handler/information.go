package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/controller"
	"software_experiment/pkg/web/plugin"
	"strconv"
)

func GetInformationByIdHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	informationIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	informationId := uint(informationIdInt)
	information, err := controller.GetInformationById(ctx, informationId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	}
	err = plugin.ChangeVisitNumRedis("information", informationId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"message":     err,
		"information": information,
	})
	return

}

func QueryInformationsHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	queryMap := c.Request.URL.Query()
	informations, num, err := controller.QueryInformations(ctx, queryMap)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":         http.StatusOK,
			"message":      err,
			"informations": informations,
			"total":        num,
		})
		return
	}
}

func NewInformationHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	information := &model.InformationModel{}
	err := c.ShouldBindJSON(information)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "信息填写错误",
		})
		return
	}
	information, err = controller.PostInformation(ctx, information)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":        http.StatusOK,
			"message":     err,
			"information": information,
		})
		return
	}
}

func ChangeInformationHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	informationIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	updateMap := make(map[string]interface{})
	err = c.BindJSON(&updateMap)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	informationId := uint(informationIdInt)
	_, err = controller.GetInformationById(ctx, informationId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	}
	err = controller.PutInformation(ctx, informationId, updateMap)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":        http.StatusOK,
			"message":     err,
			"information": nil,
		})
		return
	}
}

func DeleteInformationHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	informationIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	informationId := uint(informationIdInt)
	_, err = controller.DeleteInformation(ctx, informationId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":        http.StatusOK,
			"message":     err,
			"information": nil,
		})
		return
	}
}
