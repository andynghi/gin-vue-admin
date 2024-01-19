package system

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCasbin
//@description: Update casbin permissions
//@param: authorityId string, casbinInfos []request.CasbinInfo
//@return: error

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (casbinService *CasbinService) UpdateCasbin(AuthorityID uint, casbinInfos []request.CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityID))
	casbinService.ClearCasbin(0, authorityId)
	rules := [][]string{}
	//Do permission deduplication processing
	deduplicateMap := make(map[string]bool)
	for _, v := range casbinInfos {
		key := authorityId + v.Path + v.Method
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("The same api exists, the addition failed, please contact the administrator")
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCasbinApi
//@description: Follow API updates
//@param: oldPath string, newPath string, oldMethod string, newMethod string
//@return: error

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.GVA_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	e := casbinService.Casbin()
	err = e.LoadPolicy()
	if err != nil {
		return err
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPolicyPathByAuthorityId
//@description: Get the permission list
//@param: authorityId string
//@return: pathMaps []request.CasbinInfo

func (casbinService *CasbinService) GetPolicyPathByAuthorityId(AuthorityID uint) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	authorityId := strconv.Itoa(int(AuthorityID))
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ClearCasbin
//@description: Clear matching permissions
//@param: v int, p ...string
//@return: bool

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: RemoveFilteredPolicy
//@description: Use the database method to clean the filtered policy. This method needs to call the FreshCasbin method to take effect immediately in the system.
//@param: db *gorm.DB, authorityId string
//@return: error

func (casbinService *CasbinService) RemoveFilteredPolicy(db *gorm.DB, authorityId string) error {
	return db.Delete(&gormadapter.CasbinRule{}, "v0 = ?", authorityId).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: RemoveFilteredPolicy
//@description: Synchronize the policy of the current database. This method needs to call the FreshCasbin method to take effect immediately in the system.
//@param: db *gorm.DB, authorityId string, rules [][]string
//@return: error

func (casbinService *CasbinService) SyncPolicy(db *gorm.DB, authorityId string, rules [][]string) error {
	err := casbinService.RemoveFilteredPolicy(db, authorityId)
	if err != nil {
		return err
	}
	return casbinService.AddPolicies(db, rules)
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ClearCasbin
//@description: Clear matching permissions
//@param: v int, p ...string
//@return: bool

func (casbinService *CasbinService) AddPolicies(db *gorm.DB, rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
		})
	}
	return db.Create(&casbinRules).Error
}

func (CasbinService *CasbinService) FreshCasbin() (err error) {
	e := CasbinService.Casbin()
	err = e.LoadPolicy()
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Casbin
//@description: persist to database and introduce custom rules
//@return: *casbin.Enforcer

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.GVA_DB)
		if err != nil {
			zap.L().Error("Failed to adapt to the database, please check whether the casbin table is an InnoDB engine!", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("String failed to load model!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}
