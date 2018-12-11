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

func GetSupplyCommentByIdHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	supplyCommentIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	supplyCommentId := uint(supplyCommentIdInt)
	supplyComment, err := controller.GetSupplyCommentById(ctx, supplyCommentId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":          http.StatusOK,
			"message":       err,
			"supplyComment": supplyComment,
		})
		return
	}
}

func QuerySupplyCommentsHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	queryMap := c.Request.URL.Query()
	supplyComments, num, err := controller.QuerySupplyComments(ctx, queryMap)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":           http.StatusOK,
			"message":        err,
			"supplyComments": supplyComments,
			"total":          num,
		})
		return
	}
}

func NewSupplyCommentHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	supplyComment := &model.SupplyCommentModel{}
	err := c.ShouldBindJSON(supplyComment)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "信息填写错误",
		})
		return
	}
	supplyComment, err = controller.PostSupplyComment(ctx, supplyComment)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":          http.StatusOK,
			"message":       err,
			"supplyComment": supplyComment,
		})
		return
	}
}

func ChangeSupplyCommentHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	supplyCommentIdInt, err := strconv.Atoi(c.Param("id"))
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
	supplyCommentId := uint(supplyCommentIdInt)
	_, err = controller.GetSupplyCommentById(ctx, supplyCommentId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	}
	err = controller.PutSupplyComment(ctx, supplyCommentId, updateMap)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":          http.StatusOK,
			"message":       err,
			"supplyComment": nil,
		})
		return
	}
}

func DeleteSupplyCommentHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	supplyCommentIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	supplyCommentId := uint(supplyCommentIdInt)
	_, err = controller.DeleteSupplyComment(ctx, supplyCommentId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":          http.StatusOK,
			"message":       err,
			"supplyComment": nil,
		})
		return
	}
}
