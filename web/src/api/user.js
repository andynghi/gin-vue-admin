import service from '@/utils/request'
// @Summary User login
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
  return service({
    url: '/base/login',
    method: 'post',
    data: data
  })
}

// @Summary Get verification code
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const captcha = (data) => {
  return service({
    url: '/base/captcha',
    method: 'post',
    data: data
  })
}

// @Summary User registration
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/resige [post]
export const register = (data) => {
  return service({
    url: '/user/admin_register',
    method: 'post',
    data: data
  })
}

// @Summary Change password
// @Produce  application/json
// @Param data body {username:"string",password:"string",newPassword:"string"}
// @Router /user/changePassword [post]
export const changePassword = (data) => {
  return service({
    url: '/user/changePassword',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary Get the user list by pagination
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "Get the user list by page"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /user/getUserList [post]
export const getUserList = (data) => {
  return service({
    url: '/user/getUserList',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary Set user permissions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SetUserAuth true "Set user permissions"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Modification successful"}"
// @Router /user/setUserAuthority [post]
export const setUserAuthority = (data) => {
  return service({
    url: '/user/setUserAuthority',
    method: 'post',
    data: data
  })
}

// @Tags SysUser
// @Summary delete user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "Delete user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Modification successful"}"
// @Router /user/deleteUser [delete]
export const deleteUser = (data) => {
  return service({
    url: '/user/deleteUser',
    method: 'delete',
    data: data
  })
}

// @Tags SysUser
// @Summary Set user information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "Set user information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Modification successful"}"
// @Router /user/setUserInfo [put]
export const setUserInfo = (data) => {
  return service({
    url: '/user/setUserInfo',
    method: 'put',
    data: data
  })
}

// @Tags SysUser
// @Summary Set user information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "Set user information"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Modification successful"}"
// @Router /user/setSelfInfo [put]
export const setSelfInfo = (data) => {
  return service({
    url: '/user/setSelfInfo',
    method: 'put',
    data: data
  })
}

// @Tags User
// @Summary Set user permissions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.setUserAuthorities true "Set user permissions"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Modification successful"}"
// @Router /user/setUserAuthorities [post]
export const setUserAuthorities = (data) => {
  return service({
    url: '/user/setUserAuthorities',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary Get user information
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /user/getUserInfo [get]
export const getUserInfo = () => {
  return service({
    url: '/user/getUserInfo',
    method: 'get'
  })
}

export const resetPassword = (data) => {
  return service({
    url: '/user/resetPassword',
    method: 'post',
    data: data
  })
}
