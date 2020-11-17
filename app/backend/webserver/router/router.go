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
			//user.POST("")
			//user.PUT("/:id")
			//user.DELETE("/:id")
		}
	}
	return router
}
