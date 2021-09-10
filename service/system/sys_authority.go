package system

import (
	"errors"
	"go-class/global"
	commonReq "go-class/model/common/request"
	"go-class/model/system/tables"
	"gorm.io/gorm"
)

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

// 创建角色

func (a *AuthorityService) CreateAuthority(props tables.SysAuthority) (err error, authority tables.SysAuthority) {
	var m tables.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", props.AuthorityId).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), props
	}
	err = global.GVA_DB.Create(&props).Error
	return err, props
}

// 删除角色

func (a *AuthorityService) DeleteAuthority(props *tables.SysAuthority) (err error) {

	if !errors.Is(global.GVA_DB.Where("authority_id = ?", props.AuthorityId).First(&tables.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用,禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("parent_id = ?", props.AuthorityId).First(&tables.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色,不允许删除")
	}

	//db := global.GVA_DB.Preload("sys_base_menus").Where("authority_id = ?", props.AuthorityId).First(props)
	//err = db.Unscoped().Delete(props).Error
	//if len(props.SysBaseMenus) > 0 {
	//	err = global.GVA_DB.Model(props).Association("sys_base_menus").Delete(props.SysBaseMenus)
	//} else {
	//	err = db.Error
	//}

	if errors.Is(global.GVA_DB.Where("sys_authority_authority_id = ?", props.AuthorityId).First(&tables.SysUseAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("没有这个角色")
	}

	err = global.GVA_DB.Delete(&[]tables.SysUseAuthority{}, "sys_authority_authority_id = ?", props.AuthorityId).Error

	return err
}

// 分页获取

func (a *AuthorityService) GetAuthorityInfoList(req commonReq.PageInfo) (err error, list interface{}, total int64) {
	limit := req.PageSize
	offset := req.Page * (req.Page - 1)
	db := global.GVA_DB
	var authority []tables.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("AuthorityIdList").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = a.findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

// 查询子角色
func (a *AuthorityService) findChildrenAuthority(req *tables.SysAuthority) (err error) {
	err = global.GVA_DB.Preload("AuthorityIdList").Where("parent_id = ?", req.AuthorityId).Find(&req.Children).Error
	if len(req.Children) > 0 {
		for k := range req.Children {
			err = a.findChildrenAuthority(&req.Children[k])
		}
	}
	return err
}

// 更新角色

func (a *AuthorityService) UpdateAuthority(req tables.SysAuthority) (err error, res tables.SysAuthority) {
	err = global.GVA_DB.Where("authority_id = ?", req.AuthorityId).First(&tables.SysAuthority{}).Updates(&req).Error
	return err, req
}
