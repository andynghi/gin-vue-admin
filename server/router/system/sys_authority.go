package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority)   // Create a role
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority)   // Delete role
		authorityRouter.PUT("updateAuthority", authorityApi.UpdateAuthority)    // Update role
		authorityRouter.POST("copyAuthority", authorityApi.CopyAuthority)       // Copy role
		authorityRouter.POST("setDataAuthority", authorityApi.SetDataAuthority) //Set role resource permissions
	}
	{
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList) // Get the role list
	}
}
