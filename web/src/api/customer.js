import service from '@/utils/request'
// @Tags SysApi
// @Summary Delete customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "Delete customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /customer/customer [post]
export const createExaCustomer = (data) => {
  return service({
    url: '/customer/customer',
    method: 'post',
    data
  })
}

// @Tags SysApi
// @Summary Update customer information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "Update customer information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /customer/customer [put]
export const updateExaCustomer = (data) => {
  return service({
    url: '/customer/customer',
    method: 'put',
    data
  })
}

// @Tags SysApi
// @Summary Create a customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "Create customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /customer/customer [delete]
export const deleteExaCustomer = (data) => {
  return service({
    url: '/customer/customer',
    method: 'delete',
    data
  })
}

// @Tags SysApi
// @Summary Get single customer information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "Get single customer information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /customer/customer [get]
export const getExaCustomer = (params) => {
  return service({
    url: '/customer/customer',
    method: 'get',
    params
  })
}

// @Tags SysApi
// @Summary Get the list of authorized customers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "Get permission customer list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /customer/customerList [get]
export const getExaCustomerList = (params) => {
  return service({
    url: '/customer/customerList',
    method: 'get',
    params
  })
}
