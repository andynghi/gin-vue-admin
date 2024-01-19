package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ExaFileUploadAndDownload struct {
	global.GVA_MODEL
	Name string `json:"name" gorm:"comment:file name"`   // file name
	Url  string `json:"url" gorm:"comment:file address"` // file address
	Tag  string `json:"tag" gorm:"comment:file tag"`     // file tag
	Key  string `json:"key" gorm:"comment:number"`       // Number
}

func (ExaFileUploadAndDownload) TableName() string {
	return "exa_file_upload_and_downloads"
}
