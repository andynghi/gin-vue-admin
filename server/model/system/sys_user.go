package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid/v5"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"index;comment:UserUUID"`                                                          // UserUUID
	Username    string         `json:"userName" gorm:"index;comment:user login name"`                                               // User login name
	Password    string         `json:"-" gorm:"comment:User login password"`                                                        // User login password
	NickName    string         `json:"nickName" gorm:"default:system user;comment:user nickname"`                                   // User nickname
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:User side theme"`                                        // User side theme
	HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:user avatar"` // User avatar
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:base color"`                                            // base color
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:active color"`                                     // Active color
	AuthorityId uint           `json:"authorityId" gorm:"default:888;comment:User role ID"`                                         // User role ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:user role"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string         `json:"phone" gorm:"comment:User's mobile phone number"`                               // User's mobile phone number
	Email       string         `json:"email" gorm:"comment:user's email"`                                             // User's email
	Enable      int            `json:"enable" gorm:"default:1;comment: Whether the user is frozen 1 normal 2 frozen"` // Whether the user is frozen 1 normal 2 frozen
}

func (SysUser) TableName() string {
	return "sys_users"
}
