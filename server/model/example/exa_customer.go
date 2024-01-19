package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type ExaCustomer struct {
	global.GVA_MODEL
	CustomerName       string         `json:"customerName" form:"customerName" gorm:"comment:CustomerName"`                    // Customer name
	CustomerPhoneData  string         `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:Customer phone number"` // Customer phone number
	SysUserID          uint           `json:"sysUserId" form:"sysUserId" gorm:"comment:Management ID"`                         // Management ID
	SysUserAuthorityID uint           `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:Management role ID"`  // Management role ID
	SysUser            system.SysUser `json:"sysUser" form:"sysUser" gorm:"comment:Management details"`                        // Management details
}
