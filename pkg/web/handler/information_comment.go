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

func GetInformationCommentByIdHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	informationCommentIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	informationCommentId := uint(informationCommentIdInt)
	informationComment, err := controller.GetInformationCommentById(ctx, informationCommentId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":               http.StatusOK,
			"message":            err,
			"informationComment": informationComment,
		})
		return
	}
}

func QueryInformationCommentsHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	queryMap := c.Request.URL.Query()
	informationComments, num, err := controller.QueryInformationComments(ctx, queryMap)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":                http.StatusOK,
			"message":             err,
			"informationComments": informationComments,
			"total":               num,
		})
		return
	}
}

func NewInformationCommentHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	informationComment := &model.InformationCommentModel{}
	err := c.ShouldBindJSON(informationComment)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "信息填写错误",
		})
		return
	}
	informationComment, err = controller.PostInformationComment(ctx, informationComment)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":               http.StatusOK,
			"message":            err,
			"informationComment": informationComment,
		})
		return
	}
}

func ChangeInformationCommentHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	informationCommentIdInt, err := strconv.Atoi(c.Param("id"))
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
	informationCommentId := uint(informationCommentIdInt)
	_, err = controller.GetInformationCommentById(ctx, informationCommentId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	}
	err = controller.PutInformationComment(ctx, informationCommentId, updateMap)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":               http.StatusOK,
			"message":            err,
			"informationComment": nil,
		})
		return
	}
}

func DeleteInformationCommentHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	informationCommentIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	informationCommentId := uint(informationCommentIdInt)
	_, err = controller.DeleteInformationComment(ctx, informationCommentId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":               http.StatusOK,
			"message":            err,
			"informationComment": nil,
		})
		return
	}
}
