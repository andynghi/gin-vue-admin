package middleware

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"

	"github.com/gin-gonic/gin"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Here we use jwt authentication to obtain header information. x-token returns token information when logging in. The front-end needs to store the token in a cookie or local storage. However, it needs to negotiate the expiration time with the back-end. You can agree to refresh the token or log in again.
		token := utils.GetToken(c)
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "Not logged in or illegal access", c)
			c.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "Your account is logged in from another place or the token is invalid", c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken parses the information contained in the token
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "Authorization has expired", c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}

		// The logged-in user has been disabled by the administrator. It is necessary to invalidate the user's jwt. This is more performance consuming. If necessary, please open it yourself.
		// The logic of user deletion needs to be optimized. It consumes more performance here. If necessary, please open it yourself.

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		c.Set("claims", claims)
		c.Next()
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))
			if global.GVA_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // The blacklisting operation will only be performed when the previous acquisition is successful.
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				//Record the current active status anyway
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
	}
}
