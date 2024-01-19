package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeHistoryApi struct{}

// First
// @Tags      AutoCode
// @Summary Get meta information
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.GetById true "Request parameter"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Get meta information"
// @Router    /autoCode/getMeta [post]
func (a *AutoCodeHistoryApi) First(c *gin.Context) {
	var info request.GetById
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := autoCodeHistoryService.First(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"meta": data}, "Get successful", c)
}

// Delete
// @Tags      AutoCode
// @Summary Delete rollback records
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.GetById true "Request parameter"
// @Success 200 {object} response.Response{msg=string} "Delete rollback record"
// @Router    /autoCode/delSysHistory [post]
func (a *AutoCodeHistoryApi) Delete(c *gin.Context) {
	var info request.GetById
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = autoCodeHistoryService.Delete(&info)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// RollBack
// @Tags      AutoCode
// @Summary Roll back automatically generated code
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body systemReq.RollBack true "Request Parameters"
// @Success 200 {object} response.Response{msg=string} "Rollback automatically generated code"
// @Router    /autoCode/rollback [post]
func (a *AutoCodeHistoryApi) RollBack(c *gin.Context) {
	var info systemReq.RollBack
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = autoCodeHistoryService.RollBack(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Rollback successful", c)
}

// GetList
// @Tags      AutoCode
// @Summary Query rollback records
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body systemReq.SysAutoHistory true "Request Parameters"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Query the rollback record, return includes list, total number, page number, and number of each page"
// @Router    /autoCode/getSysHistory [post]
func (a *AutoCodeHistoryApi) GetList(c *gin.Context) {
	var search systemReq.SysAutoHistory
	err := c.ShouldBindJSON(&search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := autoCodeHistoryService.GetList(search.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     search.Page,
		PageSize: search.PageSize,
	}, "Get successful", c)
}
