package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Register User register structure
type Register struct {
	Username     string `json:"userName" example:"username"`
	Password     string `json:"passWord" example:"password"`
	NickName     string `json:"nickName" example:"nickname"`
	HeaderImg    string `json:"headerImg" example:"Avatar link"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int role id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int whether to enable"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint role id"`
	Phone        string `json:"phone" example:"phone number"`
	Email        string `json:"email" example:"email"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // Username
	Password  string `json:"password"`  // Password
	Captcha   string `json:"captcha"`   // Verification code
	CaptchaId string `json:"captchaId"` // Verification code ID
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // Extract user id from JWT to avoid unauthorized access
	Password    string `json:"password"`    // Password
	NewPassword string `json:"newPassword"` // New password
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // Role ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // Role ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                                  // Primary key ID
	NickName     string                `json:"nickName" gorm:"default:system user;comment:user nickname"`                                   // User nickname
	Phone        string                `json:"phone" gorm:"comment:User's mobile phone number"`                                             // User's mobile phone number
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                       // Role ID
	Email        string                `json:"email" gorm:"comment:user's email"`                                                           // User's email
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:user avatar"` // User avatar
	SideMode     string                `json:"sideMode" gorm:"comment:User side theme"`                                                     // User side theme
	Enable       int                   `json:"enable" gorm:"comment:Freeze user"`                                                           //Freeze user
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
