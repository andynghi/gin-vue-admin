import service from '@/utils/request'

// @Tags api
// @Summary Get the role list by pagination
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "Get the user list by page"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
export const getApiList = (data) => {
  return service({
    url: '/api/getApiList',
    method: 'post',
    data
  })
}

// @Tags Api
// @Summary Create basic api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "Create api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /api/createApi [post]
export const createApi = (data) => {
  return service({
    url: '/api/createApi',
    method: 'post',
    data
  })
}

// @Tags menu
// @Summary Get the menu based on id
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetById true "Get menu based on id"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /menu/getApiById [post]
export const getApiById = (data) => {
  return service({
    url: '/api/getApiById',
    method: 'post',
    data
  })
}

// @Tags Api
// @Summary update api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "Update api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Update successful"}"
// @Router /api/updateApi [post]
export const updateApi = (data) => {
  return service({
    url: '/api/updateApi',
    method: 'post',
    data
  })
}

// @Tags Api
// @Summary update api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "Update api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Update successful"}"
// @Router /api/setAuthApi [post]
export const setAuthApi = (data) => {
  return service({
    url: '/api/setAuthApi',
    method: 'post',
    data
  })
}

// @Tags Api
// @Summary Get all APIs without paging
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /api/getAllApis [post]
export const getAllApis = (data) => {
  return service({
    url: '/api/getAllApis',
    method: 'post',
    data
  })
}

// @Tags Api
// @Summary Delete the specified api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Api true "Delete api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Delete successfully"}"
// @Router /api/deleteApi [post]
export const deleteApi = (data) => {
  return service({
    url: '/api/deleteApi',
    method: 'post',
    data
  })
}

// @Tags SysApi
// @Summary Delete selected API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Delete successfully"}"
// @Router /api/deleteApisByIds [delete]
export const deleteApisByIds = (data) => {
  return service({
    url: '/api/deleteApisByIds',
    method: 'delete',
    data
  })
}

// FreshCasbin
// @Tags      SysApi
// @Summary refresh casbin cache
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{msg=string} "Refresh successful"
// @Router    /api/freshCasbin [get]
export const freshCasbin = () => {
  return service({
    url: '/api/freshCasbin',
    method: 'get'
  })
}
