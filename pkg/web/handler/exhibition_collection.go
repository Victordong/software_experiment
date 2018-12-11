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

func GetExhibitionCollectionByIdHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	exhibitionCollectionIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	exhibitionCollectionId := uint(exhibitionCollectionIdInt)
	exhibitionCollection, err := controller.GetExhibitionCollectionById(ctx, exhibitionCollectionId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":                 http.StatusOK,
			"message":              err,
			"exhibitionCollection": exhibitionCollection,
		})
		return
	}
}

func QueryExhibitionCollectionsHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	queryMap := c.Request.URL.Query()
	exhibitionCollections, num, err := controller.QueryExhibitionCollections(ctx, queryMap)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":                  http.StatusOK,
			"message":               err,
			"exhibitionCollections": exhibitionCollections,
			"total":                 num,
		})
		return
	}
}

func NewExhibitionCollectionHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	exhibitionCollection := &model.ExhibitionCollectionModel{}
	err := c.ShouldBindJSON(exhibitionCollection)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "信息填写错误",
		})
		return
	}
	exhibitionCollection, err = controller.PostExhibitionCollection(ctx, exhibitionCollection)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":                 http.StatusOK,
			"message":              err,
			"exhibitionCollection": exhibitionCollection,
		})
		return
	}
}

func ChangeExhibitionCollectionHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	exhibitionCollectionIdInt, err := strconv.Atoi(c.Param("id"))
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
	exhibitionCollectionId := uint(exhibitionCollectionIdInt)
	_, err = controller.GetExhibitionCollectionById(ctx, exhibitionCollectionId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	}
	err = controller.PutExhibitionCollection(ctx, exhibitionCollectionId, updateMap)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":                 http.StatusOK,
			"message":              err,
			"exhibitionCollection": nil,
		})
		return
	}
}

func DeleteExhibitionCollectionHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = plugin.SetContext(c, ctx)
	exhibitionCollectionIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	exhibitionCollectionId := uint(exhibitionCollectionIdInt)
	_, err = controller.DeleteExhibitionCollection(ctx, exhibitionCollectionId)
	if err != nil {
		errModel := err.(plugin.CustomErr)
		c.JSON(errModel.StatusCode, gin.H{
			"code": errModel.Code,
			"msg":  errModel.Information,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":                 http.StatusOK,
			"message":              err,
			"exhibitionCollection": nil,
		})
		return
	}
}
