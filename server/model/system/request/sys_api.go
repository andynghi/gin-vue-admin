package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// API paging condition query and sorting structure
type SearchApiParams struct {
	system.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey"` // Sorting
	Desc     bool   `json:"desc"`     // Sorting method: ascending false (default) | descending true
}
