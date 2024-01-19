import service from '@/utils/request'
// @Router /authority/getAuthorityList [post]
export const getAuthorityList = (data) => {
  return service({
    url: '/authority/getAuthorityList',
    method: 'post',
    data
  })
}

// @Summary Delete role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {authorityId uint} true "Delete role"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /authority/deleteAuthority [post]
export const deleteAuthority = (data) => {
  return service({
    url: '/authority/deleteAuthority',
    method: 'post',
    data
  })
}

// @Summary Create role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "Create role"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /authority/createAuthority [post]
export const createAuthority = (data) => {
  return service({
    url: '/authority/createAuthority',
    method: 'post',
    data
  })
}

// @Tags authority
// @Summary copy character
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "Copy role"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Copy successful"}"
// @Router /authority/copyAuthority [post]
export const copyAuthority = (data) => {
  return service({
    url: '/authority/copyAuthority',
    method: 'post',
    data
  })
}

// @Summary Set role resource permissions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.SysAuthority true "Set role resource permissions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Set successfully"}"
// @Router /authority/setDataAuthority [post]
export const setDataAuthority = (data) => {
  return service({
    url: '/authority/setDataAuthority',
    method: 'post',
    data
  })
}

// @Summary Modify role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "Modify role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Set successfully"}"
// @Router /authority/setDataAuthority [post]
export const updateAuthority = (data) => {
  return service({
    url: '/authority/updateAuthority',
    method: 'put',
    data
  })
}
