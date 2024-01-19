import service from '@/utils/request'
// @Tags SysDictionaryDetail
// @Summary Create SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "Create SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
export const createSysDictionaryDetail = (data) => {
  return service({
    url: '/sysDictionaryDetail/createSysDictionaryDetail',
    method: 'post',
    data
  })
}

// @Tags SysDictionaryDetail
// @Summary Delete SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "Delete SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Delete successfully"}"
// @Router /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
export const deleteSysDictionaryDetail = (data) => {
  return service({
    url: '/sysDictionaryDetail/deleteSysDictionaryDetail',
    method: 'delete',
    data
  })
}

// @Tags SysDictionaryDetail
// @Summary Update SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "Update SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Update successful"}"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
export const updateSysDictionaryDetail = (data) => {
  return service({
    url: '/sysDictionaryDetail/updateSysDictionaryDetail',
    method: 'put',
    data
  })
}

// @Tags SysDictionaryDetail
// @Summary Use id to query SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "Query SysDictionaryDetail with id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Query successful"}"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
export const findSysDictionaryDetail = (params) => {
  return service({
    url: '/sysDictionaryDetail/findSysDictionaryDetail',
    method: 'get',
    params
  })
}

// @Tags SysDictionaryDetail
// @Summary Get the SysDictionaryDetail list by pagination
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "Get the SysDictionaryDetail list in pages"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
export const getSysDictionaryDetailList = (params) => {
  return service({
    url: '/sysDictionaryDetail/getSysDictionaryDetailList',
    method: 'get',
    params
  })
}
