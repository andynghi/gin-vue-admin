package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysApi struct {
	global.GVA_MODEL
	Path        string `json:"path" gorm:"comment:api path"`                       // api path
	Description string `json:"description" gorm:"comment:api Chinese description"` // api Chinese description
	ApiGroup    string `json:"apiGroup" gorm:"comment:api group"`                  // api group
	Method      string `json:"method" gorm:"default:POST;comment:method"`          // Method: Create POST (default) | View GET | Update PUT | Delete DELETE
}

func (SysApi) TableName() string {
	return "sys_apis"
}
