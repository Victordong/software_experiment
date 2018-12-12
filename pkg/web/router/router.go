package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"software_experiment/pkg/web/handler"
	"software_experiment/pkg/web/middleware"
)

func GetRouter() *gin.Engine {

	router := gin.Default()
	{
		router.POST("/login", middleware.AuthMiddleware.LoginHandler)
		router.GET("/identify", handler.GetIdentifyID)
		router.POST("/identify", handler.VerifyCaptcha)
		router.POST("/users", handler.NewUserHandler)

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

		api.GET("/exhibitions", middleware.RolesFilterMidlle(handler.QueryExhibitionsHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/exhibitions/:id", middleware.RolesFilterMidlle(handler.GetExhibitionByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.DELETE("/exhibitions/:id", middleware.RolesFilterMidlle(handler.DeleteExhibitionHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/exhibitions", middleware.RolesFilterMidlle(handler.NewExhibitionHandler, []string{"admin", "owner", "operator", "collector"}))
		api.PUT("/exhibitions/:id", middleware.RolesFilterMidlle(handler.ChangeExhibitionHandler, []string{"admin", "owner", "operator", "collector"}))

		api.GET("/exhibition_collections/:id", middleware.RolesFilterMidlle(handler.GetExhibitionCollectionByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/exhibition_collections", middleware.RolesFilterMidlle(handler.QueryExhibitionCollectionsHandler, []string{"admin", "owner", "operator", "collector"}))
		api.DELETE("/exhibition_collections/:id", middleware.RolesFilterMidlle(handler.DeleteExhibitionCollectionHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/exhibition_collections", middleware.RolesFilterMidlle(handler.NewExhibitionCollectionHandler, []string{"admin", "owner", "operator", "collector"}))
		api.PUT("/exhibition_collections/:id", middleware.RolesFilterMidlle(handler.ChangeExhibitionCollectionHandler, []string{"admin", "owner", "operator", "collector"}))

		api.GET("/exhibitions_comments", middleware.RolesFilterMidlle(handler.QueryExhibitionCommentsHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/exhibitions_comments/:id", middleware.RolesFilterMidlle(handler.GetExhibitionCommentByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.DELETE("/exhibitions_comments/:id", middleware.RolesFilterMidlle(handler.DeleteExhibitionCommentHandler, []string{"admin", "owner"}))
		api.POST("/exhibitions_comments", middleware.RolesFilterMidlle(handler.NewExhibitionCommentHandler, []string{"admin", "owner"}))
		api.PUT("/exhibitions_comments/:id", middleware.RolesFilterMidlle(handler.ChangeExhibitionCommentHandler, []string{"admin", "owner"}))

		api.GET("/imformations", middleware.RolesFilterMidlle(handler.QueryInformationsHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/imformations/:id", middleware.RolesFilterMidlle(handler.GetInformationByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/imformations", middleware.RolesFilterMidlle(handler.NewInformationHandler, []string{"admin", "owner", "collector"}))
		api.PUT("/imformations/:id", middleware.RolesFilterMidlle(handler.ChangeInformationHandler, []string{"admin", "owner", "collector"}))
		api.DELETE("/imformations/:id", middleware.RolesFilterMidlle(handler.DeleteInformationHandler, []string{"admin", "owner", "collector"}))

		api.GET("/information_collections/:id", middleware.RolesFilterMidlle(handler.GetInformationCollectionByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.PUT("/information_collections/:id", middleware.RolesFilterMidlle(handler.ChangeInformationCollectionHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/information_collections", middleware.RolesFilterMidlle(handler.QueryInformationCollectionsHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/information_collections", middleware.RolesFilterMidlle(handler.NewInformationCollectionHandler, []string{"admin", "owner", "operator", "collector"}))
		api.DELETE("/information_collections/:id", middleware.RolesFilterMidlle(handler.DeleteInformationCollectionHandler, []string{"admin", "owner", "operator", "collector"}))

		api.GET("/information_comments/:id", middleware.RolesFilterMidlle(handler.GetInformationCommentByIdHandler, []string{"admin", "owner", "operator", "collector"}))
		api.PUT("/information_comments/:id", middleware.RolesFilterMidlle(handler.ChangeInformationCommentHandler, []string{"admin", "owner", "operator", "collector"}))
		api.GET("/information_comments", middleware.RolesFilterMidlle(handler.QueryInformationCommentsHandler, []string{"admin", "owner", "operator", "collector"}))
		api.POST("/information_comments", middleware.RolesFilterMidlle(handler.NewInformationCommentHandler, []string{"admin"}))
		api.DELETE("/information_comments/:id", middleware.RolesFilterMidlle(handler.DeleteInformationCommentHandler, []string{"admin"}))

		api.GET("/supplies/:id", middleware.RolesFilterMidlle(handler.GetSupplyByIdHandler, []string{"admin", "owner"}))
		api.PUT("/supplies/:id", middleware.RolesFilterMidlle(handler.ChangeSupplyHandler, []string{"admin"}))
		api.GET("/supplies", middleware.RolesFilterMidlle(handler.QuerySupplysHandler, []string{"admin", "owner"}))
		api.POST("/supplies", middleware.RolesFilterMidlle(handler.NewSupplyHandler, []string{"admin"}))
		api.DELETE("/supplies/:id", middleware.RolesFilterMidlle(handler.DeleteSupplyHandler, []string{"admin"}))

		api.GET("/supply_comments/:id", middleware.RolesFilterMidlle(handler.GetSupplyCommentByIdHandler, []string{"admin", "owner"}))
		api.PUT("/supply_comments/:id", middleware.RolesFilterMidlle(handler.ChangeSupplyCommentHandler, []string{"admin"}))
		api.GET("/supply_comments", middleware.RolesFilterMidlle(handler.QuerySupplyCommentsHandler, []string{"admin", "owner"}))
		api.POST("/supply_comments", middleware.RolesFilterMidlle(handler.NewSupplyCommentHandler, []string{"admin"}))
		api.DELETE("/supply_comments/:id", middleware.RolesFilterMidlle(handler.DeleteSupplyCommentHandler, []string{"admin"}))

		api.GET("/supply_collections/:id", middleware.RolesFilterMidlle(handler.GetSupplyCollectionByIdHandler, []string{"admin", "owner"}))
		api.PUT("/supply_collections/:id", middleware.RolesFilterMidlle(handler.ChangeSupplyCollectionHandler, []string{"admin"}))
		api.GET("/supply_collections", middleware.RolesFilterMidlle(handler.QuerySupplyCollectionsHandler, []string{"admin", "owner"}))
		api.POST("/supply_collections", middleware.RolesFilterMidlle(handler.NewSupplyCollectionHandler, []string{"admin"}))
		api.DELETE("/supply_collections/:id", middleware.RolesFilterMidlle(handler.DeleteSupplyCollectionHandler, []string{"admin"}))

		api.GET("/users/:username", middleware.RolesFilterMidlle(handler.GetUserByUsernameHandler, []string{"admin", "owner"}))
		api.PUT("/users/:username", middleware.RolesFilterMidlle(handler.ChangeUserHandler, []string{"admin", "owner"}))
		api.GET("/users", middleware.RolesFilterMidlle(handler.QueryUsersHandler, []string{"admin", "owner"}))
		api.DELETE("/users/:username", middleware.RolesFilterMidlle(handler.DeleteUserHandler, []string{"admin", "owner"}))
		api.PUT("/change_password", middleware.RolesFilterMidlle(handler.ChangePassWord, []string{"admin", "owner", "operator", "collector"}))

		api.POST("/images/icons", handler.UploadIconFile)
	}
	router.Static("/static", "./assets")
	router.Use(cors.Default())
	return router
}
