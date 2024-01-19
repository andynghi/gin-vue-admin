import service from '@/utils/request'
// @Tags SysOperationRecord
// @Summary Delete SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "Delete SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Delete successfully"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
export const deleteSysOperationRecord = (data) => {
  return service({
    url: '/sysOperationRecord/deleteSysOperationRecord',
    method: 'delete',
    data
  })
}

// @Tags SysOperationRecord
// @Summary Delete SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Delete successfully"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
export const deleteSysOperationRecordByIds = (data) => {
  return service({
    url: '/sysOperationRecord/deleteSysOperationRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags SysOperationRecord
// @Summary Get the SysOperationRecord list in paging
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "Get the SysOperationRecord list in pages"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
export const getSysOperationRecordList = (params) => {
  return service({
    url: '/sysOperationRecord/getSysOperationRecordList',
    method: 'get',
    params
  })
}
