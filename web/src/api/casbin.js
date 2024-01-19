import service from '@/utils/request'
// @Tags authority
// @Summary Change role api permissions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "Change role api permissions"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /casbin/UpdateCasbin [post]
export const UpdateCasbin = (data) => {
  return service({
    url: '/casbin/updateCasbin',
    method: 'post',
    data
  })
}

// @Tags casbin
// @Summary Get the permission list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "Get permission list"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
export const getPolicyPathByAuthorityId = (data) => {
  return service({
    url: '/casbin/getPolicyPathByAuthorityId',
    method: 'post',
    data
  })
}
