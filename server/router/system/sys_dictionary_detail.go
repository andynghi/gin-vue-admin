package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DictionaryDetailRouter struct{}

func (s *DictionaryDetailRouter) InitSysDictionaryDetailRouter(Router *gin.RouterGroup) {
	dictionaryDetailRouter := Router.Group("sysDictionaryDetail").Use(middleware.OperationRecord())
	dictionaryDetailRouterWithoutRecord := Router.Group("sysDictionaryDetail")
	sysDictionaryDetailApi := v1.ApiGroupApp.SystemApiGroup.DictionaryDetailApi
	{
		dictionaryDetailRouter.POST("createSysDictionaryDetail", sysDictionaryDetailApi.CreateSysDictionaryDetail)   // Create new SysDictionaryDetail
		dictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", sysDictionaryDetailApi.DeleteSysDictionaryDetail) // Delete SysDictionaryDetail
		dictionaryDetailRouter.PUT("updateSysDictionaryDetail", sysDictionaryDetailApi.UpdateSysDictionaryDetail)    // Update SysDictionaryDetail
	}
	{
		dictionaryDetailRouterWithoutRecord.GET("findSysDictionaryDetail", sysDictionaryDetailApi.FindSysDictionaryDetail)       // Get SysDictionaryDetail based on ID
		dictionaryDetailRouterWithoutRecord.GET("getSysDictionaryDetailList", sysDictionaryDetailApi.GetSysDictionaryDetailList) // Get the SysDictionaryDetail list
	}
}
