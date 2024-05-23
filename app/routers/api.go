package routers

import (
	"dbo-test-case/app/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	customerController := controllers.NewCustomerController()
	orderController := controllers.NewOrderController()

	api := r.Group("/")

	rootPrefix := api.Group("/v1")

	apiDashboard := rootPrefix.Group("/management")
	{
		customerPrefix := apiDashboard.Group("/customer")
		{
			customerPrefix.GET("/", customerController.Index)
			customerPrefix.GET("/:customerID", customerController.Show)
			customerPrefix.POST("/", customerController.Store)
			customerPrefix.PATCH("/:customerID", customerController.Update)
			customerPrefix.DELETE("/:customerID", customerController.Delete)
		}

		orderPrefix := apiDashboard.Group("/order")
		{
			orderPrefix.GET("/", orderController.Index)
			orderPrefix.GET("/:orderID", orderController.Show)
			orderPrefix.POST("/", orderController.Store)
			orderPrefix.PATCH("/:orderID", orderController.Update)
			orderPrefix.DELETE("/:orderID", orderController.Delete)
		}
	}

}
