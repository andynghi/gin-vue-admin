package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DictionaryRouter struct{}

func (s *DictionaryRouter) InitSysDictionaryRouter(Router *gin.RouterGroup) {
	sysDictionaryRouter := Router.Group("sysDictionary").Use(middleware.OperationRecord())
	sysDictionaryRouterWithoutRecord := Router.Group("sysDictionary")
	sysDictionaryApi := v1.ApiGroupApp.SystemApiGroup.DictionaryApi
	{
		sysDictionaryRouter.POST("createSysDictionary", sysDictionaryApi.CreateSysDictionary)   // Create a new SysDictionary
		sysDictionaryRouter.DELETE("deleteSysDictionary", sysDictionaryApi.DeleteSysDictionary) // Delete SysDictionary
		sysDictionaryRouter.PUT("updateSysDictionary", sysDictionaryApi.UpdateSysDictionary)    // Update SysDictionary
	}
	{
		sysDictionaryRouterWithoutRecord.GET("findSysDictionary", sysDictionaryApi.FindSysDictionary)       // Get SysDictionary based on ID
		sysDictionaryRouterWithoutRecord.GET("getSysDictionaryList", sysDictionaryApi.GetSysDictionaryList) // Get the SysDictionary list
	}
}
