// Automatically generate template SysDictionary
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// If it contains time.Time, please import the time package yourself
type SysDictionary struct {
	global.GVA_MODEL
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:Dictionary name (Chinese)"` // Dictionary name (Chinese)
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:Dictionary name (English)"` // Dictionary name (English)
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:status"`              // Status
	Desc                 string                `json:"desc" form:"desc" gorm:"column:desc;comment:description"`               // Description
	SysDictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}

func (SysDictionary) TableName() string {
	return "sys_dictionaries"
}
