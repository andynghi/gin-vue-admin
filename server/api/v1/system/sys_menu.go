package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct{}

// GetMenu
// @Tags      AuthorityMenu
// @Summary Get user dynamic routing
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param data body request.Empty true "empty"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "Get user dynamic routing and return a list including system menu details"
// @Router    /menu/getMenu [post]
func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c))
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	if menus == nil {
		menus = []system.SysMenu{}
	}
	response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "Get successful", c)
}

// GetBaseMenuTree
// @Tags      AuthorityMenu
// @Summary Get user dynamic routing
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param data body request.Empty true "empty"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenusResponse,msg=string} "Get user dynamic routing, return including system menu list"
// @Router    /menu/getBaseMenuTree [post]
func (a *AuthorityMenuApi) GetBaseMenuTree(c *gin.Context) {
	menus, err := menuService.GetBaseMenuTree()
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(systemRes.SysBaseMenusResponse{Menus: menus}, "Get successful", c)
}

// AddMenuAuthority
// @Tags      AuthorityMenu
// @Summary Add the relationship between menu and role
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body systemReq.AddMenuAuthorityInfo true "role ID"
// @Success 200 {object} response.Response{msg=string} "Add the relationship between menu and role"
// @Router    /menu/addMenuAuthority [post]
func (a *AuthorityMenuApi) AddMenuAuthority(c *gin.Context) {
	var authorityMenu systemReq.AddMenuAuthorityInfo
	err := c.ShouldBindJSON(&authorityMenu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.GVA_LOG.Error("Add failed!", zap.Error(err))
		response.FailWithMessage("Add failed", c)
	} else {
		response.OkWithMessage("Added successfully", c)
	}
}

// GetMenuAuthority
// @Tags      AuthorityMenu
// @Summary Get the specified role menu
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.GetAuthorityId true "Role ID"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Get the specified role menu"
// @Router    /menu/getMenuAuthority [post]
func (a *AuthorityMenuApi) GetMenuAuthority(c *gin.Context) {
	var param request.GetAuthorityId
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(param, utils.AuthorityIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menus, err := menuService.GetMenuAuthority(&param)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "Failed to obtain", c)
		return
	}
	response.OkWithDetailed(gin.H{"menus": menus}, "Get successful", c)
}

// AddBaseMenu
// @Tags      Menu
// @Summary New menu
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysBaseMenu true "Routing path, parent menu ID, routing name, corresponding front-end file path, sorting mark"
// @Success 200 {object} response.Response{msg=string} "New menu"
// @Router    /menu/addBaseMenu [post]
func (a *AuthorityMenuApi) AddBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu.Meta, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = menuService.AddBaseMenu(menu)
	if err != nil {
		global.GVA_LOG.Error("Add failed!", zap.Error(err))
		response.FailWithMessage("Add failed", c)
		return
	}
	response.OkWithMessage("Added successfully", c)
}

// DeleteBaseMenu
// @Tags      Menu
// @Summary delete menu
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.GetById true "menu id"
// @Success 200 {object} response.Response{msg=string} "Delete menu"
// @Router    /menu/deleteBaseMenu [post]
func (a *AuthorityMenuApi) DeleteBaseMenu(c *gin.Context) {
	var menu request.GetById
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = baseMenuService.DeleteBaseMenu(menu.ID)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// UpdateBaseMenu
// @Tags      Menu
// @Summary update menu
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysBaseMenu true "Routing path, parent menu ID, routing name, corresponding front-end file path, sorting mark"
// @Success 200 {object} response.Response{msg=string} "Update menu"
// @Router    /menu/updateBaseMenu [post]
func (a *AuthorityMenuApi) UpdateBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu.Meta, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = baseMenuService.UpdateBaseMenu(menu)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// GetBaseMenuById
// @Tags      Menu
// @Summary Get the menu based on id
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.GetById true "menu id"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenuResponse,msg=string} "Get the menu based on the id and return the system menu list"
// @Router    /menu/getBaseMenuById [post]
func (a *AuthorityMenuApi) GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menu, err := baseMenuService.GetBaseMenuById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "Get successful", c)
}

// GetMenuList
// @Tags      Menu
// @Summary Get the basic menu list by pagination
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.PageInfo true "Page number, size of each page"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Get the basic menu list by paging, and return the list, total number, page number, and number per page"
// @Router    /menu/getMenuList [post]
func (a *AuthorityMenuApi) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menuList, total, err := menuService.GetInfoList()
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     menuList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get successful", c)
}
