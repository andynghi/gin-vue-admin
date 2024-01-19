package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register)               // Administrator registration account
		userRouter.POST("changePassword", baseApi.ChangePassword)         // User changes password
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     //Set user permissions
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // Delete user
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // Set user information
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                //Set self-information
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) //Set user permission group
		userRouter.POST("resetPassword", baseApi.ResetPassword)           //Set user permission group
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // Get the user list in paging
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // Get own information
	}
}
