package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysBaseMenu struct {
	global.GVA_MODEL
	MenuLevel     uint                                                        `json:"-"`
	ParentId      string                                                      `json:"parentId" gorm:"comment:parent menu ID"`                     // Parent menu ID
	Path          string                                                      `json:"path" gorm:"comment:routing path"`                           // Routing path
	Name          string                                                      `json:"name" gorm:"comment:route name"`                             // route name
	Hidden        bool                                                        `json:"hidden" gorm:"comment:Whether to hide in the list"`          // Whether to hide in the list
	Component     string                                                      `json:"component" gorm:"comment:Corresponding front-end file path"` // Corresponding front-end file path
	Sort          int                                                         `json:"sort" gorm:"comment:sort mark"`                              // Sort mark
	Meta          `json:"meta" gorm:"embedded;comment:Additional attributes"` // Additional attributes
	SysAuthoritys []SysAuthority                                              `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu                                               `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter                                      `json:"parameters"`
	MenuBtn       []SysBaseMenuBtn                                            `json:"menuBtn"`
}

type Meta struct {
	ActiveName  string `json:"activeName" gorm:"comment:highlight menu"`
	KeepAlive   bool   `json:"keepAlive" gorm:"comment: Whether to cache"`                            // Whether to cache
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment: Is it a basic routing (under development)"` // Is it a basic routing (under development)
	Title       string `json:"title" gorm:"comment:menu name"`                                        //Menu name
	Icon        string `json:"icon" gorm:"comment:menu icon"`                                         // Menu icon
	CloseTab    bool   `json:"closeTab" gorm:"comment:Automatically close tab"`                       // Automatically close tab
}

type SysBaseMenuParameter struct {
	global.GVA_MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment: Whether the parameters carried in the address bar are params or query"` // Whether the parameters carried in the address bar are params or query
	Key           string `json:"key" gorm:"comment:key that carries parameters in the address bar"`                          //key that carries parameters in the address bar
	Value         string `json:"value" gorm:"comment:The value of the parameter carried in the address bar"`                 //The value of the parameter carried in the address bar
}

func (SysBaseMenu) TableName() string {
	return "sys_base_menus"
}
