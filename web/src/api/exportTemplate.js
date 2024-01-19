import service from '@/utils/request'

// @Tags SysExportTemplate
// @Summary Create export template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysExportTemplate true "Create export template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Created successfully"}"
// @Router /sysExportTemplate/createSysExportTemplate [post]
export const createSysExportTemplate = (data) => {
  return service({
    url: '/sysExportTemplate/createSysExportTemplate',
    method: 'post',
    data
  })
}

// @Tags SysExportTemplate
// @Summary Delete export template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysExportTemplate true "Delete export template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Delete successfully"}"
// @Router /sysExportTemplate/deleteSysExportTemplate [delete]
export const deleteSysExportTemplate = (data) => {
  return service({
    url: '/sysExportTemplate/deleteSysExportTemplate',
    method: 'delete',
    data
  })
}

// @Tags SysExportTemplate
// @Summary Batch delete export templates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete export templates in batches"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Delete successfully"}"
// @Router /sysExportTemplate/deleteSysExportTemplate [delete]
export const deleteSysExportTemplateByIds = (data) => {
  return service({
    url: '/sysExportTemplate/deleteSysExportTemplateByIds',
    method: 'delete',
    data
  })
}

// @Tags SysExportTemplate
// @Summary Update export template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysExportTemplate true "Update export template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Update successful"}"
// @Router /sysExportTemplate/updateSysExportTemplate [put]
export const updateSysExportTemplate = (data) => {
  return service({
    url: '/sysExportTemplate/updateSysExportTemplate',
    method: 'put',
    data
  })
}

// @Tags SysExportTemplate
// @Summary Use id to query and export templates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysExportTemplate true "Query the export template with id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Query successful"}"
// @Router /sysExportTemplate/findSysExportTemplate [get]
export const findSysExportTemplate = (params) => {
  return service({
    url: '/sysExportTemplate/findSysExportTemplate',
    method: 'get',
    params
  })
}

// @Tags SysExportTemplate
// @Summary Get the export template list by pagination
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "Get the export template list in pages"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /sysExportTemplate/getSysExportTemplateList [get]
export const getSysExportTemplateList = (params) => {
  return service({
    url: '/sysExportTemplate/getSysExportTemplateList',
    method: 'get',
    params
  })
}
