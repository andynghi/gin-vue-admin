package middleware

import (
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler interceptor
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.GVA_CONFIG.System.Env != "develop" {
			waitUse, _ := utils.GetClaims(c)
			//Get the requested PATH
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, global.GVA_CONFIG.System.RouterPrefix)
			// Get request method
			act := c.Request.Method
			// Get the user's role
			sub := strconv.Itoa(int(waitUse.AuthorityId))
			e := casbinService.Casbin() // Determine whether the policy exists
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDetailed(gin.H{}, "Insufficient permissions", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
