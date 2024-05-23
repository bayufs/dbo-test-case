package routers

import (
	"dbo-test-case/app/controllers"
	"dbo-test-case/app/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	customerController := controllers.NewCustomerController()
	orderController := controllers.NewOrderController()
	authController := controllers.NewAuthController()
	api := r.Group("/")

	rootPrefix := api.Group("/v1")

	apiDashboard := rootPrefix.Group("/management")
	{
		customerPrefix := apiDashboard.Group("/customer")
		{
			customerPrefix.GET("/", middlewares.ValidateToken(), customerController.Index)
			customerPrefix.GET("/:customerID", middlewares.ValidateToken(), customerController.Show)
			customerPrefix.POST("/", customerController.Store)
			customerPrefix.PATCH("/:customerID", middlewares.ValidateToken(), customerController.Update)
			customerPrefix.DELETE("/:customerID", middlewares.ValidateToken(), customerController.Delete)
		}

		orderPrefix := apiDashboard.Group("/order").Use(middlewares.ValidateToken())
		{
			orderPrefix.GET("/", orderController.Index)
			orderPrefix.GET("/:orderID", orderController.Show)
			orderPrefix.POST("/", orderController.Store)
			orderPrefix.PATCH("/:orderID", orderController.Update)
			orderPrefix.DELETE("/:orderID", orderController.Delete)
		}
	}

	apiAuth := rootPrefix.Group("/auth")
	{
		apiAuth.POST("/login", authController.Login)
	}

}
