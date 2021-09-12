package system

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gin-vue-admin-study/global"
	"gin-vue-admin-study/model/common/request"
	"gin-vue-admin-study/model/system/tables"
	"gin-vue-admin-study/utils"
	"gorm.io/gorm"
)

type UserService struct {
}

// 登录接口

func (userService *UserService) Login(u *tables.SysUser) (err error, userInter *tables.SysUser) {
	var user tables.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username  = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

// 注册接口

func (userService *UserService) Register(u tables.SysUser) (err error, userInter tables.SysUser) {
	var user tables.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

func (userService *UserService) FindUserByUuid(uuid string) (err error, user *tables.SysUser) {
	var u tables.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

func (userService *UserService) ChangePassword(u *tables.SysUser, newPassword string) (err error, userInter *tables.SysUser) {
	var user tables.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var m tables.SysUser
	db := global.GVA_DB.Model(&m)
	var userList []tables.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}

func (userService *UserService) DeleteUser(id float64) (err error) {
	var m tables.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&m).Error
	return err
}

func (userService *UserService) SetUserInfo(req tables.SysUser) (err error, res tables.SysUser) {
	err = global.GVA_DB.Updates(&req).Error
	return err, req
}

// 设置用户的权限

func (userService *UserService) SetUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	assignErr := global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&tables.SysUseAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&tables.SysUser{}).Error
	return err
}

// 设置一个用户的权限

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]tables.SysUseAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		useAuthority := []tables.SysUseAuthority{}
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, tables.SysUseAuthority{id, v})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回nil提交事务
		return nil
	})
}

// 获取用户信息

func (u *UserService) GetUserInfo(uuid uuid.UUID) (err error, res tables.SysUser) {
	var req tables.SysUser
	err = global.GVA_DB.Preload("Authority").First(&req, "uuid = ?", uuid).Error
	return err, req
}
