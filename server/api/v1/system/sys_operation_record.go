package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OperationRecordApi struct{}

// CreateSysOperationRecord
// @Tags      SysOperationRecord
// @Summary Create SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysOperationRecord true "Create SysOperationRecord"
// @Success 200 {object} response.Response{msg=string} "Create SysOperationRecord"
// @Router    /sysOperationRecord/createSysOperationRecord [post]
func (s *OperationRecordApi) CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	err := c.ShouldBindJSON(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.CreateSysOperationRecord(sysOperationRecord)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Created successfully", c)
}

// DeleteSysOperationRecord
// @Tags      SysOperationRecord
// @Summary Delete SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysOperationRecord true "SysOperationRecord model"
// @Success 200 {object} response.Response{msg=string} "Delete SysOperationRecord"
// @Router    /sysOperationRecord/deleteSysOperationRecord [delete]
func (s *OperationRecordApi) DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	err := c.ShouldBindJSON(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.DeleteSysOperationRecord(sysOperationRecord)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// DeleteSysOperationRecordByIds
// @Tags      SysOperationRecord
// @Summary Delete SysOperationRecord in batches
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.IdsReq true "Delete SysOperationRecord in batches"
// @Success 200 {object} response.Response{msg=string} "Delete SysOperationRecord in batches"
// @Router    /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (s *OperationRecordApi) DeleteSysOperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = operationRecordService.DeleteSysOperationRecordByIds(IDS)
	if err != nil {
		global.GVA_LOG.Error("Batch deletion failed!", zap.Error(err))
		response.FailWithMessage("Batch deletion failed", c)
		return
	}
	response.OkWithMessage("Batch deletion successful", c)
}

// FindSysOperationRecord
// @Tags      SysOperationRecord
// @Summary Query SysOperationRecord with id
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysOperationRecord                                  true  "Id"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Query SysOperationRecord with id"
// @Router    /sysOperationRecord/findSysOperationRecord [get]
func (s *OperationRecordApi) FindSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	err := c.ShouldBindQuery(&sysOperationRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(sysOperationRecord, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reSysOperationRecord, err := operationRecordService.GetSysOperationRecord(sysOperationRecord.ID)
	if err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithMessage("Query failed", c)
		return
	}
	response.OkWithDetailed(gin.H{"reSysOperationRecord": reSysOperationRecord}, "Query successful", c)
}

// GetSysOperationRecordList
// @Tags      SysOperationRecord
// @Summary Get the SysOperationRecord list in paging
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query request.SysOperationRecordSearch true "Page number, size of each page, search conditions"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Get the SysOperationRecord list in paging, and return the list, total number, page number, and number of each page"
// @Router    /sysOperationRecord/getSysOperationRecordList [get]
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo systemReq.SysOperationRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := operationRecordService.GetSysOperationRecordInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get successful", c)
}
