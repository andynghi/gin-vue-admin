package system

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:MenuID"`
	AuthorityId uint                   `json:"-" gorm:"comment:role ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
	Btns        map[string]uint        `json:"btns" gorm:"-"`
}

type SysAuthorityMenu struct {
	MenuId      string `json:"menuId" gorm:"comment:menuID;column:sys_base_menu_id"`
	AuthorityId string `json:"-" gorm:"comment:role ID;column:sys_authority_authority_id"`
}

func (s SysAuthorityMenu) TableName() string {
	return "sys_authority_menus"
}
