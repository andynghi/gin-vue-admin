package system

import (
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Login
// @Tags     Base
// @Summary User login
// @Produce   application/json
// @Param data body systemReq.Login true "Username, password, verification code"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=string} "Return includes user information, token, expiration time"
// @Router   /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// Determine whether the verification code is enabled
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // Whether to enable explosion-proof times
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // Cache timeout
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.GVA_LOG.Error("Login failed! The username does not exist or the password is wrong!", zap.Error(err))
			// Verification code times +1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("The username does not exist or the password is wrong", c)
			return
		}
		if user.Enable != 1 {
			global.GVA_LOG.Error("Login failed! User is prohibited from logging in!")
			// Verification code times +1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("The user is prohibited from logging in", c)
			return
		}
		b.TokenNext(c, *user)
		return
	}
	// Verification code times +1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("Verification code error", c)
}

// TokenNext issues jwt after logging in
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // Unique signature
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("Failed to obtain token!", zap.Error(err))
		response.FailWithMessage("Failed to obtain token", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "Login successful", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("Failed to set login status!", zap.Error(err))
			response.FailWithMessage("Failed to set login status", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "Login successful", c)
	} else if err != nil {
		global.GVA_LOG.Error("Failed to set login status!", zap.Error(err))
		response.FailWithMessage("Failed to set login status", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt invalidation failed", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("Failed to set login status", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "Login successful", c)
	}
}

// Register
// @Tags     SysUser
// @Summary User registration account
// @Produce   application/json
// @Param data body systemReq.Register true "Username, nickname, password, role ID"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse,msg=string} "User registered account, return includes user information"
// @Router   /user/admin_register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("Registration failed!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "Registration failed", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "Registration successful", c)
}

// ChangePassword
// @Tags      SysUser
// @Summary User changes password
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.ChangePasswordReq true "Username, original password, new password"
// @Success 200 {object} response.Response{msg=string} "User changes password"
// @Router    /user/changePassword [post]
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	u := &system.SysUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("Modification failed!", zap.Error(err))
		response.FailWithMessage("Modification failed, the original password does not match the current account", c)
		return
	}
	response.OkWithMessage("Modification successful", c)
}

// GetUserList
// @Tags      SysUser
// @Summary Get the user list by pagination
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.PageInfo true "Page number, size of each page"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Get the user list in paging, and return the list, total number, page number, and number per page"
// @Router    /user/getUserList [post]
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get successful", c)
}

// SetUserAuthority
// @Tags      SysUser
// @Summary Change user permissions
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body systemReq.SetUserAuth true "User UUID, role ID"
// @Success 200 {object} response.Response{msg=string} "Set user permissions"
// @Router    /user/setUserAuthority [post]
func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	err = userService.SetUserAuthority(userID, sua.AuthorityId)
	if err != nil {
		global.GVA_LOG.Error("Modification failed!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := utils.GetUserInfo(c)
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // Unique signature
	claims.AuthorityId = sua.AuthorityId
	if token, err := j.CreateToken(*claims); err != nil {
		global.GVA_LOG.Error("Modification failed!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("new-token", token)
		c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
		utils.SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
		response.OkWithMessage("Modification successful", c)
	}
}

// SetUserAuthorities
// @Tags      SysUser
// @Summary Set user permissions
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body systemReq.SetUserAuthorities true "User UUID, role ID"
// @Success 200 {object} response.Response{msg=string} "Set user permissions"
// @Router    /user/setUserAuthorities [post]
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.SetUserAuthorities(sua.ID, sua.AuthorityIds)
	if err != nil {
		global.GVA_LOG.Error("Modification failed!", zap.Error(err))
		response.FailWithMessage("Modification failed", c)
		return
	}
	response.OkWithMessage("Modification successful", c)
}

// DeleteUser
// @Tags      SysUser
// @Summary delete user
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.GetById true "User ID"
// @Success 200 {object} response.Response{msg=string} "Delete user"
// @Router    /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("Deletion failed, suicide failed", c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// SetUserInfo
// @Tags      SysUser
// @Summary Set user information
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysUser true "ID, username, nickname, avatar link"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Set user information"
// @Router    /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(user, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(user.AuthorityIds) != 0 {
		err = userService.SetUserAuthorities(user.ID, user.AuthorityIds)
		if err != nil {
			global.GVA_LOG.Error("Setting failed!", zap.Error(err))
			response.FailWithMessage("Setting failed", c)
			return
		}
	}
	err = userService.SetUserInfo(system.SysUser{
		GVA_MODEL: global.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("Setting failed!", zap.Error(err))
		response.FailWithMessage("Setting failed", c)
		return
	}
	response.OkWithMessage("Set successfully", c)
}

// SetSelfInfo
// @Tags      SysUser
// @Summary Set user information
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysUser true "ID, username, nickname, avatar link"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Set user information"
// @Router    /user/SetSelfInfo [put]
func (b *BaseApi) SetSelfInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = userService.SetSelfInfo(system.SysUser{
		GVA_MODEL: global.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("Setting failed!", zap.Error(err))
		response.FailWithMessage("Setting failed", c)
		return
	}
	response.OkWithMessage("Set successfully", c)
}

// GetUserInfo
// @Tags      SysUser
// @Summary Get user information
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Get user information"
// @Router    /user/getUserInfo [get]
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	ReqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "Get successful", c)
}

// ResetPassword
// @Tags      SysUser
// @Summary Reset user password
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysUser                 true  "ID"
// @Success 200 {object} response.Response{msg=string} "Reset user password"
// @Router    /user/resetPassword [post]
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var user system.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(user.ID)
	if err != nil {
		global.GVA_LOG.Error("Reset failed!", zap.Error(err))
		response.FailWithMessage("Reset failed"+err.Error(), c)
		return
	}
	response.OkWithMessage("Reset successful", c)
}
