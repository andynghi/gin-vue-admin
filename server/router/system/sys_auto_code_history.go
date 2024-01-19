package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AutoCodeHistoryRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeHistoryRouter(Router *gin.RouterGroup) {
	autoCodeHistoryRouter := Router.Group("autoCode")
	autoCodeHistoryApi := v1.ApiGroupApp.SystemApiGroup.AutoCodeHistoryApi
	{
		autoCodeHistoryRouter.POST("getMeta", autoCodeHistoryApi.First)         // Get meta information based on id
		autoCodeHistoryRouter.POST("rollback", autoCodeHistoryApi.RollBack)     // Rollback
		autoCodeHistoryRouter.POST("delSysHistory", autoCodeHistoryApi.Delete)  // Delete rollback records
		autoCodeHistoryRouter.POST("getSysHistory", autoCodeHistoryApi.GetList) // Get rollback record pagination
	}
}
