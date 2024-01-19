package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type SysExportTemplateApi struct {
}

var sysExportTemplateService = service.ServiceGroupApp.SystemServiceGroup.SysExportTemplateService

// CreateSysExportTemplate creates an export template
// @Tags SysExportTemplate
// @Summary Create export template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysExportTemplate true "Create export template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Created successfully"}"
// @Router /sysExportTemplate/createSysExportTemplate [post]
func (sysExportTemplateApi *SysExportTemplateApi) CreateSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(sysExportTemplate, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysExportTemplateService.CreateSysExportTemplate(&sysExportTemplate); err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
	} else {
		response.OkWithMessage("Created successfully", c)
	}
}

// DeleteSysExportTemplate deletes the export template
// @Tags SysExportTemplate
// @Summary Delete export template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysExportTemplate true "Delete export template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Delete successfully"}"
// @Router /sysExportTemplate/deleteSysExportTemplate [delete]
func (sysExportTemplateApi *SysExportTemplateApi) DeleteSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysExportTemplateService.DeleteSysExportTemplate(sysExportTemplate); err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
	} else {
		response.OkWithMessage("Deletion successful", c)
	}
}

// DeleteSysExportTemplateByIds delete export templates in batches
// @Tags SysExportTemplate
// @Summary Batch delete export templates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete export templates in batches"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Batch deletion successful"}"
// @Router /sysExportTemplate/deleteSysExportTemplateByIds [delete]
func (sysExportTemplateApi *SysExportTemplateApi) DeleteSysExportTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysExportTemplateService.DeleteSysExportTemplateByIds(IDS); err != nil {
		global.GVA_LOG.Error("Batch deletion failed!", zap.Error(err))
		response.FailWithMessage("Batch deletion failed", c)
	} else {
		response.OkWithMessage("Batch deletion successful", c)
	}
}

// UpdateSysExportTemplate updates the export template
// @Tags SysExportTemplate
// @Summary Update export template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysExportTemplate true "Update export template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Update successful"}"
// @Router /sysExportTemplate/updateSysExportTemplate [put]
func (sysExportTemplateApi *SysExportTemplateApi) UpdateSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(sysExportTemplate, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysExportTemplateService.UpdateSysExportTemplate(sysExportTemplate); err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
	} else {
		response.OkWithMessage("Update successful", c)
	}
}

// FindSysExportTemplate uses id to query the export template
// @Tags SysExportTemplate
// @Summary Use id to query and export templates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysExportTemplate true "Query the export template using id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Query successful"}"
// @Router /sysExportTemplate/findSysExportTemplate [get]
func (sysExportTemplateApi *SysExportTemplateApi) FindSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindQuery(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resysExportTemplate, err := sysExportTemplateService.GetSysExportTemplate(sysExportTemplate.ID); err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithMessage("Query failed", c)
	} else {
		response.OkWithData(gin.H{"resysExportTemplate": resysExportTemplate}, c)
	}
}

// GetSysExportTemplateList Gets the export template list in paging
// @Tags SysExportTemplate
// @Summary Get the export template list by pagination
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query systemReq.SysExportTemplateSearch true "Get the export template list in pages"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /sysExportTemplate/getSysExportTemplateList [get]
func (sysExportTemplateApi *SysExportTemplateApi) GetSysExportTemplateList(c *gin.Context) {
	var pageInfo systemReq.SysExportTemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysExportTemplateService.GetSysExportTemplateInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "Get successful", c)
	}
}

// ExportExcel export table
// @Tags SysExportTemplate
// @Summary Export table
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/exportExcel [get]
func (sysExportTemplateApi *SysExportTemplateApi) ExportExcel(c *gin.Context) {
	templateID := c.Query("templateID")
	if templateID == "" {
		response.FailWithMessage("Template ID cannot be empty", c)
		return
	}
	if file, name, err := sysExportTemplateService.ExportExcel(templateID); err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
	} else {
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name+utils.RandomString(6)+".xlsx")) // Rename the downloaded file
		c.Header("success", "true")
		c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", file.Bytes())
	}
}

// ExportExcel export table template
// @Tags SysExportTemplate
// @Summary Export table template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/exportExcel [get]
func (sysExportTemplateApi *SysExportTemplateApi) ExportTemplate(c *gin.Context) {
	templateID := c.Query("templateID")
	if templateID == "" {
		response.FailWithMessage("Template ID cannot be empty", c)
		return
	}
	if file, name, err := sysExportTemplateService.ExportTemplate(templateID); err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
	} else {
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name+"template.xlsx")) // Rename the downloaded file
		c.Header("success", "true")
		c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", file.Bytes())
	}
}

// ExportExcel import table
// @Tags SysImportTemplate
// @Summary Import table
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/importExcel [post]
func (sysExportTemplateApi *SysExportTemplateApi) ImportExcel(c *gin.Context) {
	templateID := c.Query("templateID")
	if templateID == "" {
		response.FailWithMessage("Template ID cannot be empty", c)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("File acquisition failed!", zap.Error(err))
		response.FailWithMessage("File acquisition failed", c)
		return
	}
	if err := sysExportTemplateService.ImportExcel(templateID, file); err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("Import successful", c)
	}
}
