package system

type SysAuthorityBtn struct {
	AuthorityId      uint           `gorm:"comment:role ID"`
	SysMenuID        uint           `gorm:"comment:menu ID"`
	SysBaseMenuBtnID uint           `gorm:"comment:menu button ID"`
	SysBaseMenuBtn   SysBaseMenuBtn ` gorm:"comment:Button Details"`
}
