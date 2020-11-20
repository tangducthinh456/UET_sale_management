package router

import (
	"SaleManagement/webserver/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		//api.POST("/login")
		//api.GET("/logout")
		user := api.Group("/users")
		{
			user.GET("", controller.HandleGETUsers)
			user.POST("", controller.HandlePOSTUsers)
			user.PUT("/:user", controller.HandlePUTUser)
			user.DELETE("/:user", controller.HandleDisableUser)
			user.POST("/:user", controller.HandleEnableUser)
		}

		provider := api.Group("/providers")
		{
			provider.GET("", controller.HandleGETProviders)
			provider.POST("", controller.HandlePOSTProviders)
			provider.PUT("/:provider", controller.HandlePUTProvider)
			provider.DELETE("/:provider", controller.HandleDisableProvider)
			provider.POST("/:provider", controller.HandleEnableProvider)
		}
	}
	return router
}
