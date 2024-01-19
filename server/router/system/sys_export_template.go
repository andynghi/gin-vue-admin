package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysExportTemplateRouter struct {
}

// InitSysExportTemplateRouter initialization export template routing information
func (s *SysExportTemplateRouter) InitSysExportTemplateRouter(Router *gin.RouterGroup) {
	sysExportTemplateRouter := Router.Group("sysExportTemplate").Use(middleware.OperationRecord())
	sysExportTemplateRouterWithoutRecord := Router.Group("sysExportTemplate")
	var sysExportTemplateApi = v1.ApiGroupApp.SystemApiGroup.SysExportTemplateApi
	{
		sysExportTemplateRouter.POST("createSysExportTemplate", sysExportTemplateApi.CreateSysExportTemplate)             // Create a new export template
		sysExportTemplateRouter.DELETE("deleteSysExportTemplate", sysExportTemplateApi.DeleteSysExportTemplate)           // Delete the export template
		sysExportTemplateRouter.DELETE("deleteSysExportTemplateByIds", sysExportTemplateApi.DeleteSysExportTemplateByIds) // Batch delete export templates
		sysExportTemplateRouter.PUT("updateSysExportTemplate", sysExportTemplateApi.UpdateSysExportTemplate)              // Update export template
		sysExportTemplateRouter.POST("importExcel", sysExportTemplateApi.ImportExcel)                                     // Update export template
	}
	{
		sysExportTemplateRouterWithoutRecord.GET("findSysExportTemplate", sysExportTemplateApi.FindSysExportTemplate)       // Get the export template based on ID
		sysExportTemplateRouterWithoutRecord.GET("getSysExportTemplateList", sysExportTemplateApi.GetSysExportTemplateList) // Get the export template list
		sysExportTemplateRouterWithoutRecord.GET("exportExcel", sysExportTemplateApi.ExportExcel)                           // Export table
		sysExportTemplateRouterWithoutRecord.GET("exportTemplate", sysExportTemplateApi.ExportTemplate)                     // Export table template
	}
}
