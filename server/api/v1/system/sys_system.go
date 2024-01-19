package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApi struct{}

// GetSystemConfig
// @Tags      System
// @Summary Get the configuration file content
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success 200 {object} response.Response{data=systemRes.SysConfigResponse,msg=string} "Get the configuration file content and return including system configuration"
// @Router    /system/getSystemConfig [post]
func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	config, err := systemConfigService.GetSystemConfig()
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "Get successful", c)
}

// SetSystemConfig
// @Tags      System
// @Summary Set configuration file content
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param data body system.System true "Set configuration file content"
// @Success 200 {object} response.Response{data=string} "Set configuration file content"
// @Router    /system/setSystemConfig [post]
func (s *SystemApi) SetSystemConfig(c *gin.Context) {
	var sys system.System
	err := c.ShouldBindJSON(&sys)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = systemConfigService.SetSystemConfig(sys)
	if err != nil {
		global.GVA_LOG.Error("Setting failed!", zap.Error(err))
		response.FailWithMessage("Setting failed", c)
		return
	}
	response.OkWithMessage("Set successfully", c)
}

// ReloadSystem
// @Tags      System
// @Summary Restart the system
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success 200 {object} response.Response{msg=string} "Restart the system"
// @Router    /system/reloadSystem [post]
func (s *SystemApi) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		global.GVA_LOG.Error("Restarting the system failed!", zap.Error(err))
		response.FailWithMessage("Failed to restart the system", c)
		return
	}
	response.OkWithMessage("Reboot the system successfully", c)
}

// GetServerInfo
// @Tags      System
// @Summary Get server information
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Get server information"
// @Router    /system/getServerInfo [post]
func (s *SystemApi) GetServerInfo(c *gin.Context) {
	server, err := systemConfigService.GetServerInfo()
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(gin.H{"server": server}, "Get successful", c)
}
