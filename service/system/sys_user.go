package system

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"go-class/global"
	"go-class/model/common/request"
	"go-class/model/system"
	"go-class/utils"
	"gorm.io/gorm"
)

type UserService struct {
}

// 登录接口

func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username  = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

// 注册接口

func (userService *UserService) Register(u system.SysUser) (err error, userInter system.SysUser) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

func (userService *UserService) FindUserByUuid(uuid string) (err error, user *system.SysUser) {
	var u system.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (err error, userInter *system.SysUser) {
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var m system.SysUser
	db := global.GVA_DB.Model(&m)
	var userList []system.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}

func (userService *UserService) DeleteUser(id float64) (err error) {
	var m system.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&m).Error
	return err
}

func (userService *UserService) SetUserInfo(req system.SysUser) (err error, res system.SysUser) {
	err = global.GVA_DB.Updates(&req).Error
	return err, req
}
