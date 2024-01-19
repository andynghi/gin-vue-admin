import service from '@/utils/request'
// @Tags InitDB
// @Summary Initialize user database
// @Produce  application/json
// @Param data body request.InitDB true "Initialization database parameters"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"Automatically created database successfully"}"
// @Router /init/initdb [post]
export const initDB = (data) => {
  return service({
    url: '/init/initdb',
    method: 'post',
    data
  })
}

// @Tags CheckDB
// @Summary Initialize user database
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"Detection completed"}"
// @Router /init/checkdb [post]
export const checkDB = () => {
  return service({
    url: '/init/checkdb',
    method: 'post'
  })
}
