package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"goClass/global"
	"goClass/model"
	"goClass/utils"
	"gorm.io/gorm"
)

// 登录接口

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username  = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

// 注册接口

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return err, u
}
