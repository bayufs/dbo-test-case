package routers

import (
	"dbo-test-case/app/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	customerController := controllers.NewCustomerController()

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
		}
	}

}
