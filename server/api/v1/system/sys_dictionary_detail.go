package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictionaryDetailApi struct{}

// CreateSysDictionaryDetail
// @Tags      SysDictionaryDetail
// @Summary Create SysDictionaryDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysDictionaryDetail true "SysDictionaryDetail model"
// @Success 200 {object} response.Response{msg=string} "Create SysDictionaryDetail"
// @Router    /sysDictionaryDetail/createSysDictionaryDetail [post]
func (s *DictionaryDetailApi) CreateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.CreateSysDictionaryDetail(detail)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Created successfully", c)
}

// DeleteSysDictionaryDetail
// @Tags      SysDictionaryDetail
// @Summary Delete SysDictionaryDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysDictionaryDetail true "SysDictionaryDetail model"
// @Success 200 {object} response.Response{msg=string} "Delete SysDictionaryDetail"
// @Router    /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func (s *DictionaryDetailApi) DeleteSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.DeleteSysDictionaryDetail(detail)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// UpdateSysDictionaryDetail
// @Tags      SysDictionaryDetail
// @Summary Update SysDictionaryDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysDictionaryDetail true "Update SysDictionaryDetail"
// @Success 200 {object} response.Response{msg=string} "Update SysDictionaryDetail"
// @Router    /sysDictionaryDetail/updateSysDictionaryDetail [put]
func (s *DictionaryDetailApi) UpdateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.UpdateSysDictionaryDetail(&detail)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// FindSysDictionaryDetail
// @Tags      SysDictionaryDetail
// @Summary Use id to query SysDictionaryDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query system.SysDictionaryDetail true "Query SysDictionaryDetail with id"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Query SysDictionaryDetail with id"
// @Router    /sysDictionaryDetail/findSysDictionaryDetail [get]
func (s *DictionaryDetailApi) FindSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	err := c.ShouldBindQuery(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(detail, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reSysDictionaryDetail, err := dictionaryDetailService.GetSysDictionaryDetail(detail.ID)
	if err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithMessage("Query failed", c)
		return
	}
	response.OkWithDetailed(gin.H{"reSysDictionaryDetail": reSysDictionaryDetail}, "Query successful", c)
}

// GetSysDictionaryDetailList
// @Tags      SysDictionaryDetail
// @Summary Get the SysDictionaryDetail list by pagination
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query request.SysDictionaryDetailSearch true "Page number, size of each page, search conditions"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Get the SysDictionaryDetail list by paging, and return the list, total number, page number, and number of each page"
// @Router    /sysDictionaryDetail/getSysDictionaryDetailList [get]
func (s *DictionaryDetailApi) GetSysDictionaryDetailList(c *gin.Context) {
	var pageInfo request.SysDictionaryDetailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dictionaryDetailService.GetSysDictionaryDetailInfoList(pageInfo)
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
