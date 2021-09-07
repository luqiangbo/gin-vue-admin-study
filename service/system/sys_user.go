package system

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"go-class/global"
	"go-class/model/common/request"
	modelSystemDb "go-class/model/system/tables"
	"go-class/utils"
	"gorm.io/gorm"
)

type UserService struct {
}

// 登录接口

func (userService *UserService) Login(u *modelSystemDb.SysUser) (err error, userInter *modelSystemDb.SysUser) {
	var user modelSystemDb.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username  = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

// 注册接口

func (userService *UserService) Register(u modelSystemDb.SysUser) (err error, userInter modelSystemDb.SysUser) {
	var user modelSystemDb.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

func (userService *UserService) FindUserByUuid(uuid string) (err error, user *modelSystemDb.SysUser) {
	var u modelSystemDb.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

func (userService *UserService) ChangePassword(u *modelSystemDb.SysUser, newPassword string) (err error, userInter *modelSystemDb.SysUser) {
	var user modelSystemDb.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var m modelSystemDb.SysUser
	db := global.GVA_DB.Model(&m)
	var userList []modelSystemDb.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}

func (userService *UserService) DeleteUser(id float64) (err error) {
	var m modelSystemDb.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&m).Error
	return err
}

func (userService *UserService) SetUserInfo(req modelSystemDb.SysUser) (err error, res modelSystemDb.SysUser) {
	err = global.GVA_DB.Updates(&req).Error
	return err, req
}

// 设置用户的权限

func (userService *UserService) SetUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	assignErr := global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&modelSystemDb.SysUseAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&modelSystemDb.SysUser{}).Error
	return err
}

// 设置一个用户的权限

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]modelSystemDb.SysUseAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		useAuthority := []modelSystemDb.SysUseAuthority{}
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, modelSystemDb.SysUseAuthority{id, v})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回nil提交事务
		return nil
	})
}
