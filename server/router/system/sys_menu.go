package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("addBaseMenu", authorityMenuApi.AddBaseMenu)           // Add a new menu
		menuRouter.POST("addMenuAuthority", authorityMenuApi.AddMenuAuthority) // Add the relationship between menu and role
		menuRouter.POST("deleteBaseMenu", authorityMenuApi.DeleteBaseMenu)     // Delete menu
		menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)     //Update menu
	}
	{
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu)                   // Get the menu tree
		menuRouterWithoutRecord.POST("getMenuList", authorityMenuApi.GetMenuList)           // Get the basic menu list in paging
		menuRouterWithoutRecord.POST("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)   // Get user dynamic routing
		menuRouterWithoutRecord.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // Get the specified role menu
		menuRouterWithoutRecord.POST("getBaseMenuById", authorityMenuApi.GetBaseMenuById)   // Get menu based on id
	}
	return menuRouter
}
