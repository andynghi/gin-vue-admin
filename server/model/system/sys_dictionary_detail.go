// Automatically generate template SysDictionaryDetail
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// If it contains time.Time, please import the time package yourself
type SysDictionaryDetail struct {
	global.GVA_MODEL
	Label           string `json:"label" form:"label" gorm:"column:label;comment:display value"`                                  // display value
	Value           int    `json:"value" form:"value" gorm:"column:value;comment:dictionary value"`                               // dictionary value
	Extend          string `json:"extend" form:"extend" gorm:"column:extend;comment:extended value"`                              // Extended value
	Status          *bool  `json:"status" form:"status" gorm:"column:status;comment:enabled status"`                              // Enabled status
	Sort            int    `json:"sort" form:"sort" gorm:"column:sort;comment:sort mark"`                                         // Sort mark
	SysDictionaryID int    `json:"sysDictionaryID" form:"sysDictionaryID" gorm:"column:sys_dictionary_id;comment:associated tag"` // associated tag
}

func (SysDictionaryDetail) TableName() string {
	return "sys_dictionary_details"
}
