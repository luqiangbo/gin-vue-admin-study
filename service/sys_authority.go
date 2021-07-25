package service

import (
	"errors"
	"goClass/global"
	"goClass/model"
	"gorm.io/gorm"
)

func CreateAuthority(props model.SysAuthority) (err error, authority model.SysAuthority) {
	var m model.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", props.AuthorityId).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), props
	}
	err = global.GVA_DB.Create(&props).Error
	return err, props
}
