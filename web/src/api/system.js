import service from '@/utils/request'
// @Tags systrm
// @Summary Get the configuration file content
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Return successful"}"
// @Router /system/getSystemConfig [post]
export const getSystemConfig = () => {
  return service({
    url: '/system/getSystemConfig',
    method: 'post'
  })
}

// @Tags system
// @Summary Set configuration file content
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body sysModel.System true
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Return successful"}"
// @Router /system/setSystemConfig [post]
export const setSystemConfig = (data) => {
  return service({
    url: '/system/setSystemConfig',
    method: 'post',
    data
  })
}

// @Tags system
// @Summary Get server running status
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Return successful"}"
// @Router /system/getServerInfo [post]
export const getSystemState = () => {
  return service({
    url: '/system/getServerInfo',
    method: 'post',
    donNotShowLoading: true
  })
}
