package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")

	apiPublicRouterWithoutRecord := RouterPub.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)               // Create Api
		apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // Delete Api
		apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // Get a single Api message
		apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // Update api
		apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // Delete the selected api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", apiRouterApi.GetAllApis) // Get all apis
		apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // Get the Api list
	}
	{
		apiPublicRouterWithoutRecord.GET("freshCasbin", apiRouterApi.FreshCasbin) // Refresh casbin permissions
	}
}
