package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityApi struct{}

// CreateAuthority
// @Tags      Authority
// @Summary Create role
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysAuthority true "Permission id, permission name, parent role id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "Create role, return including system role details"
// @Router    /authority/createAuthority [post]
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority, authBack system.SysAuthority
	var err error

	if err = c.ShouldBindJSON(&authority); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err = utils.Verify(authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if authBack, err = authorityService.CreateAuthority(authority); err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed"+err.Error(), c)
		return
	}
	err = casbinService.FreshCasbin()
	if err != nil {
		global.GVA_LOG.Error("Creation successful, permission refresh failed.", zap.Error(err))
		response.FailWithMessage("Creation successful, permission refresh failed."+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "Created successfully", c)
}

// CopyAuthority
// @Tags      Authority
// @Summary copy character
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body response.SysAuthorityCopyResponse true "Old role id, new permission id, new permission name, new parent role id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "Copy role, return including system role details"
// @Router    /authority/copyAuthority [post]
func (a *AuthorityApi) CopyAuthority(c *gin.Context) {
	var copyInfo systemRes.SysAuthorityCopyResponse
	err := c.ShouldBindJSON(&copyInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(copyInfo, utils.OldAuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(copyInfo.Authority, utils.AuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authBack, err := authorityService.CopyAuthority(copyInfo)
	if err != nil {
		global.GVA_LOG.Error("Copy failed!", zap.Error(err))
		response.FailWithMessage("Copy failed"+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "Copy successful", c)
}

// DeleteAuthority
// @Tags      Authority
// @Summary Delete role
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysAuthority true "Delete role"
// @Success 200 {object} response.Response{msg=string} "Delete role"
// @Router    /authority/deleteAuthority [post]
func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var authority system.SysAuthority
	var err error
	if err = c.ShouldBindJSON(&authority); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// Before deleting a role, you need to determine whether there is a user using this role.
	if err = authorityService.DeleteAuthority(&authority); err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed"+err.Error(), c)
		return
	}
	_ = casbinService.FreshCasbin()
	response.OkWithMessage("Deletion successful", c)
}

// UpdateAuthority
// @Tags      Authority
// @Summary Update character information
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysAuthority true "Permission id, permission name, parent role id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "Update role information, return including system role details"
// @Router    /authority/updateAuthority [post]
func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var auth system.SysAuthority
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(auth, utils.AuthorityVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authority, err := authorityService.UpdateAuthority(auth)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed"+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "Update successful", c)
}

// GetAuthorityList
// @Tags      Authority
// @Summary Get the role list by pagination
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.PageInfo true "Page number, size of each page"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Get the role list in paging, and return the list, total number, page number, and number per page"
// @Router    /authority/getAuthorityList [post]
func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := authorityService.GetAuthorityInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get successful", c)
}

// SetDataAuthority
// @Tags      Authority
// @Summary Set role resource permissions
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysAuthority true "Set role resource permissions"
// @Success 200 {object} response.Response{msg=string} "Set role resource permissions"
// @Router    /authority/setDataAuthority [post]
func (a *AuthorityApi) SetDataAuthority(c *gin.Context) {
	var auth system.SysAuthority
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(auth, utils.AuthorityIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authorityService.SetDataAuthority(auth)
	if err != nil {
		global.GVA_LOG.Error("Setting failed!", zap.Error(err))
		response.FailWithMessage("Setting failed"+err.Error(), c)
		return
	}
	response.OkWithMessage("Set successfully", c)
}
