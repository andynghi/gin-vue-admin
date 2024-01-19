package system

import (
	"errors"
	"fmt"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ast"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"go.uber.org/zap"
)

var RepeatErr = errors.New("Repeat creation")

type AutoCodeHistoryService struct{}

var AutoCodeHistoryServiceApp = new(AutoCodeHistoryService)

// CreateAutoCodeHistory creates code generator history
// RouterPath : RouterPath@RouterString;RouterPath2@RouterString2
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) CreateAutoCodeHistory(meta, structName, structCNName, autoCodePath string, injectionMeta string, tableName string, apiIds string, Package string, BusinessDB string) error {
	return global.GVA_DB.Create(&system.SysAutoCodeHistory{
		Package:       Package,
		RequestMeta:   meta,
		AutoCodePath:  autoCodePath,
		InjectionMeta: injectionMeta,
		StructName:    structName,
		StructCNName:  structCNName,
		TableName:     tableName,
		ApiIDs:        apiIds,
		BusinessDB:    BusinessDB,
	}).Error
}

// First get the code generator history data based on id
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) First(info *request.GetById) (string, error) {
	var meta string
	return meta, global.GVA_DB.Model(system.SysAutoCodeHistory{}).Select("request_meta").Where("id = ?", info.Uint()).First(&meta).Error
}

// Repeat detects repetitions
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) Repeat(businessDB, structName, Package string) bool {
	var count int64
	global.GVA_DB.Model(&system.SysAutoCodeHistory{}).Where("business_db = ? and struct_name = ? and package = ? and flag = 0", businessDB, structName, Package).Count(&count)
	return count > 0
}

// RollBack rollback
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) RollBack(info *systemReq.RollBack) error {
	md := system.SysAutoCodeHistory{}
	if err := global.GVA_DB.Where("id = ?", info.ID).First(&md).Error; err != nil {
		return err
	}
	// Clear API table

	ids := request.IdsReq{}
	idsStr := strings.Split(md.ApiIDs, ";")
	for i := range idsStr[0 : len(idsStr)-1] {
		id, err := strconv.Atoi(idsStr[i])
		if err != nil {
			return err
		}
		ids.Ids = append(ids.Ids, id)
	}
	err := ApiServiceApp.DeleteApisByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("ClearTag DeleteApiByIds:", zap.Error(err))
	}
	// delete table
	if info.DeleteTable {
		if err = AutoCodeServiceApp.DropTable(md.BusinessDB, md.TableName); err != nil {
			global.GVA_LOG.Error("ClearTag DropTable:", zap.Error(err))
		}
	}
	// Delete Files

	for _, path := range strings.Split(md.AutoCodePath, ";") {

		// Add security judgment patch:
		_path, err := filepath.Abs(path)
		if err != nil || _path != path {
			continue
		}

		// migration
		nPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root,
			"rm_file", time.Now().Format("20060102"), filepath.Base(filepath.Dir(filepath.Dir(path))), filepath.Base(filepath.Dir(path)), filepath.Base(path))
		// Determine whether the target file exists
		for utils.FileExist(nPath) {
			fmt.Println("File already exists:", nPath)
			nPath += fmt.Sprintf("_%d", time.Now().Nanosecond())
		}
		err = utils.FileMove(path, nPath)
		if err != nil {
			global.GVA_LOG.Error("file move err ", zap.Error(err))
		}
		//_ = utils.DeLFile(path)
	}
	// clear injection
	for _, v := range strings.Split(md.InjectionMeta, ";") {
		// RouterPath@functionName@RouterString
		meta := strings.Split(v, "@")
		if len(meta) == 3 {
			_ = utils.AutoClearCode(meta[0], meta[2])
		}
	}

	ast.RollBackAst(md.Package, md.StructName)

	md.Flag = 1
	return global.GVA_DB.Save(&md).Error
}

// Delete delete historical data
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) Delete(info *request.GetById) error {
	return global.GVA_DB.Where("id = ?", info.Uint()).Delete(&system.SysAutoCodeHistory{}).Error
}

// GetList obtains system historical data
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (autoCodeHistoryService *AutoCodeHistoryService) GetList(info request.PageInfo) (list []response.AutoCodeHistory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysAutoCodeHistory{})
	var entities []response.AutoCodeHistory
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&entities).Error
	return entities, total, err
}
