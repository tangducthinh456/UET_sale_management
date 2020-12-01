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

		customer := api.Group("/customers")
		{
			customer.GET("", controller.HandleGETCustomers)
			customer.POST("", controller.HandlePOSTCustomers)
			customer.PUT("/:customer", controller.HandlePUTCustomer)
			customer.DELETE("/:customer", controller.HandleDisableCustomer)
			customer.POST("/:customer", controller.HandleEnableCustomer)
		}

		group := api.Group("/groups")
		{
			group.GET("", controller.HandleGetGroups)
			group.POST("", controller.HandlePostGroup)
		}

		product := api.Group("/products")
		{
			product.GET("", controller.HandleGETProducts)
			product.POST("", controller.HandlePOSTProducts)
			product.PUT("/:product", controller.HandlePUTProduct)
			product.DELETE("/:product", controller.HandleDisableProduct)
			product.POST("/:product", controller.HandleEnableProduct)
		}

		bill := api.Group("/bills")
		{
			bill.GET("", controller.HandleGETBills)
			bill.POST("", controller.HandlePOSTBills)
			bill.PUT("/:bill", controller.HandlePUTBill)
			bill.DELETE("/:bill", controller.HandleDeleteBill)
		}

		imp := api.Group("/imports")
		{
			imp.GET("", controller.HandleGETImports)
			imp.POST("", controller.HandlePOSTImports)
			imp.PUT("/:import", controller.HandlePUTImport)
			imp.DELETE("/:import", controller.HandleDeleteImport)
		}

	}
	return router
}
