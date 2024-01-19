import service from '@/utils/request'
// @Tags jwt
// @Summary jwt added to the blacklist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Blacklisted successfully"}"
// @Router /jwt/jsonInBlacklist [post]
export const jsonInBlacklist = () => {
  return service({
    url: '/jwt/jsonInBlacklist',
    method: 'post'
  })
}
