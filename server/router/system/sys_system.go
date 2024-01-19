package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		sysRouter.POST("getSystemConfig", systemApi.GetSystemConfig) // Get the configuration file content
		sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig) //Set the configuration file content
		sysRouter.POST("getServerInfo", systemApi.GetServerInfo)     // Get server information
		sysRouter.POST("reloadSystem", systemApi.ReloadSystem)       // Restart the service
	}
}
