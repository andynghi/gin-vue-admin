package system

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeApi struct{}

// PreviewTemp
// @Tags      AutoCode
// @Summary Preview the created code
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.AutoCodeStruct true "Preview creation code"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Preview the created code"
// @Router    /autoCode/preview [post]
func (autoApi *AutoCodeApi) PreviewTemp(c *gin.Context) {
	var a system.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Pretreatment() // Process the go keyword
	a.PackageT = utils.FirstUpper(a.Package)
	autoCode, err := autoCodeService.PreviewTemp(a)
	if err != nil {
		global.GVA_LOG.Error("Preview failed!", zap.Error(err))
		response.FailWithMessage("Preview failed", c)
	} else {
		response.OkWithDetailed(gin.H{"autoCode": autoCode}, "Preview successful", c)
	}
}

// CreateTemp
// @Tags      AutoCode
// @Summary automatic code template
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.AutoCodeStruct true "Create automatic code"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Created successfully"}"
// @Router    /autoCode/createTemp [post]
func (autoApi *AutoCodeApi) CreateTemp(c *gin.Context) {
	var a system.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Pretreatment()
	var apiIds []uint
	if a.AutoCreateApiToSql {
		if ids, err := autoCodeService.AutoCreateApi(&a); err != nil {
			global.GVA_LOG.Error("Automated creation failed! Please clear the garbage data yourself!", zap.Error(err))
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape("Automation creation failed! Please clear the junk data yourself!"))
			return
		} else {
			apiIds = ids
		}
	}
	a.PackageT = utils.FirstUpper(a.Package)
	err := autoCodeService.CreateTemp(a, apiIds...)
	if err != nil {
		if errors.Is(err, system.ErrAutoMove) {
			c.Writer.Header().Add("success", "true")
			c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
		} else {
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
			_ = os.Remove("./ginvueadmin.zip")
		}
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s" , filename) to rename the downloaded file
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		_ = os.Remove("./ginvueadmin.zip")
	}
}

// GetDB
// @Tags      AutoCode
// @Summary Get all current databases
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Get all current databases"
// @Router    /autoCode/getDatabase [get]
func (autoApi *AutoCodeApi) GetDB(c *gin.Context) {
	businessDB := c.Query("businessDB")
	dbs, err := autoCodeService.Database(businessDB).GetDB(businessDB)
	var dbList []map[string]interface{}
	for _, db := range global.GVA_CONFIG.DBList {
		var item = make(map[string]interface{})
		item["aliasName"] = db.AliasName
		item["dbName"] = db.Dbname
		item["disable"] = db.Disable
		item["dbtype"] = db.Type
		dbList = append(dbList, item)
	}
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs, "dbList": dbList}, "Get successful", c)
	}
}

// GetTables
// @Tags      AutoCode
// @Summary Get all tables in the current database
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Get all tables in the current database"
// @Router    /autoCode/getTables [get]
func (autoApi *AutoCodeApi) GetTables(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	businessDB := c.Query("businessDB")
	tables, err := autoCodeService.Database(businessDB).GetTables(businessDB, dbName)
	if err != nil {
		global.GVA_LOG.Error("Query table failed!", zap.Error(err))
		response.FailWithMessage("Query table failed", c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, "Get successful", c)
	}
}

// GetColumn
// @Tags      AutoCode
// @Summary Get all fields of the current table
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Get all fields of the current table"
// @Router    /autoCode/getColumn [get]
func (autoApi *AutoCodeApi) GetColumn(c *gin.Context) {
	businessDB := c.Query("businessDB")
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	tableName := c.Query("tableName")
	columns, err := autoCodeService.Database(businessDB).GetColumn(businessDB, tableName, dbName)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, "Get successful", c)
	}
}

// CreatePackage
// @Tags      AutoCode
// @Summary Create package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysAutoCode true "Create package"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Package created successfully"
// @Router    /autoCode/createPackage [post]
func (autoApi *AutoCodeApi) CreatePackage(c *gin.Context) {
	var a system.SysAutoCode
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.AutoPackageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := autoCodeService.CreateAutoCode(&a)
	if err != nil {

		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
	} else {
		response.OkWithMessage("Created successfully", c)
	}
}

// GetPackage
// @Tags      AutoCode
// @Summary Get package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Package created successfully"
// @Router    /autoCode/getPackage [post]
func (autoApi *AutoCodeApi) GetPackage(c *gin.Context) {
	pkgs, err := autoCodeService.GetPackage()
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
	} else {
		response.OkWithDetailed(gin.H{"pkgs": pkgs}, "Get successful", c)
	}
}

// DelPackage
// @Tags      AutoCode
// @Summary delete package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysAutoCode true "Create package"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Delete package successfully"
// @Router    /autoCode/delPackage [post]
func (autoApi *AutoCodeApi) DelPackage(c *gin.Context) {
	var a system.SysAutoCode
	_ = c.ShouldBindJSON(&a)
	err := autoCodeService.DelPackage(a)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
	} else {
		response.OkWithMessage("Deletion successful", c)
	}
}

// AutoPlug
// @Tags      AutoCode
// @Summary Create plug-in template
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysAutoCode true "Create plug-in template"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Plug-in template created successfully"
// @Router    /autoCode/createPlug [post]
func (autoApi *AutoCodeApi) AutoPlug(c *gin.Context) {
	var a system.AutoPlugReq
	err := c.ShouldBindJSON(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Snake = strings.ToLower(a.PlugName)
	a.NeedModel = a.HasRequest || a.HasResponse
	err = autoCodeService.CreatePlug(a)
	if err != nil {
		global.GVA_LOG.Error("Preview failed!", zap.Error(err))
		response.FailWithMessage("Preview failed", c)
		return
	}
	response.Ok(c)
}

// InstallPlugin
// @Tags      AutoCode
// @Summary Install plug-in
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     plug  formData  file                                              true  "this is a test file"
// @Success 200 {object} response.Response{data=[]interface{},msg=string} "Plug-in installed successfully"
// @Router    /autoCode/installPlugin [post]
func (autoApi *AutoCodeApi) InstallPlugin(c *gin.Context) {
	header, err := c.FormFile("plug")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	web, server, err := autoCodeService.InstallPlugin(header)
	webStr := "web plug-in installed successfully"
	serverStr := "server plug-in installed successfully"
	if web == -1 {
		webStr = "The web-side plug-in was not successfully installed. Please unzip and install it according to the document. If it is a pure back-end plug-in, please ignore this prompt."
	}
	if server == -1 {
		serverStr = "The server-side plug-in was not successfully installed. Please unzip and install it according to the document. If it is a pure front-end plug-in, please ignore this prompt."
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData([]interface{}{
		gin.H{
			"code": web,
			"msg":  webStr,
		},
		gin.H{
			"code": server,
			"msg":  serverStr,
		}}, c)
}

// PubPlug
// @Tags      AutoCode
// @Summary packaging plug-in
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysAutoCode true "Packaging plug-in"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Packaging plug-in successful"
// @Router    /autoCode/pubPlug [get]
func (autoApi *AutoCodeApi) PubPlug(c *gin.Context) {
	plugName := c.Query("plugName")
	zipPath, err := autoCodeService.PubPlug(plugName)
	if err != nil {
		global.GVA_LOG.Error("Packaging failed!", zap.Error(err))
		response.FailWithMessage("Packaging failed"+err.Error(), c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("Packaging successful, file path is: %s", zipPath), c)
}
