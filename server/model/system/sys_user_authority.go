package system

// SysUserAuthority is the connection table between sysUser and sysAuthority
type SysUserAuthority struct {
	SysUserId               uint `gorm:"column:sys_user_id"`
	SysAuthorityAuthorityId uint `gorm:"column:sys_authority_authority_id"`
}

func (s *SysUserAuthority) TableName() string {
	return "sys_user_authority"
}
