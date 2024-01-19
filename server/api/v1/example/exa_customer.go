package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/example/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerApi struct{}

// CreateExaCustomer
// @Tags      ExaCustomer
// @Summary Create a customer
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body example.ExaCustomer true "Customer username, customer mobile phone number"
// @Success 200 {object} response.Response{msg=string} "Create customer"
// @Router    /customer/customer [post]
func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer, utils.CustomerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customer.SysUserID = utils.GetUserID(c)
	customer.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = customerService.CreateExaCustomer(customer)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Created successfully", c)
}

// DeleteExaCustomer
// @Tags      ExaCustomer
// @Summary Delete customer
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body example.ExaCustomer true "Customer ID"
// @Success 200 {object} response.Response{msg=string} "Delete customer"
// @Router    /customer/customer [delete]
func (e *CustomerApi) DeleteExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = customerService.DeleteExaCustomer(customer)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// UpdateExaCustomer
// @Tags      ExaCustomer
// @Summary Update customer information
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body example.ExaCustomer true "Customer ID, customer information"
// @Success 200 {object} response.Response{msg=string} "Update customer information"
// @Router    /customer/customer [put]
func (e *CustomerApi) UpdateExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer, utils.CustomerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = customerService.UpdateExaCustomer(&customer)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// GetExaCustomer
// @Tags      ExaCustomer
// @Summary Get single customer information
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query example.ExaCustomer true "Customer ID"
// @Success 200 {object} response.Response{data=exampleRes.ExaCustomerResponse,msg=string} "Get single customer information, return including customer details"
// @Router    /customer/customer [get]
func (e *CustomerApi) GetExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindQuery(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(customer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := customerService.GetExaCustomer(customer.ID)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaCustomerResponse{Customer: data}, "Get successful", c)
}

// GetExaCustomerList
// @Tags      ExaCustomer
// @Summary Get the list of authorized customers in paging
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query request.PageInfo true "Page number, size of each page"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Get the list of authorized customers in paging, and return the list, total number, page number, and number per page"
// @Router    /customer/customerList [get]
func (e *CustomerApi) GetExaCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customerList, total, err := customerService.GetCustomerInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     customerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get successful", c)
}
