package example

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	exaCustomerApi := v1.ApiGroupApp.ExampleApiGroup.CustomerApi
	{
		customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)   // Create a customer
		customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)    // Update customer
		customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer) // Delete customer
	}
	{
		customerRouterWithoutRecord.GET("customer", exaCustomerApi.GetExaCustomer)         // Get single customer information
		customerRouterWithoutRecord.GET("customerList", exaCustomerApi.GetExaCustomerList) // Get the customer list
	}
}
