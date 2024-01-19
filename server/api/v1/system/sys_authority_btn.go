package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityBtnApi struct{}

// GetAuthorityBtn
// @Tags      AuthorityBtn
// @Summary Get permission button
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.SysAuthorityBtnReq true "Menu id, role id, selected button id"
// @Success 200 {object} response.Response{data=response.SysAuthorityBtnRes,msg=string} "Return list successfully"
// @Router    /authorityBtn/getAuthorityBtn [post]
func (a *AuthorityBtnApi) GetAuthorityBtn(c *gin.Context) {
	var req request.SysAuthorityBtnReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := authorityBtnService.GetAuthorityBtn(req)
	if err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Error(err))
		response.FailWithMessage("Query failed", c)
		return
	}
	response.OkWithDetailed(res, "Query successful", c)
}

// SetAuthorityBtn
// @Tags      AuthorityBtn
// @Summary Set permission button
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.SysAuthorityBtnReq true "Menu id, role id, selected button id"
// @Success 200 {object} response.Response{msg=string} "Return list successfully"
// @Router    /authorityBtn/setAuthorityBtn [post]
func (a *AuthorityBtnApi) SetAuthorityBtn(c *gin.Context) {
	var req request.SysAuthorityBtnReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authorityBtnService.SetAuthorityBtn(req)
	if err != nil {
		global.GVA_LOG.Error("Allocation failed!", zap.Error(err))
		response.FailWithMessage("Allocation failed", c)
		return
	}
	response.OkWithMessage("Assignment successful", c)
}

// CanRemoveAuthorityBtn
// @Tags      AuthorityBtn
// @Summary Set permission button
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{msg=string} "Delete successfully"
// @Router    /authorityBtn/canRemoveAuthorityBtn [post]
func (a *AuthorityBtnApi) CanRemoveAuthorityBtn(c *gin.Context) {
	id := c.Query("id")
	err := authorityBtnService.CanRemoveAuthorityBtn(id)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}
