import service from '@/utils/request'
// @Tags SysDictionary
// @Summary Create SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "Create SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /sysDictionary/createSysDictionary [post]
export const createSysDictionary = (data) => {
  return service({
    url: '/sysDictionary/createSysDictionary',
    method: 'post',
    data
  })
}

// @Tags SysDictionary
// @Summary delete SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "Delete SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Delete successfully"}"
// @Router /sysDictionary/deleteSysDictionary [delete]
export const deleteSysDictionary = (data) => {
  return service({
    url: '/sysDictionary/deleteSysDictionary',
    method: 'delete',
    data
  })
}

// @Tags SysDictionary
// @Summary Update SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "Update SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Update successful"}"
// @Router /sysDictionary/updateSysDictionary [put]
export const updateSysDictionary = (data) => {
  return service({
    url: '/sysDictionary/updateSysDictionary',
    method: 'put',
    data
  })
}

// @Tags SysDictionary
// @Summary Query SysDictionary with id
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "Query SysDictionary with id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Query successful"}"
// @Router /sysDictionary/findSysDictionary [get]
export const findSysDictionary = (params) => {
  return service({
    url: '/sysDictionary/findSysDictionary',
    method: 'get',
    params
  })
}

// @Tags SysDictionary
// @Summary Get the SysDictionary list by pagination
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "Get the SysDictionary list in pages"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /sysDictionary/getSysDictionaryList [get]
export const getSysDictionaryList = (params) => {
  return service({
    url: '/sysDictionary/getSysDictionaryList',
    method: 'get',
    params
  })
}
