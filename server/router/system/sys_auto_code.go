package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	autoCodeApi := v1.ApiGroupApp.SystemApiGroup.AutoCodeApi
	{
		autoCodeRouter.GET("getDB", autoCodeApi.GetDB)                  // Get the database
		autoCodeRouter.GET("getTables", autoCodeApi.GetTables)          // Get the tables of the corresponding database
		autoCodeRouter.GET("getColumn", autoCodeApi.GetColumn)          // Get all field information of the specified table
		autoCodeRouter.POST("preview", autoCodeApi.PreviewTemp)         // Get automatically created code preview
		autoCodeRouter.POST("createTemp", autoCodeApi.CreateTemp)       // Create automation code
		autoCodeRouter.POST("createPackage", autoCodeApi.CreatePackage) // Create package
		autoCodeRouter.POST("getPackage", autoCodeApi.GetPackage)       // Get the package
		autoCodeRouter.POST("delPackage", autoCodeApi.DelPackage)       // Delete package
		autoCodeRouter.POST("createPlug", autoCodeApi.AutoPlug)         // Automatic plug-in package template
		autoCodeRouter.POST("installPlugin", autoCodeApi.InstallPlugin) // Automatically install plug-ins
		autoCodeRouter.POST("pubPlug", autoCodeApi.PubPlug)             // Package plug-in
	}
}
