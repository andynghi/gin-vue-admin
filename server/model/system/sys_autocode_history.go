package system

import (
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// SysAutoCodeHistory automatically migrates code records for rollback and replay use
type SysAutoCodeHistory struct {
	global.GVA_MODEL
	Package       string `json:"package"`
	BusinessDB    string `json:"businessDB"`
	TableName     string `json:"tableName"`
	RequestMeta   string `gorm:"type:text" json:"requestMeta,omitempty"`   // Structured information passed in by the front end
	AutoCodePath  string `gorm:"type:text" json:"autoCodePath,omitempty"`  // Other meta information path;path
	InjectionMeta string `gorm:"type:text" json:"injectionMeta,omitempty"` // Injected content RouterPath@functionName@RouterString;
	StructName    string `json:"structName"`
	StructCNName  string `json:"structCNName"`
	ApiIDs        string `json:"apiIDs,omitempty"` // api table registration content
	Flag          int    `json:"flag"`             // Indicates the corresponding status 0 represents creation, 1 represents rollback...
}

// ToRequestIds ApiIDs conversion request.IdsReq
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *SysAutoCodeHistory) ToRequestIds() request.IdsReq {
	if m.ApiIDs == "" {
		return request.IdsReq{}
	}
	slice := strings.Split(m.ApiIDs, ";")
	ids := make([]int, 0, len(slice))
	length := len(slice)
	for i := 0; i < length; i++ {
		id, _ := strconv.ParseInt(slice[i], 10, 32)
		ids = append(ids, int(id))
	}
	return request.IdsReq{Ids: ids}
}
