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

func GetSupplyCollectionByIdHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	supplyCollectionIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	supplyCollectionId := uint(supplyCollectionIdInt)
	supplyCollection, err := controller.GetSupplyCollectionById(ctx, supplyCollectionId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":             http.StatusOK,
			"message":          err,
			"supplyCollection": supplyCollection,
		})
		return
	}
}

func QuerySupplyCollectionsHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	queryMap := c.Request.URL.Query()
	supplyCollections, num, err := controller.QuerySupplyCollections(ctx, queryMap)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":              http.StatusOK,
			"message":           err,
			"supplyCollections": supplyCollections,
			"total":             num,
		})
		return
	}
}

func NewSupplyCollectionHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	supplyCollection := &model.SupplyCollectionModel{}
	err := c.ShouldBindJSON(supplyCollection)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "信息填写错误",
		})
		return
	}
	supplyCollection, err = controller.PostSupplyCollection(ctx, supplyCollection)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":             http.StatusOK,
			"message":          err,
			"supplyCollection": supplyCollection,
		})
		return
	}
}

func ChangeSupplyCollectionHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	supplyCollectionIdInt, err := strconv.Atoi(c.Param("id"))
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
	supplyCollectionId := uint(supplyCollectionIdInt)
	_, err = controller.GetSupplyCollectionById(ctx, supplyCollectionId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	}
	err = controller.PutSupplyCollection(ctx, supplyCollectionId, updateMap)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":             http.StatusOK,
			"message":          err,
			"supplyCollection": nil,
		})
		return
	}
}

func DeleteSupplyCollectionHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	supplyCollectionIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	supplyCollectionId := uint(supplyCollectionIdInt)
	_, err = controller.DeleteSupplyCollection(ctx, supplyCollectionId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":             http.StatusOK,
			"message":          err,
			"supplyCollection": nil,
		})
		return
	}
}
