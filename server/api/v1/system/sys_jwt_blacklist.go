package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtApi struct{}

// JsonInBlacklist
// @Tags      Jwt
// @Summary jwt added to the blacklist
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{msg=string} "jwt added to blacklist"
// @Router    /jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := utils.GetToken(c)
	jwt := system.JwtBlacklist{Jwt: token}
	err := jwtService.JsonInBlacklist(jwt)
	if err != nil {
		global.GVA_LOG.Error("jwt invalidation failed!", zap.Error(err))
		response.FailWithMessage("jwt invalidation failed", c)
		return
	}
	utils.ClearToken(c)
	response.OkWithMessage("jwt invalidated successfully", c)
}
