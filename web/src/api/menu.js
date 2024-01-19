import service from '@/utils/request'
// @Summary User login to obtain dynamic routing
// @Produce  application/json
// @Param You can fill in nothing, just adjust it
// @Router /menu/getMenu [post]
export const asyncMenu = () => {
  return service({
    url: '/menu/getMenu',
    method: 'post'
  })
}

// @Summary Get the menu list
// @Produce  application/json
// @Param {
//  page     int
//	pageSize int
// }
// @Router /menu/getMenuList [post]
export const getMenuList = (data) => {
  return service({
    url: '/menu/getMenuList',
    method: 'post',
    data
  })
}

// @Summary Add a new basic menu
// @Produce  application/json
// @Param menu Object
// @Router /menu/getMenuList [post]
export const addBaseMenu = (data) => {
  return service({
    url: '/menu/addBaseMenu',
    method: 'post',
    data
  })
}

// @Summary Get the basic route list
// @Produce  application/json
// @Param You can fill in nothing, just adjust it
// @Router /menu/getBaseMenuTree [post]
export const getBaseMenuTree = () => {
  return service({
    url: '/menu/getBaseMenuTree',
    method: 'post'
  })
}

// @Summary Add user menu association relationship
// @Produce  application/json
// @Param menus Object authorityId string
// @Router /menu/getMenuList [post]
export const addMenuAuthority = (data) => {
  return service({
    url: '/menu/addMenuAuthority',
    method: 'post',
    data
  })
}

// @Summary Get the user menu association relationship
// @Produce  application/json
// @Param authorityId string
// @Router /menu/getMenuAuthority [post]
export const getMenuAuthority = (data) => {
  return service({
    url: '/menu/getMenuAuthority',
    method: 'post',
    data
  })
}

// @Summary delete menu
// @Produce  application/json
// @Param ID float64
// @Router /menu/deleteBaseMenu [post]
export const deleteBaseMenu = (data) => {
  return service({
    url: '/menu/deleteBaseMenu',
    method: 'post',
    data
  })
}

// @Summary Modify menu list
// @Produce  application/json
// @Param menu Object
// @Router /menu/updateBaseMenu [post]
export const updateBaseMenu = (data) => {
  return service({
    url: '/menu/updateBaseMenu',
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
// @Router /menu/getBaseMenuById [post]
export const getBaseMenuById = (data) => {
  return service({
    url: '/menu/getBaseMenuById',
    method: 'post',
    data
  })
}
