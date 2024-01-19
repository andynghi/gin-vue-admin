package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApiApi struct{}

// CreateApi
// @Tags      SysApi
// @Summary Create basic api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysApi true "api path, api Chinese description, api group, method"
// @Success 200 {object} response.Response{msg=string} "Create basic api"
// @Router    /api/createApi [post]
func (s *SystemApiApi) CreateApi(c *gin.Context) {
	var api system.SysApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiService.CreateApi(api)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Created successfully", c)
}

// DeleteApi
// @Tags      SysApi
// @Summary delete api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysApi                  true  "ID"
// @Success 200 {object} response.Response{msg=string} "Delete api"
// @Router    /api/deleteApi [post]
func (s *SystemApiApi) DeleteApi(c *gin.Context) {
	var api system.SysApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiService.DeleteApi(api)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// GetApiList
// @Tags      SysApi
// @Summary Get the API list by pagination
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body systemReq.SearchApiParams true "Get the API list by pagination"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Get the API list in paging, and return the list, total number, page number, and number per page"
// @Router    /api/getApiList [post]
func (s *SystemApiApi) GetApiList(c *gin.Context) {
	var pageInfo systemReq.SearchApiParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := apiService.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
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

// GetApiById
// @Tags      SysApi
// @Summary Get api based on id
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.GetById true "Get api based on id"
// @Success 200 {object} response.Response{data=systemRes.SysAPIResponse} "Get api based on id, return including api details"
// @Router    /api/getApiById [post]
func (s *SystemApiApi) GetApiById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	api, err := apiService.GetApiById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(systemRes.SysAPIResponse{Api: api}, "Get successful", c)
}

// UpdateApi
// @Tags      SysApi
// @Summary Modify the basic api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysApi true "api path, api Chinese description, api group, method"
// @Success 200 {object} response.Response{msg=string} "Modify basic api"
// @Router    /api/updateApi [post]
func (s *SystemApiApi) UpdateApi(c *gin.Context) {
	var api system.SysApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiService.UpdateApi(api)
	if err != nil {
		global.GVA_LOG.Error("Modification failed!", zap.Error(err))
		response.FailWithMessage("Modification failed", c)
		return
	}
	response.OkWithMessage("Modification successful", c)
}

// GetAllApis
// @Tags      SysApi
// @Summary Get all APIs without paging
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{data=systemRes.SysAPIListResponse,msg=string} "Get all APIs without paging, return including api list"
// @Router    /api/getAllApis [post]
func (s *SystemApiApi) GetAllApis(c *gin.Context) {
	apis, err := apiService.GetAllApis()
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(systemRes.SysAPIListResponse{Apis: apis}, "Get successful", c)
}

// DeleteApisByIds
// @Tags      SysApi
// @Summary Delete selected API
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "ID"
// @Success 200 {object} response.Response{msg=string} "Delete selected Api"
// @Router    /api/deleteApisByIds [delete]
func (s *SystemApiApi) DeleteApisByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiService.DeleteApisByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// FreshCasbin
// @Tags      SysApi
// @Summary refresh casbin cache
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{msg=string} "Refresh successful"
// @Router    /api/freshCasbin [get]
func (s *SystemApiApi) FreshCasbin(c *gin.Context) {
	err := casbinService.FreshCasbin()
	if err != nil {
		global.GVA_LOG.Error("Refresh failed!", zap.Error(err))
		response.FailWithMessage("Refresh failed", c)
		return
	}
	response.OkWithMessage("Refresh successful", c)
}
