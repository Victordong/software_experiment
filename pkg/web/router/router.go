package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/handler"
	"software_experiment/pkg/web/middleware"
)

func GetRouter() *gin.Engine {

	router := gin.Default()
	{
		router.POST("/login", middleware.AuthMiddleware.LoginHandler)
	}

	auth := router.Group("/auth")
	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
		auth.GET("/current_user", handler.GetCurrentUserHandler)
	}

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{

		api.GET("/crops", middleware.RolesFilterMidlle(handler.QueryCropsHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/crops/:id", middleware.RolesFilterMidlle(handler.GetCropByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.DELETE("/crops/:id", middleware.RolesFilterMidlle(handler.DeleteCropHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/crops", middleware.RolesFilterMidlle(handler.NewCropHandler, []string{"admin", "owner", "operator", "collector"}))
		api.PUT("/crops/:id", middleware.RolesFilterMidlle(handler.ChangeCropHandler, []string{"admin", "owner", "operator", "collector"}))

		api.GET("/farmlands/:id", middleware.RolesFilterMidlle(handler.GetFarmlandHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/farmlands", middleware.RolesFilterMidlle(handler.QueryFarmlandsHandler, []string{"admin", "owner", "operator", "collector"}))
		api.DELETE("/farmlands/:id", middleware.RolesFilterMidlle(handler.DeleteFarmlandHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/farmlands", middleware.RolesFilterMidlle(handler.NewFarmlandHandler, []string{"admin", "owner", "operator", "collector"}))
		api.PUT("/farmlands/:id", middleware.RolesFilterMidlle(handler.ChangeFarmlandHandler, []string{"admin", "owner", "operator", "collector"}))

		api.GET("/farmers", middleware.RolesFilterMidlle(handler.QueryFarmersHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/farmers/:id", middleware.RolesFilterMidlle(handler.GetFarmerByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.DELETE("/farmers/:id", middleware.RolesFilterMidlle(handler.DeleteFarmerHandler, []string{"admin", "owner"}))
		api.POST("/farmers", middleware.RolesFilterMidlle(handler.NewFarmerHandler, []string{"admin", "owner"}))
		api.PUT("/farmers/:id", middleware.RolesFilterMidlle(handler.ChangeFarmerHandler, []string{"admin", "owner"}))

		api.GET("/farm_types", middleware.RolesFilterMidlle(handler.QueryFarmTypesHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/farm_types/:id", middleware.RolesFilterMidlle(handler.GetFarmTypeByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/farm_types", middleware.RolesFilterMidlle(handler.NewFarmTypeHandler, []string{"admin", "owner", "collector"}))
		api.PUT("/farm_types/:id", middleware.RolesFilterMidlle(handler.ChangeFarmTypeHandler, []string{"admin", "owner", "collector"}))
		api.DELETE("/farm_types/:id", middleware.RolesFilterMidlle(handler.DeleteFarmTypeHandler, []string{"admin", "owner", "collector"}))

		api.GET("/fertilizers/:id", middleware.RolesFilterMidlle(handler.GetFertilizerByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.PUT("/fertilizers/:id", middleware.RolesFilterMidlle(handler.ChangeFertilizerHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/fertilizers", middleware.RolesFilterMidlle(handler.QueryFertilizersHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/fertilizers", middleware.RolesFilterMidlle(handler.NewFertilizerHandler, []string{"admin", "owner", "operator", "collector"}))
		api.DELETE("/fertilizers/:id", middleware.RolesFilterMidlle(handler.DeleteFertilizerHandler, []string{"admin", "owner", "operator", "collector"}))

		api.GET("/shops/:id", middleware.RolesFilterMidlle(handler.GetShopByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.PUT("/shops/:id", middleware.RolesFilterMidlle(handler.ChangeShopHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/shops", middleware.RolesFilterMidlle(handler.QueryShopsHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/shops", middleware.RolesFilterMidlle(handler.NewShopHandler, []string{"admin"}))
		api.DELETE("/shops/:id", middleware.RolesFilterMidlle(handler.DeleteShopHandler, []string{"admin"}))

		api.GET("/roles/:id", middleware.RolesFilterMidlle(handler.GetRoleByIdHandler, []string{"admin", "owner"}))
		api.PUT("/roles/:id", middleware.RolesFilterMidlle(handler.ChangeRoleHandler, []string{"admin"}))
		api.GET("/roles", middleware.RolesFilterMidlle(handler.QueryRolesHandler, []string{"admin", "owner"}))
		api.POST("/roles", middleware.RolesFilterMidlle(handler.NewRoleHandler, []string{"admin"}))
		api.DELETE("/roles/:id", middleware.RolesFilterMidlle(handler.DeleteRoleHandler, []string{"admin"}))

		api.GET("/operators/:username", middleware.RolesFilterMidlle(handler.GetOperatorByUsernameHandler, []string{"admin", "owner"}))
		api.PUT("/operators/:username", middleware.RolesFilterMidlle(handler.ChangeOperatorHandler, []string{"admin", "owner"}))
		api.GET("/operators", middleware.RolesFilterMidlle(handler.QueryOperatorsHandler, []string{"admin", "owner"}))
		api.DELETE("/operators/:username", middleware.RolesFilterMidlle(handler.DeleteOperatorHandler, []string{"admin", "owner"}))
		api.POST("/operators", middleware.RolesFilterMidlle(handler.NewOperatorHandler, []string{"admin", "owner"}))
		api.PUT("/change_password", middleware.RolesFilterMidlle(handler.ChangePassword, []string{"admin", "owner", "operator", "collector"}))

		api.GET("/stock_flows", middleware.RolesFilterMidlle(handler.QueryStockFlowsHandler, []string{"admin", "owner", "operator"}))
		api.POST("/stock_flows", middleware.RolesFilterMidlle(handler.NewStockFlowHandler, []string{"admin", "owner", "operator"}))

		api.GET("/stocks", middleware.RolesFilterMidlle(handler.QueryStocksrHandler, []string{"admin", "owner", "operator"}))
		api.POST("/stocks", middleware.RolesFilterMidlle(handler.NewStockHandler, []string{"admin", "owner", "operator"}))

		api.GET("/consume_record_sums", middleware.RolesFilterMidlle(handler.QueryConsumeRecordSumHandler, []string{"admin", "owner", "operator"}))
		api.GET("/consume_records", middleware.RolesFilterMidlle(handler.QueryConsumeRecordHandler, []string{"admin", "owner", "operator"}))
		api.POST("/consumes", middleware.RolesFilterMidlle(handler.NewConsume, []string{"admin", "owner", "operator"}))
		api.GET("/consume_packages", middleware.RolesFilterMidlle(handler.QueryConsumeRecordPackageHandler, []string{"admin", "owner", "operator"}))

		api.POST("/calculate/getElementMixtures", middleware.RolesFilterMidlle(handler.GetElementMixturesHandler, []string{"admin", "owner", "operator"}))
		api.POST("/calculate/getFertilizerMixtures", middleware.RolesFilterMidlle(handler.GetFertilizerMixturesHandler, []string{"admin", "owner", "operator"}))
		api.POST("/calculate/getFertilizerPrice", middleware.RolesFilterMidlle(handler.GetFertilizersPriceHandler, []string{"admin", "owner", "operator"}))

		api.GET("/fertilizer_kinds", middleware.RolesFilterMidlle(handler.GetFertilizerKindHandler, []string{"admin", "owner", "operator", "collector"}))

		api.POST("/shop_messages", middleware.RolesFilterMidlle(handler.NewShopMessage, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/shop_messages/batch", middleware.RolesFilterMidlle(handler.NewShopMessageBatch, []string{"admin", "owner", "collector"}))
		api.GET("/shop_messages", middleware.RolesFilterMidlle(handler.QueryShopMessage, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/my/shop_messages", middleware.RolesFilterMidlle(handler.QueryMyShopMessage, []string{"admin", "owner", "operator", "collector"}))

		api.POST("/notices", middleware.RolesFilterMidlle(handler.NewNotice, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/notices/batch", middleware.RolesFilterMidlle(handler.NewNoticeBatch, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/notices", middleware.RolesFilterMidlle(handler.QueryNotice, []string{"admin", "owner", "operator", "collector"}))

		api.GET("/farmer_numbers/:id_number", middleware.RolesFilterMidlle(handler.GetFarmerNumberByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/farmer_numbers", middleware.RolesFilterMidlle(handler.NewFarmerNumberHandler, []string{"admin", "owner", "operator"}))
		api.GET("/farmer_numbers", middleware.RolesFilterMidlle(handler.QueryFarmNumbers, []string{"admin", "owner", "operator", "collector"}))
		api.DELETE("/farmer_numbers/:id", middleware.RolesFilterMidlle(handler.DeleteFarmerNumberHandler, []string{"admin", "owner", "operator", "collector"}))

		api.GET("/notice_templates", middleware.RolesFilterMidlle(handler.QueryNoticeTemplates, []string{"admin", "owner", "operator", "collector"}))

		api.GET("/images/crops", middleware.RolesFilterMidlle(handler.GetCropImagesHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/images/crops", handler.UploadFile)
	}
	router.Static("/static", "./assets")
	router.Use(cors.Default())
	return router
}
