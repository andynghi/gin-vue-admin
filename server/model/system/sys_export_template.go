// Automatically generate template SysExportTemplate
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Export template structure SysExportTemplate
type SysExportTemplate struct {
	global.GVA_MODEL
	Name         string `json:"name" form:"name" gorm:"column:name;comment:template name;"`                              //Template name
	TableName    string `json:"tableName" form:"tableName" gorm:"column:table_name;comment:table name;"`                 //Table name
	TemplateID   string `json:"templateID" form:"templateID" gorm:"column:template_id;comment:template identification;"` //Template identification
	TemplateInfo string `json:"templateInfo" form:"templateInfo" gorm:"column:template_info;type:text;"`                 //Template information
}
