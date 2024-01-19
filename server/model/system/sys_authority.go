package system

import (
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time       // Creation time
	UpdatedAt       time.Time       // Update time
	DeletedAt       *time.Time      `sql:"index"`
	AuthorityId     uint            `json:"authorityId" gorm:"not null;unique;primary_key;comment:role ID;size:90"` // Role ID
	AuthorityName   string          `json:"authorityName" gorm:"comment:role name"`                                 // role name
	ParentId        *uint           `json:"parentId" gorm:"comment:parent role ID"`                                 // Parent role ID
	DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children        []SysAuthority  `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users           []SysUser       `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter   string          `json:"defaultRouter" gorm:"comment:default menu;default:dashboard"` //Default menu (default dashboard)
}

func (SysAuthority) TableName() string {
	return "sys_authorities"
}
