package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SysBaseMenuBtn struct {
	global.GVA_MODEL
	Name          string `json:"name" gorm:"comment:button key"`
	Desc          string `json:"desc" gorm:"Button remarks"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:MenuID"`
}
