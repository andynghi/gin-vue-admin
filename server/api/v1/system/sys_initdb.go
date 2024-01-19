package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type DBApi struct{}

// InitDB
// @Tags     InitDB
// @Summary Initialize user database
// @Produce  application/json
// @Param data body request.InitDB true "Initialization database parameters"
// @Success 200 {object} response.Response{data=string} "Initialize user database"
// @Router   /init/initdb [post]
func (i *DBApi) InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error("Database configuration already exists!")
		response.FailWithMessage("Database configuration already exists", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("Parameter verification failed!", zap.Error(err))
		response.FailWithMessage("Parameter verification failed", c)
		return
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.GVA_LOG.Error("Automatic database creation failed!", zap.Error(err))
		response.FailWithMessage("Automatic database creation failed, please check the background log and initialize after checking", c)
		return
	}
	response.OkWithMessage("Automatically created database successfully", c)
}

// CheckDB
// @Tags     CheckDB
// @Summary Initialize user database
// @Produce  application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Initialize user database"
// @Router   /init/checkdb [post]
func (i *DBApi) CheckDB(c *gin.Context) {
	var (
		message  = "Go to initialize database"
		needInit = true
	)

	if global.GVA_DB != nil {
		message = "Database does not need to be initialized"
		needInit = false
	}
	global.GVA_LOG.Info(message)
	response.OkWithDetailed(gin.H{"needInit": needInit}, message, c)
}
